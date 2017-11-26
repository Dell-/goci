package cmd

import (
	"github.com/urfave/cli"
	"github.com/emicklei/go-restful"
	"net/http"
	log "github.com/go-clog/clog"
	"github.com/Dell-/goci/pkg/settings"
	"github.com/Dell-/goci/pkg/global"
	"github.com/Dell-/goci/api"
	"strconv"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"path"
)

var Rest = cli.Command{
	Name:        "rest",
	Usage:       "Start rest service",
	Description: "",
	Action:      runRest,
}

func runRest(ctx *cli.Context) error {
	global.Initialize()

	wsContainer := restful.NewContainer()

	// Rest endpoints
	auth := api.AuthResource{}
	users := api.UsersResource{}
	wsContainer.Add(auth.WebService())
	wsContainer.Add(users.WebService())

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		WebServicesURL:                settings.SERVER.ApiUrl,
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	wsContainer.Add(restfulspec.NewOpenAPIService(config))

	Addr := settings.SERVER.HTTPAddr + ":" + strconv.Itoa(settings.SERVER.HTTPPort)
	log.Info("Start listening on " + Addr)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir(path.Join(settings.SERVER.StaticRootPath, "dist")))))

	server := &http.Server{Addr: Addr, Handler: wsContainer}

	return server.ListenAndServe()
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Goci",
			Description: "API Doc",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "0.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}
