package config

type ServerConfig struct {
	HttpPort    string `yaml:"http_port"`
	GrpcPort    string `yaml:"grpc_port"`
	ServiceName string `yaml:"service_name"`
}
