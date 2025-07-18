package types

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "encoding/json"
)

type ID string

// MarshalBSON converts to MongoDB ObjectID
func (id ID) MarshalBSON() ([]byte, error) {
    oid, err := primitive.ObjectIDFromHex(string(id))
    if err != nil {
        return nil, err
    }
    return bson.Marshal(oid)
}

// UnmarshalBSON converts from MongoDB ObjectID
func (id *ID) UnmarshalBSON(data []byte) error {
    var oid primitive.ObjectID
    if err := bson.Unmarshal(data, &oid); err != nil {
        return err
    }
    *id = ID(oid.Hex())
    return nil
}

// Optional: Support JSON marshalling
func (id ID) MarshalJSON() ([]byte, error) {
    return json.Marshal(string(id))
}

func (id *ID) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    *id = ID(s)
    return nil
}

// Optional: String() method for debugging
func (id ID) String() string {
	return string(id)
}