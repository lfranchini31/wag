package swagger

import (
	"github.com/go-openapi/strfmt"

	"gopkg.in/mgo.v2/bson"
)

// InitCustomFormats adds wag's custom formats to the global go-openapi/strfmt Default registry.
func InitCustomFormats() {
	m := mongoID("")
	strfmt.Default.Add("mongo-id", &m, isMongoID)
}

func isMongoID(str string) bool {
	return bson.IsObjectIdHex(str)
}

type mongoID string

// MarshalText turns this instance into text
func (e mongoID) MarshalText() ([]byte, error) {
	return []byte(string(e)), nil
}

// UnmarshalText hydrates this instance from text
func (e *mongoID) UnmarshalText(data []byte) error { // validation is performed later on
	*e = mongoID(string(data))
	return nil
}

// String representation of the Mongo ID.
func (e mongoID) String() string {
	return string(e)
}
