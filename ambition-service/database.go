package ambition

// Input to database must have necessary fields populated
type Database interface {
	// CreateAction requires Action.Name and Action.UserID
	CreateAction(in *Action) (*Action, error)
	// CreateOccurrence requires CreateOccurrenceRequest.UserID ,
	// .Occurrence.ActionId, and Occurrence.Datetime
	CreateOccurrence(in *Occurrence) (*Occurrence, error)
	// CreateAction requires Action.Name and Action.UserID
	ReadActionByID(id int64) (*Action, error)
	// CreateAction requires Action.Name and Action.UserID
	ReadActionByNameAndUserID(name string, userId int64) (*Action, error)
	// TODO: Add queries for all rpc's
}
