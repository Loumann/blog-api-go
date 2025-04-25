package repos

import (
	"blog-api-go/config"
	"blog-api-go/internal/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var dbDriveName = "postgres"

func NewBusinessDatabase(config *config.Config, env *models.Environment) *sqlx.DB {
	conn := getConnectDatabase(env, config)
	if conn == nil {
		log.Fatal("failed to establish database connection")
	}
	db := sqlx.NewDb(conn, dbDriveName)
	return db
}

func getConnectDatabase(env *models.Environment, config *config.Config) *sql.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		env.PostgresUser,
		env.PostgresPassword,
		config.Dbname,
		config.SSlmode,
	)

	db, err := sql.Open(
		dbDriveName,
		connString,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf(`database open connect: %v`, err.Error())
	}
	return db
}
