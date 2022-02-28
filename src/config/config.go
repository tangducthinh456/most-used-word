package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CONFIG_FILE_NAME       = "config"          // const string file name of config file
	CONFIG_FILE_TYPE       = "yaml"            // const string file type of config file
	CONFIG_SERVER_HOST     = "server.Host"     // viper get value from field's name and struct's name of file yaml
	CONFIG_SERVER_PORT     = "server.Port"     // viper get value from field's name and struct's name of file yaml
	CONFIG_SERVER_PROTOCOL = "server.Protocol" // viper get value from field's name and struct's name of file yaml
)

var (
	serverConf   *ServerConfig // variable represents server config
	readerConfig *viper.Viper  // variable represents viper which read value from file yaml
)

type ServerConfig struct {
	ServerHost string // host ip of server
	ServerPort string // port of server
	Protocol   string // http or https
}

// init() is a special function called when another package import this config package
// This config package is implement in Singleton pattern, which has private constructor and
// a public method to access object, guarantee there is one object created in all program
func init() {

	if readerConfig == nil {
		readerConfig = viper.New()
		readerConfig.SetConfigName(CONFIG_FILE_NAME)
		readerConfig.SetConfigType(CONFIG_FILE_TYPE)
		readerConfig.AddConfigPath(".")
		readerConfig.AddConfigPath("../../..")

		if err := readerConfig.ReadInConfig(); err != nil {
			log.Fatalf("Can not read the config file : %+v", err)
		}
	}

	if serverConf == nil {
		serverConf = &ServerConfig{
			ServerHost: readerConfig.GetString(CONFIG_SERVER_HOST),
			ServerPort: readerConfig.GetString(CONFIG_SERVER_PORT),
			Protocol:   readerConfig.GetString(CONFIG_SERVER_PROTOCOL),
		}
	}
}

// Getter() of server config
func GetServerConfig() *ServerConfig {
	return serverConf
}

// Setter() of server config
func SetServerConfig(serv *ServerConfig) {
	serverConf = serv
}
