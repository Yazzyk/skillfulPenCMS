package models

type Conf struct {
	AppName  string
	Server   Server
	Logger   Log
	Database DB
}

type Log struct {
	Level string
	Path  string
}

type Server struct {
	Port         uint
	EnableLogger bool
}

type DB struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
}
