package jwt

import (
	"github.com/afavendanor/twittor.git/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/*
	GenerarJWT se encarga de generar el token de autenticaciòn
 */
func GenerarJWT(t models.Usuario) (string, error)  {

	miClave := []byte("MastersdelDesarrollo")

	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellido": t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

