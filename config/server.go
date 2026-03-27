package config

// ServerList :
type ServerList struct {
	Rest Server
}

// Server is struct for server conf
type Server struct {
	TLS             bool `mapstructure:"tls"`
	Name            string
	Host            string
	Port            int
	Path            string
	SecretKey       string
	Timeout         int
	ApplicationCode string
	Env             string
}

type MessageBroker struct {
	Adapter string
	Host    string
	Port    string
}
