package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/afavendanor/twittor.git/middlew"
	"github.com/afavendanor/twittor.git/routers"
)

/*
	Manejadores setea mi puerto, el handler e inicia el servidor
*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/tweetByUser", middlew.ChequeoBD(middlew.ValidarJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(routers.EliminarTweet))).Methods("DELETE")


	router.HandleFunc("/avatar", middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.ChequeoBD(middlew.ValidarJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/banner", middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.ChequeoBD(middlew.ValidarJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/usuarios", middlew.ChequeoBD(middlew.ValidarJWT(routers.ListaUsuarios))).Methods("GET")

	router.HandleFunc("/tweets", middlew.ChequeoBD(middlew.ValidarJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
