package auth

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"github.com/boatprakit/graphql/user"
	"github.com/google/uuid"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Verify(email, password string) (bool, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	isFound := rows.Next()

	if !isFound {
		return false, nil
	}
	var u user.U
	err = rows.Scan(&u.UserID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return false, err
	}
	isSuccess := comparePasswords(u.Password, password)
	return isSuccess, nil

}

func (s *Storage) RegisterUser(input RegisterInput) (user.U, error) {
	userId := uuid.New().String()
	hashedPassword := hashPassword(input.Password)
	result, err := s.db.Exec("INSERT INTO users (user_id, email, password) VALUES ($1, $2, $3)", userId, input.Email, hashedPassword)
	if err != nil {
		return user.U{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user.U{}, err
	}

	var u = user.U{
		ID:       int(id),
		UserID:   userId,
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}
	return u, nil
}

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	sum := hasher.Sum([]byte("salt"))
	hashedPassword := hex.EncodeToString(sum)
	return hashedPassword
}

func comparePasswords(hashedPassword string, password string) bool {
	hashedInputPassword := hashPassword(password)
	return hashedPassword == hashedInputPassword
}
