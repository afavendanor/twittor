package bd

import (
	"github.com/afavendanor/twittor.git/models"
	"golang.org/x/crypto/bcrypt"
)

/*
	IntentoLogin realiza el chequeo del email ante la bd
 */
func IntentoLogin(email string, password string) (models.Usuario, bool)  {
	usu, encontrado, _ := CheckExistUser(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
