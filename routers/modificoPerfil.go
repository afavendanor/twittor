package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/models"
	"net/http"
)

/*	ModificarPerfil se utiliza para modificar datos del usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request)  {

	var t models.Usuario
	var status bool

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos " + err.Error(), http.StatusBadRequest)
		return
	}

	status, err = bd.ModificarRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurriò un error al intentar modificar el registro. Reintente nuevamente " + err.Error(), http.StatusInternalServerError)
	}

	if status == false {
		http.Error(w, "No se logrò modificar el registro del usuario. Reintente nuevamente " + err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

