package db

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"token"`
}

type VM struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Owner  string `json:"owner"`
}

var Users []User
var VMs []VM
