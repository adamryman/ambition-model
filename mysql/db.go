package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	//"github.com/adamryman/db"

	pb "github.com/adamryman/ambition-model/ambition-service"
)

func Open(conn string) (pb.Database, error) {
	d, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to mysql with %s", conn)
	}
	if err := d.Ping(); err != nil {
		return nil, errors.Wrapf(err, "cannot make initial database connection to %s", conn)
	}

	return &database{d}, nil
}

type database struct {
	db *sql.DB
}

func (d *database) CreateAction(in *pb.Action) (*pb.Action, error) {
	const query = `INSERT actions SET action_name=?, user_id=?`
	id, err := exec(d.db, query, in.GetName, in.GetUserID)
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

func (d *database) CreateOccurrence(in *pb.Occurrence) (*pb.Occurrence, error) {
	const query = `INSERT occurrences SET action_id=?, datetime=?, data=?`
	id, err := exec(d.db, query, in.GetActionID(), in.GetDatetime(), in.GetData())
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

func (d *database) ReadActionByID(id int64) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE id=?`
	resp := d.db.QueryRow(query, id)
	var action pb.Action
	err := resp.Scan(action.ID, action.Name, action.UserID)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

func (d *database) ReadActionByNameAndUserID(name string, userID int64) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE action_name=?, user_id=?`
	resp := d.db.QueryRow(query, name, userID)
	var action pb.Action
	err := resp.Scan(action.ID, action.Name, action.UserID)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

// exec calls db.db.Exec with passed arguments and returns the id of the LastInsertId
func exec(db *sql.DB, query string, args ...interface{}) (int64, error) {
	resp, err := db.Exec(query, args...)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to exec query: %v", query)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, errors.Wrapf(err, "unable to get last id after query: %v", query)
	}

	return id, nil
}
