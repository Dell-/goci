package cmd

import (
	"github.com/urfave/cli"
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
)

var Rest = cli.Command{
	Name:        "rest",
	Usage:       "Start rest service",
	Description: "",
	Action:      runRest,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "3000",
			Usage: "Temporary port number to prevent conflict",
		},
	},
}

type Token struct {
	access_token string
}

type TokenResource struct {
}

func init()  {

}

func createToken(req *restful.Request, resp *restful.Response) {
	login, err := req.BodyParameter("login")

	if err != nil || login == "" {
		log.Println("Error: empty login")
		return
	}

	pass, err := req.BodyParameter("password")

	if err != nil || pass == "" {
		log.Println("Error: empty password")
		return
	}

	log.Println("Create token for > " + login + "|" + pass)
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// usr/pwd = admin/admin
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "admin" || p != "admin" {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "")
		return
	}
	chain.ProcessFilter(req, resp)
}

func (p TokenResource) Register(container *restful.Container) {

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/tokens").Filter(basicAuthenticate).To(createToken))
	container.Add(ws)
}

func runRest(ctx *cli.Context) error {

	t := TokenResource{}
	wsContainer := restful.NewContainer()
	t.Register(wsContainer)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"POST"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	log.Print("Start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}

	log.Fatal(server.ListenAndServe())

	return nil
}
