package packages

import (
	"errors"

	"github.com/google/uuid"
)

type Id = uuid.UUID

type User struct {
	UserId    Id     `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type application struct {
	data map[Id]User
}

func FindAll(data map[Id]User) ([]User, error) {
	var user []User
	for _, users := range data {
		user = append(user, users)
	}
	return user, nil
}

func FindById(data map[Id]User, idProcurado Id) (User, error) {
	usuario, existe := data[idProcurado]
	if !existe {
		return User{}, errors.New("user not found")
	}

	return usuario, nil

}

func Insert(data map[Id]User, newUser User) User {
	NewId := uuid.New()
	newUser.UserId = NewId
	data[NewId] = newUser

	user := data[NewId]
	return user

}

func Update(idProcurado Id, userUpdates User, data map[Id]User) (User, error) {
	usuario, existe := data[idProcurado]
	if !existe {
		return User{}, errors.New("user not found")
	}

	usuario.FirstName = userUpdates.FirstName
	usuario.LastName = userUpdates.LastName
	usuario.Biography = userUpdates.Biography

	data[idProcurado] = usuario
	return usuario, nil
}

func Delete(data map[Id]User, idProcurado Id) (User, error) {
	usuario, existe := data[idProcurado]
	if !existe {
		return User{}, errors.New("user not found")
	}

	delete(data, idProcurado)

	return usuario, nil
}
