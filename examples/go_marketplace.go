package main

type User struct {
	id              string
	password        string
	tempPassword    string
	name            string
	email           string
	emailVerified   uint8
	phone           string
	banner          string
	banned          uint8
	bannedReason    string
	suspended       uint8
	suspendedReason string
	deleted         uint8
	created         string
	modified        string
}

type Images struct {
	id        string
	parentId  string
	productId string
	size      string
	name      string
	nameTag   string
	deleted   uint8
	created   string
	modified  string
}

type Publications struct {
	id        string
	userId    string
	title     string
	body      string
	price     string
	quantity  string
	status    uint8
	published uint8
	banned    uint8
	deleted   uint8
	created   string
	modified  string
}

type Banners struct {
	id       string
	parentId string
	userId   string
	size     string
	name     string
	nameTag  string
	deleted  uint8
	created  string
	modified string
}
