package model

import "github.com/go-openapi/strfmt"

type User struct {
    Base
	Name 		string
    Username 	string 
	Email  		*string          
	Password  	*string          
    Role_id 	int32 
	Last_access *strfmt.DateTime
}