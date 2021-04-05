package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(f http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (this *User) ToJson() ([]byte, error) {
	return json.Marshal(this)
}
