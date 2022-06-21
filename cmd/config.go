/**
 * Configuration that used by many application main.
 * e.g: "cmd/grpc-app" or "cmd/rest-app"
 *
 * There is no validation here so the code that uses this config need to validate
 * and check if particular key that they need is already set.
 */
package cmd

type Config struct {
	GrpcPort int

	MongoHost string
	MongoPort int
	// In miliseconds how long connection should established
	// until considered as failed
	MongoConnectTimeout int

	FileDirectory string
}
