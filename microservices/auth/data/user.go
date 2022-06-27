package data

import "errors"

type User struct {
	Id       string
	Email    string
	Password string
	Fullname string
	Role     int
}

var UserAlreadyExists = errors.New("User already exists in the database")

// for the moment use no DB, only local storage
var UserList = []User{
	{
		Id:       "1",
		Email:    "ciucur.daniel@email.com",
		Password: "abcde",
		Fullname: "Ciucur Daniel",
		Role:     1,
	},
	{
		Id:       "2",
		Email:    "mike.jon@email.com",
		Password: "mike123",
		Fullname: "Mike Jon",
		Role:     1,
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
		if u.Email == UserList[i].Email {
			return true
		}
	}
	return false
}
