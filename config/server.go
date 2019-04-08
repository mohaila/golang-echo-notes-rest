package config

import (
	"os"
	"strconv"
)

var (
	sc *ServerConfig
)

// GetServerConfig returns the app config
func GetServerConfig() *ServerConfig {
	if sc != nil {
		return sc
	}

	sc = &ServerConfig{}
	s, ok := os.LookupEnv(address)
	if !ok {
		sc.Address = "localhost"
	} else {
		sc.Address = s
	}

	s, ok = os.LookupEnv(dbport)
	if !ok {
		sc.Port = 3000
	} else {
		p, err := strconv.Atoi(s)
		if err != nil {
			sc.Port = 3000
		} else {
			sc.Port = p
		}
	}

	return sc
}
