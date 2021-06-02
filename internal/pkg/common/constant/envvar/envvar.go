package envvar

// DatabaseEnvVars defines environment variable names for database.
type DatabaseEnvVars struct {
	DriverName         string
	Host               string
	Port               string
	SSLMode            string
	ConnectTimeout     string
	User               string
	Password           string
	DbName             string
	ConnMaxIdlePercent string
	ConnMaxLifetime    string
	CheckTable         string
}

const (
	appPrefix = "BASEGO_"
	gcpPrefix = "GOOGLE_"
)

// App Configs
var (
	Environment = withAppPrefix("ENV")
	LogLevel    = withAppPrefix("LOG_LEVEL")
	ServerPort  = withAppPrefix("SERVER_PORT")
	BackendURL  = withAppPrefix("BACKEND_URL")
	FrontendURL = withAppPrefix("FRONTEND_URL")
)

// Database Configs
var Database = DatabaseEnvVars{
	DriverName:         withAppPrefix("DB_DRIVER"),
	Host:               withAppPrefix("DB_HOST"),
	Port:               withAppPrefix("DB_PORT"),
	SSLMode:            withAppPrefix("DB_SSLMODE"),
	ConnectTimeout:     withAppPrefix("DB_CONNECT_TIMEOUT"),
	User:               withAppPrefix("DB_USER"),
	Password:           withAppPrefix("DB_PASSWORD"),
	DbName:             withAppPrefix("DB_NAME"),
	ConnMaxIdlePercent: withAppPrefix("DB_CONN_MAX_IDLE_PERCENTAGE"),
	ConnMaxLifetime:    withAppPrefix("DB_CONN_MAX_LIFETIME"),
	CheckTable:         withAppPrefix("DB_CHECK_TABLE"),
}

// Redis Configs
var (
	RedisHost           = withAppPrefix("REDIS_HOST")
	RedisPort           = withAppPrefix("REDIS_PORT")
	RedisDatabase       = withAppPrefix("REDIS_DATABASE")
	RedisPassword       = withAppPrefix("REDIS_PASSWORD")
	RedisMaxConnections = withAppPrefix("REDIS_MAX_CONNECTIONS")
)

// Asset Configs
var Asset = struct{ BaseStaticAssetURL string }{
	BaseStaticAssetURL: withAppPrefix("BASE_STATIC_ASSET_URL"),
}

// Storage Configs
var Storage = struct{ DefaultStorage, LocalDirPath string }{
	DefaultStorage: withAppPrefix("DEFAULT_STORAGE"),
	LocalDirPath:   withAppPrefix("LOCAL_STORAGE_DIRPATH"),
}

// Google Cloud Storage
var Google = struct{ AppCredentials, ProjectID, StorageBucket, StorageClass, StorageLocation string }{
	AppCredentials:  withGooglePrefix("APPLICATION_CREDENTIALS"),
	ProjectID:       withGooglePrefix("PROJECT_ID"),
	StorageBucket:   withGooglePrefix("STORAGE_BUCKET"),
	StorageClass:    withGooglePrefix("STORAGE_CLASS"),
	StorageLocation: withGooglePrefix("STORAGE_LOCATION"),
}

// SMTP Configs
var SMTP = struct{ Host, Port, Username, Password, FromEmail, FromName string }{
	Host:      withAppPrefix("SMTP_HOST"),
	Port:      withAppPrefix("SMTP_PORT"),
	Username:  withAppPrefix("SMTP_USERNAME"),
	Password:  withAppPrefix("SMTP_PASSWORD"),
	FromEmail: withAppPrefix("SMTP_FROM_EMAIL"),
	FromName:  withAppPrefix("SMTP_FROM_NAME"),
}

func withAppPrefix(key string) string {
	return appPrefix + key
}

func withGooglePrefix(key string) string {
	return gcpPrefix + key
}
