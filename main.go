package main

import (
	"database/sql"
	_ "database/sql"
	db "erp-back/catalog/persistence/sqlc"
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
		serverAddress = "0.0.0.0:8799"
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

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	queries := db.New(conn)

	router := gin.Default()
	rest.RouteUpsertCategory(router)
	err = router.Run(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
