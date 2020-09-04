package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/server"
)

const (
	// Environnement variable name containing the API host
	apiAddressEnvVar = "API_HOSTNAME"
	// Environnement variable name containing DSN to connect to database
	databaseDsnEnvVar = "DB_DSN"
	// Environnement variable name containg API's config log
	apiLogFileEnvVar = "API_LOG_FILE"
	// Message for missing environnement variable
	missingEnvVar = "Missing environnement variable `%s`"
	// Database type (must be one of gorm supported database see https://gorm.io/docs/connecting_to_the_database.html).
	// If you want to chose another SQL Database than (mysql/mariadb), you will also need to import its driver
	sqlDatabase = "mysql"
)

// List all environnement variables required for the API to run
var environnementVariables = []string{apiAddressEnvVar, databaseDsnEnvVar, apiLogFileEnvVar}

// getEnvironnementSettings gets and checks all system environment variables are set up
func getEnvironnementSettings() map[string]string {
	environnementValues := map[string]string{}

	for _, envVariableName := range environnementVariables {
		value := os.Getenv(envVariableName)
		if value == "" {
			panic(fmt.Sprintf(missingEnvVar, envVariableName))
		}
		environnementValues[envVariableName] = value
	}

	return environnementValues
}

func main() {
	var (
		apiServer     *server.Server
		dbConnHandler *gorm.DB
		err           error
	)

	envSettings := getEnvironnementSettings()
	if dbConnHandler, err = gorm.Open(sqlDatabase, envSettings[databaseDsnEnvVar]); err != nil {
		panic("Initialize Database connection pool handler:" + err.Error())
	}
	models.SetDb(dbConnHandler)
	apiServer, err = server.Initialize(nil, dbConnHandler, envSettings[apiAddressEnvVar], envSettings[apiLogFileEnvVar])
	if err != nil {
		panic(err.Error())
	}

	defer apiServer.Close()
	apiServer.Start()
}
