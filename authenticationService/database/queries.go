package database

var GetUserQuery = `SELECT id, email, first_name, last_name, password FROM users WHERE email=$1`
var SignupQuery = `INSERT INTO users (id, email, first_name, last_name, password) VALUES($1, $2, $3, $4, $5)`
var DeleteUserQuery = `DELETE FROM users WHERE email=$1`
