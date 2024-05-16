package resolvers

import (
	"database/sql"

	"github.com/boatprakit/graphql/models"
	"github.com/graphql-go/graphql"
)

func NewContactResolver(db *sql.DB) *QueryResolver {
	return &QueryResolver{db}
}

func (r *QueryResolver) Contacts(p graphql.ResolveParams) (interface{}, error) {
	rows, _ := r.db.Query("SELECT * FROM contact")
	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.FirstName,
			&contact.LastName,
			&contact.GenderID,
			&contact.DOB,
			&contact.Email,
			&contact.Phone,
			&contact.Address,
			&contact.PhotoPath,
			&contact.CreatedAt,
			&contact.CreatedBy)

		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (r *QueryResolver) ContactById(p graphql.ResolveParams) (interface{}, error) {
	rows, _ := r.db.Query("SELECT * FROM contact WHERE contact_id = $1", p.Args["id"])
	var contact models.Contact
	if rows.Next() {
		err := rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.FirstName,
			&contact.LastName,
			&contact.GenderID,
			&contact.DOB,
			&contact.Email,
			&contact.Phone,
			&contact.Address,
			&contact.PhotoPath,
			&contact.CreatedAt,
			&contact.CreatedBy)

		if err != nil {
			return models.Contact{}, err
		}
	}
	return contact, nil
}
