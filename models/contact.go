package models

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
