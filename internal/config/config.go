package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/thediveo/enumflag/v2"
)

const (
	json = iota
	text
)

var (
	configPath       string
	httpAddress      string
	httpTimeoutRead  int
	httpTimeoutWrite int
	logLevel         slog.Level
	logFormat        int
	databaseHost     string
	databasePort     int
	databaseName     string
	databaseUser     string
	//databasePassword string //commented while unused
)

var logLevelIds = map[slog.Level][]string{
	slog.LevelDebug: {"debug"},
	slog.LevelInfo:  {"info"},
	slog.LevelWarn:  {"warning", "warn"},
	slog.LevelError: {"error"},
}

var logFormatIds = map[int][]string{
	json: {"json"},
	text: {"text"},
}

type Config struct {
	HTTP     HTTP     `mapstructure:"http-config"`
	Logger   Log      `mapstructure:"log"`
	Database Database `mapstructure:"database"`
	Version  string
	Build    string
}

type HTTP struct {
	Address string      `mapstructure:"address"`
	Timeout HTTPTimeout `mapstructure:"timeout"`
}

type HTTPTimeout struct {
	Read  int `mapstructure:"read"`
	Write int `mapstructure:"write"`
}

type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func New(version, build string) (*Config, error) {
	// Set enviroment prefix and bind to viper.
	viper.SetEnvPrefix("EGRA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Scan CLI flags and bind to viper.
	pflag.StringVarP(&configPath, "config", "c", viper.GetViper().ConfigFileUsed(), "path to config file")
	pflag.StringVarP(&httpAddress, "http.address", "a", "127.0.0.1:8080", "http listening address")
	pflag.IntVarP(&httpTimeoutRead, "http.timeout.read", "t", 5, "http read timeout")
	pflag.IntVarP(&httpTimeoutWrite, "http.timeout.write", "w", 5, "http write timeout")
	pflag.VarP(
		enumflag.New(&logLevel, "log.level", logLevelIds, enumflag.EnumCaseSensitive),
		"log.level", "l", "log level",
	)
	pflag.VarP(
		enumflag.New(&logFormat, "log.format", logFormatIds, enumflag.EnumCaseSensitive),
		"log.format", "f", "log format",
	)
	pflag.StringVarP(&databaseHost, "database.host", "H", "localhost", "database host")
	pflag.IntVarP(&databasePort, "database.port", "P", 5432, "database port")
	pflag.StringVarP(&databaseName, "database.name", "N", "gapi", "database name")
	pflag.StringVarP(&databaseUser, "database.user", "U", "postgres", "database user")
	pflag.Parse()

	if err := viper.BindPFlag("http.address", pflag.Lookup("http.address")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("http.timeout.read", pflag.Lookup("http.timeout.read")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("http.timeout.write", pflag.Lookup("http.timeout.write")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("log.level", pflag.Lookup("log.level")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("log.format", pflag.Lookup("log.format")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("database.host", pflag.Lookup("database.host")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("database.port", pflag.Lookup("database.port")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("database.name", pflag.Lookup("database.name")); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag("database.user", pflag.Lookup("database.user")); err != nil {
		return nil, err
	}

	if configPath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configPath)
	} else {
		// Set config name, path and type.
		viper.AddConfigPath("configs")
		viper.AddConfigPath("/etc/example-gorilla-rest-api")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// Discover and load the configuration file from disk.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// unmarshal unmarshals the config into a Struct
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.Version = version
	cfg.Build = build

	return &cfg, nil
}
