package bd

import "golang.org/x/crypto/bcrypt"

/*
	EncriptarPassword es la funcion encargada de encriptar una contraseña
*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //Es la cantidad de 2 elevado a este valor para pasar la encriptada, en este caso lo encripta 256 veces
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
