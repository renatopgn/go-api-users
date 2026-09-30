package packages

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Id = uuid.UUID

type User struct {
	UserId    Id     `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

var ctx = context.Background()

func FindAll(db *pgxpool.Pool) ([]User, error) {
	var users []User

	res, err := db.Query(ctx, `SELECT id, first_name, last_name, biography FROM users`)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var user User

		err := res.Scan(
			&user.UserId,
			&user.FirstName,
			&user.LastName,
			&user.Biography,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func FindById(db *pgxpool.Pool, idProcurado Id) (User, error) {
	var user User

	res := db.QueryRow(ctx, `SELECT * FROM users WHERE id = ($1)`, idProcurado)

	err := res.Scan(
		&user.UserId,
		&user.FirstName,
		&user.LastName,
		&user.Biography,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Insert(db *pgxpool.Pool, newUser User) (User, error) {

	var created User

	id := uuid.New()

	res := db.QueryRow(ctx, `INSERT INTO users (id, first_name, last_name, biography) VALUES ($1, $2, $3, $4) RETURNING id, first_name, last_name, biography`,
		id, newUser.FirstName, newUser.LastName, newUser.Biography)

	err := res.Scan(
		&created.UserId,
		&created.FirstName,
		&created.LastName,
		&created.Biography,
	)

	if err != nil {
		return User{}, err
	}

	return created, nil

}

func Update(idProcurado Id, userUpdates User, db *pgxpool.Pool) (User, error) {
	var userUpdated User

	res := db.QueryRow(ctx, `UPDATE users SET first_name = $1 , last_name = $2, biography = $3 WHERE id = $4 RETURNING id, first_name, last_name, biography`,
		userUpdates.FirstName, userUpdates.LastName, userUpdates.Biography, idProcurado,
	)

	err := res.Scan(
		&userUpdated.UserId,
		&userUpdated.FirstName,
		&userUpdated.LastName,
		&userUpdated.Biography,
	)

	if err != nil {
		return User{}, err
	}

	return userUpdated, nil

}

func Delete(db *pgxpool.Pool, idProcurado Id) (User, error) {
	var userDeleted User

	res := db.QueryRow(ctx, `DELETE FROM users WHERE id = ($1) RETURNING id, first_name, last_name, biography`, idProcurado)

	err := res.Scan(
		&userDeleted.UserId,
		&userDeleted.FirstName,
		&userDeleted.LastName,
		&userDeleted.Biography,
	)

	if err != nil {
		return User{}, err
	}

	return userDeleted, nil
}
