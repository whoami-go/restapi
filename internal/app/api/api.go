package api

type API struct {
}

func New() *API {
	return &API{}
}
func (api *API) Start() error {
	return nil
}
