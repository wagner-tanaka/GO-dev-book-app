package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// take the body and convert to a user struct
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	// connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// defer the database connection close
	userRepository := repositories.NewUsersRepository(db)

	// create user, passing the user from request
	userID, err := userRepository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	// return the user ID
	w.Write([]byte(fmt.Sprintf("User created with ID %d", userID)))

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users!"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User!"))
}
