package main

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/pflag"

	"github/vriaan/footballmanagerapi/endpoints"
	"github/vriaan/footballmanagerapi/middlewares"
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
var (
	environnementVariables = []string{apiAddressEnvVar, databaseDsnEnvVar, apiLogFileEnvVar, apiAuthorizationSecretKey}
	Usage                  = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprint(os.Stderr, "To run following environment variables must be set and not empty:\n")
		for _, envVar := range environnementVariables {
			fmt.Fprintf(os.Stderr, "\t> %s\n", envVar)
		}
		fmt.Fprint(os.Stderr, "You can also set up the API verbose mode using GIN_MODE (please, refer to godoc gin for how to use)\n")
		fmt.Fprint(os.Stderr, "\n")
	}
)

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
		apiServer *serverapi.API
		err       error
	)
	pflag.Usage = Usage
	pflag.Parse()

	envSettings := getEnvironnementSettings()
	if err = models.InitDatabaseConnection(sqlDatabase, envSettings[databaseDsnEnvVar]); err != nil {
		panic("Initialize Database connection pool handler:" + err.Error())
	}
	defer models.GetDB().Close()

	middlewares.SetAuthorizationPassphrase(envSettings[apiAuthorizationSecretKey])

	endpointList := endpoints.Get()
	apiServer, err = serverapi.Initialize(
		&endpointList,
		envSettings[apiAddressEnvVar],
		envSettings[apiLogFileEnvVar],
	)
	if err != nil {
		panic(err.Error())
	}

	defer apiServer.Close()
	apiServer.Start()
}
