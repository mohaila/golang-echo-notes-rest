package config

type (
	// DBConfig is the PSQL config
	DBConfig struct {
		Host     string
		DPort    int
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	// ServerConfig is the server config
	ServerConfig struct {
		Address string
		Port    int
	}
)

const (
	dbhost     = "DBHOST"
	dbport     = "DBPORT"
	dbuser     = "DBUSER"
	dbpassword = "DBPASSWORD"
	dbname     = "DBNAME"
	address    = "ADDRESS"
	port       = "PORT"
	sslmode    = "DBSSLMODE"
)
