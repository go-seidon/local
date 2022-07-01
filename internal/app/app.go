package app

const (
	ENV_DEV  = "dev"
	ENV_STG  = "stg"
	ENV_TEST = "test"
	ENV_PROD = "prod"
)

type App interface {
	Run() error
	Stop() error
}
