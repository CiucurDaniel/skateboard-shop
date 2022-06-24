package data

type user struct {
	id       string
	email    string
	password string
	fullname string
	role     int
}

// for the moment use no DB, only local storage
var UserList = []user{
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
