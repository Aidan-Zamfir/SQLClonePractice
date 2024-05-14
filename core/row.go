package core

import "github.com/marianogappa/sqlparser/query"


type Row struct {
	id string
	username string
	password string
}

func RowFromQuery(q query.Query) Row {
	r := Row{} //create empty row
	for i, f := range q.Fields { // for index, field_name of range query of q.Fields
		 if f == "id" {
			r.id = q.Inserts[0][i]
		 } else if f == "username" {
			r.username = q.Inserts[0][i]
		 } else if f == "password" {
			r.password = q.Inserts[0][i]
		 }
	}
	return r //return row just created
}