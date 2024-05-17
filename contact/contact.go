package contact

import "github.com/graphql-go/graphql"

type Contact struct {
	ID        int64       `db:"contact_id" json:"contactId"`
	Name      string      `db:"name" json:"name"`
	FirstName string      `db:"first_name" json:"firstName"`
	LastName  string      `db:"last_name" json:"lastName"`
	GenderID  int         `db:"gender_id" json:"genderId"`
	DOB       interface{} `db:"dob" json:"dob"`
	Email     string      `db:"email" json:"email"`
	Phone     string      `db:"phone" json:"phone"`
	Address   string      `db:"address" json:"address"`
	PhotoPath string      `db:"photo_path" json:"photoPath"`
	CreatedAt string      `db:"created_at" json:"createdAt"`
	CreatedBy string      `db:"created_by" json:"createdBy"`
}

var ContactType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Contact",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"name":      &graphql.Field{Type: graphql.String},
		"firstName": &graphql.Field{Type: graphql.String},
		"lastName":  &graphql.Field{Type: graphql.String},
		"genderId":  &graphql.Field{Type: graphql.Int},
		"dob":       &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"phone":     &graphql.Field{Type: graphql.String},
		"address":   &graphql.Field{Type: graphql.String},
	},
},
)
