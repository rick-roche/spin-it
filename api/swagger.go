package api

import (
	"net/http"
	"path/filepath"
)

func (a *API) initialiseSwagger() {
	usp, _ := filepath.Abs("./third_party/swaggerui")
	ui := http.FileServer(http.Dir(usp))
	a.Router.Handle("/*", ui)

	jsp, _ := filepath.Abs("./api/swagger/")
	js := http.FileServer(http.Dir(jsp))
	a.Router.Handle("/swagger.json", js)
}
