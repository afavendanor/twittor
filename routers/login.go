package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/jwt"
	"github.com/afavendanor/twittor.git/models"
)


/*
	Login realiza el login del usuario
 */
func Login(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invàlidos" + err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", http.StatusBadRequest)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invàlidos" , http.StatusBadRequest)
		return
	}

	jwtkey, err := jwt.GenerarJWT(documento)
	if err != nil {
		http.Error(w, "Ocurriò un error al intentar generar el token correspondiente " + err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin {
		Token : jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*
		Se genera y se setea cookie del usuario desde el backend
	 */
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtkey,
		Expires: expirationTime,
	})

}

