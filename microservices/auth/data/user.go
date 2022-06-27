package data

import "errors"

type User struct {
	id       string
	email    string
	password string
	fullname string
	role     int
}

var UserAlreadyExists = errors.New("User already exists in the database")

// for the moment use no DB, only local storage
var UserList = []User{
	{
		id:       "1",
		email:    "ciucur.daniel@email.com",
		password: "abcde",
		fullname: "Ciucur Daniel",
		role:     1,
	},
	{
		id:       "2",
		email:    "mike.jon@email.com",
		password: "mike123",
		fullname: "Mike Jon",
		role:     1,
	},
}

func AddUserToDb(u User) error {
	if ok := checkUserExist(u); ok == false {
		UserList = append(UserList, u)
		return nil
	}
	return UserAlreadyExists
}

func checkUserExist(u User) bool {
	for i := 0; i < len(UserList); i++ {
		if u.email == UserList[i].email {
			return true
		}
	}
	return false
}
