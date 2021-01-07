package config

import (
	"os"

	"github.com/jinzhu/gorm"

	// Included to activate postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	VarDatabaseHost     = "localhost"
	VarDatabaseUsername = "postgres"
	VarDatabasePassword = "admin"
	VarDatabaseName     = "postgres"
	VarHttpPort         = "8080"
	VarGrpcAddr         = ":9000"
)

func InitVars() {
	get := func(org string, name string) (env string) {
		env = os.Getenv(name)
		if env != "" {
			return
		}
		env = org
		return
	}
	VarDatabaseHost = get(VarDatabaseHost, "DATABASE_HOST")
	VarDatabaseUsername = get(VarDatabaseUsername, "DATABASE_USERNAME")
	VarDatabasePassword = get(VarDatabasePassword, "DATABASE_PASSWORD")
	VarDatabaseName = get(VarDatabaseName, "DATABASE_NAME")
	VarHttpPort = get(VarHttpPort, "HTTP_PORT")
	VarGrpcAddr = get(VarGrpcAddr, "GRPC_ADDR")
}

func Connect() (*gorm.DB, error) {
	connectionURL := "host=" + VarDatabaseHost +
		" user=" + VarDatabaseUsername +
		" dbname=" + VarDatabaseName +
		" password=" + VarDatabasePassword +
		" sslmode=disable"
	return gorm.Open("postgres", connectionURL)
}
