package config

func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		ServerHost: "0.0.0.0",
		ServerPort: 9000,
	}
}
