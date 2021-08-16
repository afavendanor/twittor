package models

/*
	RespuestaLogin se utiliza para retornar el token del login
 */
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}