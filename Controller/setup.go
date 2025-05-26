package Controller

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AccountName  string
	AccountKey   string
	Server       string
	Port         int
	User         string
	Password     string
	Database     string
	Server2      string
	Port2        int
	User2        string
	Password2    string
	Database2    string
	Database3    string
	Database4    string
	Database5    string
	SmtpServer   string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		return nil, err
	}

	dbPort2, err := strconv.Atoi(os.Getenv("port2"))
	if err != nil {
		return nil, err
	}

	smtpPortdb, err := strconv.Atoi(os.Getenv("smtpPort"))
	if err != nil {
		return nil, err
	}

	Cfg := &Config{
		AccountName:  os.Getenv("accountName"),
		AccountKey:   os.Getenv("accountKey"),
		Server:       os.Getenv("server"),
		Port:         dbPort,
		User:         os.Getenv("user"),
		Password:     os.Getenv("password"),
		Database:     os.Getenv("database"),
		Server2:      os.Getenv("server2"),
		Port2:        dbPort2,
		User2:        os.Getenv("user2"),
		Password2:    os.Getenv("password2"),
		Database2:    os.Getenv("database2"),
		Database3:    os.Getenv("database3"),
		Database4:    os.Getenv("database4"),
		Database5:    os.Getenv("database5"),
		SmtpPort:     smtpPortdb,
		SmtpServer:   os.Getenv("smtpServer"),
		SmtpUsername: os.Getenv("smtpUsername"),
		SmtpPassword: os.Getenv("smtpPassword"),
	}

	return Cfg, nil
}
