package openapi

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ClientDataDao interface {
	PostgresCreate()
	PostgresRead()
}

func (u *ClientData) PostgresCreate(postgres Postgres) {
	var pgCtrImpl PostgresController = &postgres
	pgCtrImpl.Open()
	defer pgCtrImpl.Close()

	// Insert data
	err := postgres.db.Create(u).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (u *ClientData) PostgresRead(postgres Postgres) []ClientData {
	var pgCtrImpl PostgresController = &postgres
	pgCtrImpl.Open()
	defer pgCtrImpl.Close()

	var clientData []ClientData
	err := postgres.db.Where("client_id = ?", u.ClientId).Find(&clientData).Error
	if err != nil {
		log.Fatal(err)
	}
	return clientData
}
