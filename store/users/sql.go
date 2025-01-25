package users

import (
	"database/sql"

	"github.com/todo-list/models"
)

type userStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *userStore {
	return &userStore{DB: db}
}

func (us *userStore) CreateUser(user models.User) error {
	_, err := us.DB.Exec(`INSERT INTO users (name, email, password,country,
	occupation, phone) VALUES ($1,$2,$3,$4,$5,$6)`, user.Name, user.Email, user.Password,
		user.Country, user.Occupation, user.Phone)
	return err
}

func (us *userStore) GetUserByEmail(email string) (*models.User, error) {
	var (
		user models.User
	)
	if err := us.DB.QueryRow("SELECT * FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Country,
			&user.Occupation, &user.Phone); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil

}

func (us *userStore) UpdatePasswordById(password string, id int) error {
	_, err := us.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id)
	if err != nil {
		return err
	}

	return nil
}

func (us *userStore) GetUserById(id int) (*models.User, error) {
	var (
		user models.User
	)
	if err := us.DB.QueryRow("SELECT * FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Email, &user.Name); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (us *userStore) UpdateUserDetailsById(user models.User, id int) (*models.User, error) {
	_, err := us.DB.Exec("UPDATE SET phone_number=$2, country=$3, occupation=$4 WHERE id=$5", id)
	if err != nil {
		return nil, err
	}

	return &user, err
}
