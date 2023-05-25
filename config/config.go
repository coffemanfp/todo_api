package config

// Config is a interface to get the config of a given implementation.
type Config interface {
	// Get will get all the ConfigInfo available in the implementation
	Get() ConfigInfo
}

// ConfigInfo is the common	 structure to contain all the config fields.
type ConfigInfo struct {
	Server               server               `yaml:"server"`
	PostgreSQLProperties postgreSQLProperties `yaml:"psql"`
}

type server struct {
	Port           int      `yaml:"port"`
	Host           string   `yaml:"host"`
	AllowedOrigins []string `yaml:"allowed_origins"`
	SecretKey      string   `yaml:"secret_key"`
	JWTLifespan    int      `yaml:"jwt_lifespan"`
}

type postgreSQLProperties struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
