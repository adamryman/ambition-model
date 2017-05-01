package mysql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-model/ambition-service"
)

func Open(conn string) (*Database, error) {
	d, err := sql.Open("sqlite3", conn)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to mysql with %s", conn)
	}
	if err := d.Ping(); err != nil {
		return nil, errors.Wrapf(err, "cannot make initial database connection to %s", conn)
	}

	if err := setupDB(d); err != nil {
		return nil, err
	}

	return &Database{d}, nil
}

func setupDB(db *sql.DB) error {
	const actions = `CREATE TABLE IF NOT EXISTS actions(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				action_name varchar(255),
				user_id integer);`
	_, err := db.Exec(actions)
	if err != nil {
		return err
	}

	const occurrences = `CREATE TABLE IF NOT EXISTS occurrences(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				action_id varchar(255),
				datetime varchar(255),
				data varchar(255));`
	_, err = db.Exec(occurrences)
	if err != nil {
		return err
	}

	return nil
}

type Database struct {
	db *sql.DB
}

func (d *Database) CreateAction(in *pb.Action) (*pb.Action, error) {
	const query = `INSERT INTO actions(action_name, user_id) VALUES (?, ?)`
	id, err := exec(d.db, query, in.GetName(), in.GetUserID())
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

func (d *Database) CreateOccurrence(in *pb.Occurrence) (*pb.Occurrence, error) {
	const query = `INSERT INTO occurrences(action_id, datetime, data) VALUES (?, ?, ?)`
	id, err := exec(d.db, query, in.GetActionID(), in.GetDatetime(), in.GetData())
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

func (d *Database) ReadActionByID(id int64) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE id=?`
	resp := d.db.QueryRow(query, id)
	var action pb.Action
	err := resp.Scan(&action.ID, &action.Name, &action.UserID)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

func (d *Database) ReadActionByNameAndUserID(name string, userID int64) (*pb.Action, error) {
	const query = `SELECT * FROM actions WHERE action_name=? AND user_id=?`
	resp := d.db.QueryRow(query, name, userID)
	var action pb.Action
	err := resp.Scan(&action.ID, &action.Name, &action.UserID)
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
