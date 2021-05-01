package models

import (
	"database/sql"
	"time"

	"github.com/miguel1411/go-api-test/src/database"
	"gorm.io/gorm"
)

// gorm.Model definition
type Client struct {
	gorm.Model
	ID                uint `gorm:"primaryKey"`
	Nombre            string
	Apellido          string
	Correo            string `gorm:"unique"`
	Edad              uint8
	EntidadFederativa string
	Orders            []Order
	ActivatedAt       sql.NullTime
	CreatedAt         time.Time
	UpdatedAt         time.Time
	// DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func DataClient() {
	db := database.GetConnection()

	db.AutoMigrate(&Client{})

	db.Create(&Client{Nombre: "MIGUEL", Apellido: "CORONA", Correo: "MCORL95@GMAIL.COM", Edad: 26, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "GRECIA", Apellido: "GARCIA", Correo: "GGM@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "JESUS", Apellido: "LOPEZ", Correo: "LOPRX@GMAIL.COM", Edad: 22, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "ESTALI", Apellido: "MARQUEZ", Correo: "BOMBA@GMAIL.COM", Edad: 12, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "CHOCOLATE", Apellido: "NEGRO", Correo: "PERRO@GMAIL.COM", Edad: 14, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "CHOCOLATE", Apellido: "BLANCO", Correo: "GATO@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "MANUEL", Apellido: "GARCIA", Correo: "MANUEL@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "JHIN", Apellido: "CUATRO", Correo: "JHIN@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "ASHE", Apellido: "HIELO", Correo: "ASHE@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "SIVIR", Apellido: "AHCHA", Correo: "SIVIR@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "SION", Apellido: "GIGANTE", Correo: "SION@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "RYZE", Apellido: "AMD", Correo: "RYZE@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "LUX", Apellido: "ANA", Correo: "LUX@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "MISS", Apellido: "FORTUNE", Correo: "MISS@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "MORGANA", Apellido: "BRUJA", Correo: "MORGANA@GMAIL.COM", Edad: 24, EntidadFederativa: "SONORA"})
	db.Create(&Client{Nombre: "PYKE", Apellido: "PIRATA", Correo: "PYKE@GMAIL.COM", Edad: 25, EntidadFederativa: "SONORA"})

}
