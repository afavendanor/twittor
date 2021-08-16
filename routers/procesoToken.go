package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/models"
)

/*	Email es el valor del email del usuario utilizado en todos los endpoints */
var Email string
/*	IDUsuario es el valor del ID devuelo del modelo, que se usarà en todos los endpoints */
var IDUsuario string

/*	ProcesarToken se utiliza para extraer los valores del  token */
func ProcesarToken(token string) (*models.Claim, bool, string, error)  {
	miClave := []byte("MastersdelDesarrollo")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.CheckExistUser(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid{
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}

