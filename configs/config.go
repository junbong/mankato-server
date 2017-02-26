package configs

type Config struct {
	Foo        string  `yaml:"foo"`
	Server     struct {
		           Host string `yaml:"host"`
		           Port int    `yaml:"port"`
	           }
	Collection struct {
		           TtlScanPeriodMillis int `yaml:"ttl_scan_period"`
	           }
}
