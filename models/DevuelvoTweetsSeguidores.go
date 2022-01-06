package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
	Devualve la estructura para leer tweets
*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id, omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId, omitempty"`
	UsuarioRelacionID string             `bson:"usuarioRelacionId" json:"userRelationId, omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje, omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha, omitempty"`
		ID      string    `bson:"_id" json:"_id, omitempty"`
	}
}