package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github/vriaan/footballmanagerapi/endpoints"
	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/serverapi"
)

const (
	// Environnement variable name containing the API host
	apiAddressEnvVar = "API_HOSTNAME"
	// Environnement variable name containing DSN to connect to database
	databaseDsnEnvVar = "DB_DSN"
	// Environnement variable name containg API's config log
	apiLogFileEnvVar = "API_LOG_FILE"
	// Environnement variable name for authorization tokens to be encrypted/decrypted
	apiAuthorizationSecretKey = "AUTH_SECRET"
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
		fmt.Println(value)
		environnementValues[envVariableName] = value
	}

	return environnementValues
}

func main() {
	var (
		apiServer     *serverapi.API
		dbConnHandler *gorm.DB
		err           error
	)

	// TODO: Implement config argument fallback on environment variables ?
	envSettings := getEnvironnementSettings()
	if dbConnHandler, err = gorm.Open(sqlDatabase, envSettings[databaseDsnEnvVar]); err != nil {
		panic("Initialize Database connection pool handler:" + err.Error())
	}
	models.SetDb(dbConnHandler)

	endpointList := endpoints.Get()
	apiServer, err = serverapi.Initialize(
		&endpointList,
		dbConnHandler,
		envSettings[apiAddressEnvVar],
		envSettings[apiLogFileEnvVar],
	)
	if err != nil {
		panic(err.Error())
	}

	defer apiServer.Close()
	apiServer.Start()
}
