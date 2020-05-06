package openapi

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	postgres = Postgres{
		host:     "192.168.99.100",
		port:     "5432",
		user:     "root",
		dbname:   "workermonitordb",
		password: "admin",
		sslmode:  "disable",
	}
)

type PostgresController interface {
	Open()
	Close()
}

type Postgres struct {
	db       *gorm.DB
	host     string
	port     string
	user     string
	dbname   string
	password string
	sslmode  string
}

func NewPostgresController() *Postgres {
	p := new(Postgres)
	return p
}

func (u *Postgres) Open() {
	dbAccessParam := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		u.host, u.port, u.user, u.dbname, u.password, u.sslmode,
	)
	var err error
	u.db, err = gorm.Open("postgres", dbAccessParam)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *Postgres) Close() {
	u.db.Close()
}

func tempFunc() {
	// clientData := ClientData{
	// 	1,
	// 	"Taro Sato",
	// 	"2020-04-30 12:00:00",
	// 	"2020-04-30 22:00:00",
	// }
	// clientData.PostgresCreate(postgres)

	var temp ClientData
	temp.ClientId = 1
	targetdata := temp.PostgresRead(postgres)
	fmt.Println(targetdata)
}
