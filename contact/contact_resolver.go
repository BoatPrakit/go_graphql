package contact

import (
	"github.com/graphql-go/graphql"
)

type ContactResolver struct {
	contact *Storage
}

func NewContactResolver(contactStorage *Storage) *ContactResolver {
	return &ContactResolver{
		contact: contactStorage,
	}
}

func (r *ContactResolver) Contacts(p graphql.ResolveParams) (interface{}, error) {
	contacts, err := r.contact.FindAll()
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *ContactResolver) ContactById(p graphql.ResolveParams) (interface{}, error) {
	contact, err := r.contact.FindByID(p.Args["id"].(int))
	if err != nil {
		return nil, err
	}

	return contact, nil
}
