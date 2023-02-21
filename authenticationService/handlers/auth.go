package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/rwiteshbera/microservices_demo/authenticationService/api"
	"github.com/rwiteshbera/microservices_demo/authenticationService/database"

	"github.com/rwiteshbera/microservices_demo/authenticationService/helpers"
	"github.com/rwiteshbera/microservices_demo/authenticationService/utils"
)

const dbTimeOut = 2 * time.Second

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(api.AuthHandlerInstance.Env.DB_DRIVER, api.AuthHandlerInstance.Env.DB_CONNECTION_STRING)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(3 * time.Minute)
	return db, nil
}

func SignupRouter(res http.ResponseWriter, req *http.Request) {
	userData := &database.User{}

	// Read user data from request body
	err := helpers.ReadJSON(res, req, userData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	// Hash the password
	newHashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a new user Id
	UID := uuid.New()

	// Connect with database
	db, err := ConnectDB()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert data in database
	_, err = db.ExecContext(ctx, database.SignupQuery, UID, userData.Email, userData.FirstName, userData.LastName, newHashedPassword)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a response
	err = helpers.WriteJSON(res, http.StatusOK, UID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get user data by email
func GetUser(res http.ResponseWriter, req *http.Request) {
	var userEmail string

	// Read user data from request body
	err := helpers.ReadJSON(res, req, userEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	// Connect with database
	db, err := ConnectDB()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var user database.User
	err = db.QueryRowContext(ctx, database.GetUserQuery, userEmail).Scan(&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a response
	err = helpers.WriteJSON(res, http.StatusOK, user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete user by email
func DeleteUserRouter(res http.ResponseWriter, req *http.Request) {
	var userEmail string

	// Read user data from request body
	err := helpers.ReadJSON(res, req, userEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	// Connect with database
	db, err := ConnectDB()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.ExecContext(ctx, database.DeleteUserQuery, userEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a response
	err = helpers.WriteJSON(res, http.StatusOK, "deleted")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
