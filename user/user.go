package user

type U struct {
	ID       int    `db:"id" json:"id"`
	UserID   string `db:"user_id" json:"userId"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
