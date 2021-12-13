package models

import (
	"errors"
	"fmt"
)

/*

	"webservice/telemetryService"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	//slice that holds pointers to user objects
	users []*User
	// at package level, I don't need :=
	nextID = 1
)

func GetUsers() []*User {

	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User can't have ID already!")
	}
	u.ID = nextID
	nextID++

	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found, fool!", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found, fool!", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			// take slice of users up to index i & append to it all users after index and on
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found, fool!", id)
}
