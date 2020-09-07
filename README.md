# API

Welcome on football Manager API !
The project is the basis ground API written in Go to let you manage your football team and is far from over. It will continue to evolve.

### Tech

The project uses some known golang packages (you can easily find more here https://github.com/avelino/awesome-go):

* [gin] - The API framework used to power the project
* [gorm] - The ORM to efficiently query Databases
* [jwt-go] - The Authorization system is build atop JWT (JSON Web Tokens)
* [pflag] - To Parse command arguments
* [testify] - To efficiently write tests


### How to use
Since go 1.11 we have a pretty management system integrated to go commands tool: go modules.
The go.mod / go.sum files handles already all the dependencies used in the project so everything should be imported for you at build (except if those libraries need you to install system packages).
The project uses a vendor mostly to avoid downloading the depen over and over

### Architecture
* bin/ - binary file will be build in there
* endpoints/ - Contains endpoints declarations & code
    * actions - Contains all endpoint actions code
* middlewares/ - Code in between the API & the endpoint action (middlewares)
* migrations/ - Databases migration files will be placed here
* models/ - Data models build atop gorm
* serverapi/ - The API server build atop gin code is here
* test/ - Contains test data and a go package containing code to write tests more easily
    * data - Contains test data (for now as .sql file(s))
* vendor/ - Dependencies are vendored
* docker-compose.yml - Classic file to manage containers with docker-compose (more details in docker section)
* Makefile - Makes life easier when it comes about testing or doing manual operations
* go.mod/go.sum - Go dependencies management system files
* main.go - If you need to start reading code, start with this file
* Readme.md - Guess what you are reading

### Starting the API within a container
So far the API does not need to run in background. A dedicated configuration will have be done in this case.

To set up the API within docker:
```sh
$ make start_docker_api
```

(With the current docker-compose config) You can call the API using
* From Outside the container network
```sh
$ curl -i "192.168.0.4:8081/ping"
```
from Inside the API container network
```sh
$ curl -i "0.0.0.0:8081/ping"
```
The API is running using configuration from environment variables:
* GIN_MODE - The API mode to show more or less informations about errors & what the API does (optional)
* API_HOSTNAME - the API host (mandatory)
* DB_DSN - Database DSN for connection (mandatory)
* API_LOG_FILE - The log file the api is writing to is specified in this file (mandatory)
* AUTH_SECRET - JWT secret passphrase to encrypt/decrypt authorization tokens


You can print the full list using:
```sh
$ apibinary --help
```

#### Testing & Benchmarking
For now the API tests depend on a database container (mock could come later).
So the tests are run inside a container.
You can run the tests using following command
```sh
$ make run_docker_test
```
similary benchmarks can be run using
```sh
$ make run_docker_benchmark
```
### Docker-compose configurations
You will find 4 containers working by pair in distinct networks. The pair are not meant to interact.

* A pair API - Database used to run the API (because we might later supposed the API & databases containers will diverge from the tests pair version)
* Another pair API - Database used only to run tests/benchmarks.

The Makefile contains some targets to interact smoothly with the containers.
Note: You may need to do a make clean from time to time (clean up some docker stuffs)

### Todos

 - Improve code coverage
 - Use argument to setup the api (see github.com/spf13/pflag) to get configurations from a file/etc (see github.com/spf13/viper) then fallback on environment variables.
 - Use github.com/sirupsen/logru as logger
 - Write middlewares/decorators to log request parameters/response body
 - Database test passwords should not be crystal clear (hash them at least)
 - Tests could use a Datbase Mock

License
----

MIT
