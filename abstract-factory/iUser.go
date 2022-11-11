package main

type IUser interface {
	getName() string
	getSobrenome() string
	getPassword() string
}

type User struct {
	nome      string
	sobrenome string
	password  string
}
