package middlew

import (
	"github.com/afavendanor/twittor.git/routers"
	"net/http"
)

/*
	ValidarJTW nos permite validar el JWT que viene en la peticion
 */
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token " + err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

