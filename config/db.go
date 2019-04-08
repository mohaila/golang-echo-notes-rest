package config

import (
	"os"
	"strconv"
)

var (
	dc *DBConfig
)

// GetDBConfig returns DB config
func GetDBConfig() *DBConfig {
	if dc != nil {
		return dc
	}

	dc = &DBConfig{}
	s, ok := os.LookupEnv(dbhost)
	if !ok {
		dc.Host = "localhost"
	} else {
		dc.Host = s
	}

	s, ok = os.LookupEnv(dbport)
	if !ok {
		dc.DPort = 5432
	} else {
		p, err := strconv.Atoi(s)
		if err != nil {
			dc.DPort = 5432
		} else {
			dc.DPort = p
		}
	}

	s, ok = os.LookupEnv(dbuser)
	if !ok {
		dc.User = "go"
	} else {
		dc.User = s
	}

	s, ok = os.LookupEnv(dbpassword)
	if !ok {
		dc.Password = ""
	} else {
		dc.Password = s
	}

	s, ok = os.LookupEnv(dbname)
	if !ok {
		dc.Name = "go"
	} else {
		dc.Name = s
	}

	return dc
}
