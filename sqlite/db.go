package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-model/ambition-service"
)

type database struct {
	db *sql.DB
}

func InitDatabase(conn string) (*database, error) {
	d, err := sql.Open("sqlite3", conn)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to %s", conn)
	}
	return &database{d}, nil
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

func (d *database) CreateOccurrence(in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	occurrence := in.GetOccurrence()

	const query = `INSERT occurrences SET action_id=?, datetime=?, data=?`
	id, err := exec(d.db, query, occurrence.GetActionID(), occurrence.GetDatetime(), occurrence.GetData())
	if err != nil {
		return nil, err
	}
	occurrence.ID = id

	return occurrence, nil
}

func (d *database) ReadActionByID(in *pb.Action) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE id=?`
	resp := d.db.QueryRow(query, in.GetID())
	err := resp.Scan(in.ID, in.Name, in.UserID)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (d *database) ReadActionByUserIdAndName(in *pb.Action) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE action_name=?, user_id=?`
	resp := d.db.QueryRow(query, in.GetName, in.GetUserID())
	err := resp.Scan(in.ID, in.Name, in.UserID)
	if err != nil {
		return nil, err
	}

	return in, nil
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