package main

import (
	"encoding/json"
	"net/http"
)

// type Middleware recibe y retorna http.HandlerFunc así todas las funciones que cuenten con esta firma se consideran de tipo middleware. Reciben y retornan lo mismo para poder encadenarse y pasar el request entre varios middleware, hasta llegar al handler.
type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
