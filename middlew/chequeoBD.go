package middlew

import (
	"github.com/afavendanor/twittor.git/bd"
	"net/http"
)

/*
	ChequeoBD es el middleware que me permite conocer el estado de la base de datos
*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
