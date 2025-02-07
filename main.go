package main

import (
	"fmt"
)

/*
GET /events                   Get a list of available events
GET /events/<id>              Get a specific event
POST /events                  Create a new bookable event (Auth required)
PUT /events/<id>              Update an event (Auth required, Only by creator)
DELETE /events/<id>           Delete an event (Auth required, Only by creator)
POST /signup                  Create a new user
POST /login                   Authenticate a user (Auth token JWT, used for auth required routes)
POST /events/<id>/register    Register user for an event (Auth required)
DELETE /events/<id>/register  Cancel registration (Auth required)

*/

func main(){
	fmt.Println("REST API!")

}