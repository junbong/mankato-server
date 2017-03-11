package configs

type Config struct {
	Server            struct {
		                  Host              string `yaml:"host"`
		                  Port              int    `yaml:"port"`
		                  ClientMaxBodySize string `yaml:"client_max_body_size"`
	                  }
	Collection        struct {
		                  DefaultName         string `yaml:"default_name"`
		                  TtlScanPeriodMillis int `yaml:"ttl_scan_period"`
	                  }
	ExpirationHandler []ExpHandlerConfigs `yaml:"expiration_handler"`
}

type ExpHandlerConfigs struct {
	Type       string `yaml:"type"`
	Enable     bool `yaml:"enable"`
	Properties map[string]string `yaml:"properties"`
}
