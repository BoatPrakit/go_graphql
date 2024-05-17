package contact

import "database/sql"

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) FindAll() ([]Contact, error) {
	rows, err := s.db.Query("SELECT * FROM contact")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
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
			&contact.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (s *Storage) FindByID(id int) (Contact, error) {
	rows, err := s.db.Query("SELECT * FROM contact WHERE contact_id = $1", id)
	if err != nil {
		return Contact{}, err
	}

	defer rows.Close()

	var contact Contact
	for rows.Next() {
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
			&contact.CreatedBy,
		)

		if err != nil {
			return Contact{}, err
		}
	}

	return contact, nil
}
