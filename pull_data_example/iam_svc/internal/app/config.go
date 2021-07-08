package app

type Config struct {
	Port     int    // The HTTP listening port.
	EnvStage string // The environment stage (dev|qa|prod).
	Db       struct {
		DSN          string // The data source name.
		MaxOpenConns int    // The max number of open connections.
		MaxIdleConns int    // The max number of open connections.
		MaxIdleTime  string // // The max idle time after which the connection is closed.
	}
}
