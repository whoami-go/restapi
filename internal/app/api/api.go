package api

// base api server instance description
type API struct {
	//unexported field!
	config *Config
}

// api constructor : build base api instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// start http server / configure logger , router and etc
func (api *API) Start() error {
	return nil
}
