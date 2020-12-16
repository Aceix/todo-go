package models

import "github.com/Kamva/mgm/v2"

type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Desc             string `json:"desc" bson:"desc"`
	Done             bool   `json:"done" bson:"done"`
}

func CreateTodo(title, desc string) *Todo {
	return &Todo{
		Title: title,
		Desc:  desc,
		Done:  false,
	}
}
