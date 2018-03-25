package models


import (
	t "time"
	"gopkg.in/mgo.v2/bson"
)

type(	
	User struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPasswords []byte `json:"hashPassword,omitempty"`
	}
	Task struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy string `json:"createdBy"`
		Name string `json:"name"`
		Description string `json:"description"`
		CreatedOn t.Time  `json:"createdOn,omitempty"`
		Due t.Time `json:"due,omitempty"`
		Status string `json:"status,omitempty"`
		Tags []string `json:"tags,omitempty"`
	}

	TaskNote struct{
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId bson.ObjectId `json:"taskId"`
		Description string `json:"description"`
		CreatedOn t.Time `json:"createdOn,omitempty"`
	}

)