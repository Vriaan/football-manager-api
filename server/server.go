package server

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github/vriaan/footballmanagerapi/server/endpoints"
)

const (
	// log datetime format layout
	apiDatetimeFormatLayout = "2006-01-02 15:04:05"
	// format 2006-01-02 15:04:05 | 127.0.0.1 | Latency 989.158µs | resp body size %d | 404 | GET "/footballer" {name=roro} [Error not found]
	apiLogFormat = "%s | From %s | Latency %s | resp body size %d | %d | %s %s %s\n"
)

// Server represents the API server build atop gin engine framework
type Server struct {
	Engine *gin.Engine
	// DbConnHandler handles a pool of database connections
	DbConnHandler  *gorm.DB
	Address        string
	LogFileHandler *os.File
}

// Start initiates the API server
func (s *Server) Start() {
	s.Engine.Run(s.Address)
}

// Close cleanly closes the API server that need to be shutdown
func (s *Server) Close() {
	s.DbConnHandler.Close()
	if s.LogFileHandler != nil {
		s.LogFileHandler.Close()
	}
}

// Initialize creates a new API server fully setup
func Initialize(providedAPIEngine *gin.Engine, dbConnHandler *gorm.DB, address, logFilePath string) (server *Server, err error) {
	var logFileHandler *os.File
	gin.DisableConsoleColor()

	// it is possible to passdown an already configurated apiEngine (can be useful for test purpose),
	// but if none provided, we create one
	apiEngine := providedAPIEngine
	if apiEngine == nil {
		apiEngine = gin.New()
		//TODO add another middlewares to log request parameters & response body
		apiEngine.Use(
			gin.LoggerWithFormatter(customLogFormat),
			gin.Recovery(),
		)

		if logFileHandler, err = createLogFile(logFilePath); err != nil {
			return
		}
		multiWriter := io.MultiWriter(logFileHandler, os.Stdout)
		gin.DefaultWriter = multiWriter
		gin.DefaultErrorWriter = multiWriter
	}

	err = endpoints.Register(apiEngine)
	server = &Server{
		Engine:         apiEngine,
		DbConnHandler:  dbConnHandler,
		Address:        address,
		LogFileHandler: logFileHandler,
	}
	return
}

// customLogFormatter returns the API logs format
func customLogFormat(param gin.LogFormatterParams) string {
	error := ""
	if param.ErrorMessage != "" {
		error = "[" + param.ErrorMessage + "]"
	}

	return fmt.Sprintf(
		apiLogFormat,
		param.TimeStamp.Format(apiDatetimeFormatLayout),
		param.ClientIP,
		param.Latency,
		param.BodySize,
		param.StatusCode,
		param.Method,
		param.Path,
		error,
	)
}

// initLogger creates every thing for custom logs (log filen, format)
func createLogFile(logFilePath string) (logFileHandler *os.File, err error) {
	logFilePath = "/var/log/api/api.log"
	logPath := filepath.Dir(logFilePath)

	if _, err = os.Stat(logPath); os.IsNotExist(err) {
		if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
			return
		}
	} else {
		return
	}

	logFileHandler, err = os.OpenFile(
		logFilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	return
}
