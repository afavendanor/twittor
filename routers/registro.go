package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/models"
	"net/http"
)

/*
	Registro es la funcion para crear el usuario en la base de datos
*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es rquerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener minimo 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.CheckExistUser(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado en base de datos con ese email", http.StatusConflict)
		return
	}

	_, status, err := bd.InsertUSer(t)
	if err != nil {
		http.Error(w, "Ocurriò un error al insertar el registro en base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
