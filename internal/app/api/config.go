package api

// general instance for api server of REST app
type Config struct {
	//PORT
	BindAddr string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
