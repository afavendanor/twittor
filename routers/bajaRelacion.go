package routers

import (
	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/models"
	"net/http"
)

/*
	Se encarga de borar una relacion entre dos users
 */
func BajaRelacion(w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID!", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al borrar la relación " + err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
