package chat_server

import (
	"flag"
	"time"
)

type flags struct {
	Name           string
	Port           string
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	DB             string
	AllowedIP      string
	AllowedMethods string
	SessionAddr    string
}

var config flags

func init() {
	flag.StringVar(&config.Name, "project name", "chat-server", "set name of project")
	flag.StringVar(&config.Port, "port", ":8084", "service port")
	flag.DurationVar(&config.WriteTimeout, "write timeout", 15*time.Second, "timeout for write")
	flag.DurationVar(&config.ReadTimeout, "read timeout", 15*time.Second, "timeout for read")
	flag.StringVar(&config.DB, "database DSN", "user=postgres password=postgres dbname=postgres sslmode=disable", "DSN for database")
	flag.StringVar(&config.AllowedIP, "allowed IP", "http://212.109.223.57:10002", "IP for CORS")
	flag.StringVar(&config.AllowedMethods, "allowed HTTP methods", "POST, GET, PUT, DELETE, OPTIONS", "HTTP methids for CORS")
	flag.StringVar(&config.SessionAddr, "session addres", "127.0.0.1:8081", "addres of session microservice")
}