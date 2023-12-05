package repository

const (
	CreateUser = `INSERT INTO users (name, surname, password, email, balance, is_register_social, is_paid, role)
VALUES (:name, :surname, :password, :email, :balance, :is_register_social, :is_paid, :role);`
	FindByEmail = `SELECT * from users where email=?`
	GetUsers    = `SELECT * from users;`
	GetByID     = `SELECT * from users where id=?;`
)
