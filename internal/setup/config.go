package setup

import (
	"sync"

	"github.com/spf13/viper"
)

var configOnce *sync.Once = &sync.Once{}
var config *viper.Viper = viper.New()
var currentConfig *Config

type Config struct {
	Exchange      string
	RoutingKey    string
	QueueName     string
	BrokerUrl     string
	ZincSearchUrl string
	ZincIndex     string
	ZincUser      string
	ZincPass      string
	LogLevel      int
	Version       string
}

const (
	Exchange      = "EXCHANGE_NAME"
	RoutingKey    = "ROUTING_KEY"
	QueueName     = "QUEUE"
	BrokerUrl     = "BROKER_URL"
	ZincSearchUrl = "URL"
	ZincIndex     = "INDEX"
	ZincUser      = "USER"
	ZincPass      = "PASS"
	LogLevel      = "LOG_LEVEL"
	Version       = "VERSION"
)

func Get() *Config {
	configOnce.Do(func() {
		config.SetEnvPrefix("ZINC")
		config.BindEnv(Exchange)
		config.BindEnv(RoutingKey)
		config.BindEnv(QueueName)
		config.BindEnv(BrokerUrl)
		config.BindEnv(ZincSearchUrl)
		config.BindEnv(ZincIndex)
		config.BindEnv(ZincUser)
		config.BindEnv(ZincPass)
		config.BindEnv(LogLevel)
		config.BindEnv(Version)

		currentConfig = &Config{
			Exchange:      config.GetString(Exchange),
			RoutingKey:    config.GetString(RoutingKey),
			QueueName:     config.GetString(QueueName),
			BrokerUrl:     config.GetString(BrokerUrl),
			ZincSearchUrl: config.GetString(ZincSearchUrl),
			ZincIndex:     config.GetString(ZincIndex),
			ZincUser:      config.GetString(ZincUser),
			ZincPass:      config.GetString(ZincPass),
			LogLevel:      config.GetInt(LogLevel),
			Version:       config.GetString(Version),
		}
	})

	return currentConfig
}
