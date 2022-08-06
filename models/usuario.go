package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID        primitive.ObjectID `bson: "_id, omitempty" json : "id"`
	Nombre    string             `bson: "nombre" json: "nombre, omitempty"`
	Apellido  string             `bson: "apellido" json: "apellido, omitempty"`
	FechaNaci time.Time          `bson: "fechaNaci" json: "fechaNaci, omitempty"`
	Email     string             `bson: "email" json: "email, omitempty"`
	Password  string             `bson: "password" json: "password, omitempty"`
	Avatar    string             `bson: "avatar" json: "avatar, omitempty"`
	Banner    string             `bson: "banner" json: "banner, omitempty"`
	Biografia string             `bson: "biografia" json: "biografia, omitempty"`
	Ubicacion string             `bson: "ubicacion" json: "ubicacion, omitempty"`
	SitioWeb  string             `bson: "sitioWeb" json: "sitioWeb, omitempty"`
}
