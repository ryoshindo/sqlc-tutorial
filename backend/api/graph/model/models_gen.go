// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}

type SigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
