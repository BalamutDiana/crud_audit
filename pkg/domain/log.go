package audit

import (
	"time"
)

const (
	ENTITY_USER  = "USER"
	ENTITY_MOVIE = "MOVIE"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}
