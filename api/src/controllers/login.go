package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// pega os dados do body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// cria um usuário vazio
	var user models.User
	// trata o erro, decodifica o body e coloca o body no user
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// faz uma conexao com o database
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// cria um repositorio de usuarios, esse repositorio contem a conexao com o database e o user
	userRepository := repositories.NewUsersRepository(db)

	// busca o usuario no banco de dados pelo email
	userFromDatabase, err := userRepository.SearchByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userFromDatabase.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
	}

	token, _ := auth.CreateToken(userFromDatabase.ID)
	fmt.Println(token)
}
