package main

import (
	_ "database/sql"
	create_category "erp-back/catalog/usecases/create_category/rest"

	update_category "erp-back/catalog/usecases/update_category/rest"
	"erp-back/framework"
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

	//conn, err := sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("cannot connect to db:", err)
	//}
	//defer func(conn *sql.DB) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("cannot close connection to db:", err)
	//	}
	//}(conn)
	framework.InitDatabase()

	router := gin.Default()
	create_category.RouteCreateCategory(router)
	update_category.RouteUpdateCategory(router)
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
