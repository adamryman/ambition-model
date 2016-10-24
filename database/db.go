package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-model/ambition-service"
)

type Configuration struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

var db *sql.DB

func New() error {
	config := Configuration{
		DBName:     os.Getenv("MYSQL_DATABASE"),
		DBUser:     os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBPort:     os.Getenv("MYSQL_PORT"),
		DBHost:     os.Getenv("MYSQL_HOST"),
	}

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	tempdb, err := sql.Open("mysql", dbString)
	if err != nil {
		return errors.Wrapf(err, "could not open db connection %v", dbString)
	}

	err = tempdb.Ping()
	if err != nil {
		return errors.Wrapf(err, "could not ping database %v", dbString)
	}

	db = tempdb

	return nil
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

func CreateAction(req *pb.CreateActionRequest) (*pb.Action, error) {
	var action pb.Action
	const query = `INSERT actions SET action_name=?, user_id=?`

	id, err := exec(db, query, req.ActionName, req.UserId)
	if err != nil {
		return nil, err
	}

	action.ActionId = id
	action.UserId = req.UserId
	action.ActionName = req.ActionName

	return &action, nil
}

func ReadAction(req *pb.ReadActionRequest) (*pb.Action, error) {
	var action pb.Action
	const query = `SELECT * FROM actions WHERE action_id=?`

	resp := db.QueryRow(query, req.ActionId)

	err := resp.Scan(&action.ActionId, &action.ActionId, nil, &action.UserId)
	if err != nil {
		return nil, err
	}

	return &action, nil

}

func CreateOccurrence(req *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	var occurrence pb.Occurrence
	const query = `INSERT occurrences SET action_id=?, time=?`

	id, err := exec(db, query, req.ActionId, req.Datetime)
	if err != nil {
		return nil, err
	}

	occurrence.OccurrenceId = id
	occurrence.ActionId = req.ActionId
	occurrence.Datetime = req.Datetime

	return &occurrence, nil
}

/*
func (db DB) CreateUser(user) error {
	const query = `INSERT INTO users (username, email, password_salt, hashed_password) VALUES ($1,$2,$3,$4)`

	_, err := db.Exec(query, user.UserName, user.Email, user.PasswordSalt, user.HashedPassword)

	return err
}


func (db DB) GetUserByUserName(userName string) (*User, error) {
	const query = `SELECT id, email, password_salt, hashed_password FROM users WHERE username = $1`
	var reval User

	err := db.QueryRow(query, userName).Scan(&reval.Id, &reval.Email, &reval.PasswordSalt, &reval.HashedPassword)
	reval.UserName = userName
	return &reval, err
}

func (db DB) GetUserById(id int) (*User, error) {
	const query = `SELECT username, email, password_salt, hashed_password FROM users WHERE id = $1`
	var reval User

	err := db.QueryRow(query, id).Scan(&reval.UserName, &reval.Email, &reval.PasswordSalt, &reval.HashedPassword)
	reval.Id = id
	return &reval, err
}

func (db DB) InsertSession(userId int, hashedToken string) error {
	const query = `INSERT INTO sessions (user_id, hashed_token) VALUES ($1, $2)`

	_, err := db.Exec(query, userId, hashedToken)

	return err
}

func (db DB) GetSessionKeysByUserId(userId int) ([]string, error) {
	const query = `SELECT hashed_token FROM sessions, users WHERE user_id = $1 and users.id = sessions.user_id`
	var reval []string

	rows, err := db.Query(query, userId)
	defer rows.Close()
	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		check(err)
		reval = append(reval, key)
	}
	return reval, err
}

func (db DB) DeleteSessionByUserId(userId int) error {
	const query = `DELETE FROM sessions WHERE user_id = $1`

	_, err := db.Exec(query, userId)

	return err
}

// ----------------------------- Sets  ----------------------------- //
func (db DB) GetSets() ([]Set, error) {
	const query = `SELECT * FROM sets`
	var reval []Set

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var set Set
		err := rows.Scan(&set.Id, &set.SetName)
		check(err)
		reval = append(reval, set)
	}
	return reval, err
}

func (db DB) GetSetById(id int) (*Set, error) {
	const query = `SELECT set_name FROM sets WHERE id = $1`
	var reval Set
	err := db.QueryRow(query, id).Scan(&reval.SetName)
	reval.Id = id
	return &reval, err
}

func (db DB) InsertSet(set *Set) error {
	const query = `INSERT INTO sets (set_name) VALUES ($1)`

	_, err := db.Exec(query, set.SetName)

	return err
}

func (db DB) DeleteSetById(setId int) error {
	const query = `DELETE FROM sets WHERE id = $1`

	_, err := db.Exec(query, setId)

	return err
}

// ----------------------------- Actions  ----------------------------- //

func (db DB) GetActions() ([]Action, error) {
	const query = `SELECT id, action_name, set_id FROM actions`
	var reval []Action

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.Id, &action.ActionName, &action.SetId)
		check(err)
		reval = append(reval, action)
	}
	return reval, err
}

func (db DB) GetActionById(id int) (*Action, error) {
	const query = `SELECT action_name, set_id, user_id FROM actions WHERE id = $1`
	var reval Action
	err := db.QueryRow(query, id).Scan(&reval.ActionName, &reval.SetId, &reval.UserId)
	reval.Id = id

	return &reval, err
}

func (db DB) GetActionsByUserId(id int) ([]Action, error) {
	const query = `SELECT id, action_name, set_id, user_id FROM actions WHERE user_id = $1`
	var reval []Action

	rows, err := db.Query(query, id)
	defer rows.Close()
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.Id, &action.ActionName, &action.SetId, &action.UserId)
		check(err)
		reval = append(reval, action)
	}

	return reval, err
}

func (db DB) InsertAction(action *Action) error {
	const query = `INSERT INTO actions (action_name, set_id, user_id) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, action.ActionName, action.SetId, action.UserId)

	return err
}

func (db DB) UpdateActionById(action *Action) error {
	const query = `UPDATE actions SET action_name = $1, set_id = $2, user_id = $3 WHERE id = $4)`

	_, err := db.Exec(query, action.ActionName, action.SetId, action.UserId, action.Id)

	return err

}

func (db DB) DeleteActionById(actionId int) error {
	const query = `DELETE FROM actions WHERE id = $1`

	_, err := db.Exec(query, actionId)

	return err
}

// ----------------------------- Occurrences  ----------------------------- //

func (db DB) GetOccurrenceById(id int) (*Occurrence, error) {
	const query = `SELECT (action_name, time) FROM occurrences WHERE id = $1`
	var reval Occurrence
	err := db.QueryRow(query, id).Scan(&reval.ActionId, &reval.Time)
	reval.Id = id
	return &reval, err
}

func (db DB) GetOccurrencesOfAction(id int) ([]Occurrence, error) {
	const query = `SELECT * FROM occurrences WHERE action_id = $1`
	var reval []Occurrence

	rows, err := db.Query(query, id)
	defer rows.Close()
	for rows.Next() {
		var occurrence Occurrence
		err := rows.Scan(&occurrence.Id, &occurrence.ActionId, &occurrence.Time)
		check(err)
		reval = append(reval, occurrence)
	}
	return reval, err
}

func (db DB) InsertOccurrence(occurrence *Occurrence) error {
	const query = `INSERT INTO occurrences (action_id, time) VALUES ($1, $2)`

	_, err := db.Exec(query, occurrence.ActionId, occurrence.Time)

	return err
}

func (db DB) DeleteOccurrenceById(occurrenceId int) error {
	const query = `DELETE FROM occurrences WHERE id = $1`

	_, err := db.Exec(query, occurrenceId)

	return err
}

// ------------ Table Creation and Dropping -------------------

func (db DB) CreateUserTable() error {
	const query = `CREATE TABLE users(id SERIAL PRIMARY KEY, username varchar(255), email varchar(255), password_salt varchar(30), hashed_password varchar(255))`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropUserTable() error {
	const query = `DROP TABLE users`

	_, err := db.Exec(query)

	return err
}

func (db DB) CreateSessionTable() error {
	const query = `CREATE TABLE sessions(id SERIAL PRIMARY KEY, user_id integer, hashed_token varchar(255))`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropSessionTable() error {
	const query = `DROP TABLE sessions`

	_, err := db.Exec(query)

	return err
}

func (db DB) CreateSetTable() error {
	const query = `CREATE TABLE sets(id SERIAL PRIMARY KEY, set_name varchar(255))`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropSetTable() error {
	const query = `DROP TABLE sets`

	_, err := db.Exec(query)

	return err
}


// FUTURE:
// Will allow combining CreateTable and DropTable functions
func getTable(obj interface{}) string {
	switch obj.(type) {
	default:
		return "unknown"
	case Action, *Action:
		return "actions"
	case Set, *Set:
		return "sets"
	}
}
*/
