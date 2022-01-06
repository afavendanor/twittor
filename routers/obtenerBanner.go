package routers

import (
	"github.com/afavendanor/twittor.git/bd"
	"io"
	"net/http"
	"os"
)

/*
	obtiene el banner de bd y seridor
*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID!", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	openfile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, openfile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusInternalServerError)
		return
	}

}
