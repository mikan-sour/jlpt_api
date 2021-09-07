package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jedzeins/jlpt_api/userService/src/models"
	"github.com/jedzeins/jlpt_api/userService/src/services"
	"github.com/jedzeins/jlpt_api/userService/src/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	// check if not post
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Can only send post requests"))
		return
	}

	// get username and pw from json
	decoder := json.NewDecoder(r.Body)
	var credentials models.CredentialsPost
	decodeErr := decoder.Decode(&credentials)

	if decodeErr != nil {
		msg := errors.New(fmt.Errorf("error decoding the body: %w", decodeErr).Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg.Error()))
		return
	}

	// service to register user
	user, err := services.RegisterUser(credentials)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// service to register session

	session, err := services.RegisterSession(user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// service to get jwt refresh & access tokens

	refreshToken, accessToken, err := services.GenerateSignedTokens(user.ID, session.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// set refresh token as cookie
	utils.SetRefreshTokenCookie(refreshToken, w)

	// add access token to response
	user.AccessToken = accessToken

	// response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(accessToken))
	json.NewEncoder(w).Encode(user)

	return

	// service to create session (also get's refresh token)

	// service to get access token

	// status := services.HealthcheckService()
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(status)
}
