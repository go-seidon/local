package app

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppEnv     string `env:"APP_ENV"`
	AppVersion string `env:"APP_VERSION"`
	AppDebug   bool   `env:"APP_DEBUG"`

	RESTAppHost string `env:"REST_APP_HOST"`
	RESTAppPort int    `env:"REST_APP_PORT"`

	RPCAppHost string `env:"RPC_APP_HOST"`
	RPCAppPort int    `env:"RPC_APP_PORT"`

	DBProvider string `env:"DB_PROVIDER"`

	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLPort     int    `env:"MYSQL_PORT"`
	MySQLUser     string `env:"MYSQL_USER"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
	MySQLDBName   string `env:"MYSQL_DB_NAME"`
}
