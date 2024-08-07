package config

type IntegrationConfig struct {
	CoreService    ServiceConfig `yaml:"core_service"`
	MonitorService ServiceConfig `yaml:"monitor_service"`
	JobService     ServiceConfig `yaml:"job_service"`
}

type ServiceConfig struct {
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}
