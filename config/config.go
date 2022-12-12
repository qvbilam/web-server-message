package config

type ServerConfig struct {
	Name                 string               `mapstructure:"name" json:"name"`
	Host                 string               `mapstructure:"host" json:"host"`
	Port                 int64                `mapstructure:"port" json:"port"`
	Tags                 []string             `mapstructure:"tags" json:"tags"`
	UserServerConfig     UserServerConfig     `mapstructure:"user-server" json:"user-server"`
	MessageServerConfig  MessageServerConfig  `mapstructure:"message-server" json:"message-server"`
	RabbitMQServerConfig RabbitMQServerConfig `mapstructure:"rabbit-server" json:"rabbit-server"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type MessageServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type RabbitMQServerConfig struct {
	Host                string `mapstructure:"host" json:"host"`
	Port                int64  `mapstructure:"port" json:"port"`
	Name                string `mapstructure:"name" json:"name"`
	QueueSuffix         string `mapstructure:"queue_suffix" json:"queue_suffix"`
	MessageExchangeName string
	MessageQueueName    string
}
