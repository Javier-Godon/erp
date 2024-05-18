package main

import (
	_ "database/sql"
	"erp-back/catalog/usecases/upsert_category/rest"
	"github.com/gin-gonic/gin"
	//"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {

	const (
		dbDriver = "postgres"
		//dbSource      = "postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public"
		dbSource      = "postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public"
		serverAddress = "0.0.0.0:8887"
	)

	//m, err := migrate.New(
	//	"file://db/migration",
	//	//"postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public")
	//	"postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := m.Up(); err != nil {
	//	log.Fatal(err)
	//}

	router := gin.Default()
	rest.RouteUpsertCategory(router)
	err := router.Run("0.0.0.0:8799")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
