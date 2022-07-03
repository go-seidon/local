package app

const (
	ENV_LOCAL = "local"
	ENV_TEST  = "test"

	ENV_DEV  = "dev"
	ENV_STG  = "stg"
	ENV_PROD = "prod"
)

type App interface {
	Run() error
	Stop() error
}
