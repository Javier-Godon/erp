package framework

//_ "github.com/jackc/pgx/v5"
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func InitDatabase() {
	const (
		dbDriver = "postgres"
		//dbSource      = "postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public"
		dbSource = "postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public"
	)
	//database, err := sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("cannot connect to db:", err)
	//}
	//defer func(database *sql.DB) {
	//	err := database.Close()
	//	if err != nil {
	//		log.Fatal("cannot close connection to db:", err)
	//	}
	//}(database)
	//
	////defer func(conn *sql.DB) {
	////	err := conn.Close()
	////	if err != nil {
	////		log.Fatal("cannot close connection to db:", err)
	////	}
	////}(database)
	//database.SetMaxIdleConns(10)
	//database.SetMaxOpenConns(10)
	//err = database.Ping()
	//if err != nil {
	//	log.Fatal("cannot ping db:", err)
	//}
	//DB = database

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	DB = pool

}
