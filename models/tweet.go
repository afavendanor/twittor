package models

/*
	Tweet se utiliza para decodificar mensaje de entrada
 */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje, omitempty"`
}
