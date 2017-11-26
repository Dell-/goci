package api

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/Dell-/goci/api/middleware"
)

// UserResource is the REST layer to the Auth domain
type UsersResource struct {
	// normally one would use DAO (data access object)
}

// WebService creates a new service that can handle REST requests for User resources.
func (u UsersResource) WebService() *restful.WebService {
	webService := new(restful.WebService)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET"},
		CookiesAllowed: false}
	webService.Filter(cors.Filter).Filter(middleware.BearerAuthenticate)
	webService.
	Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"users"}

	webService.Route(webService.GET("/current").To(u.currentUser).
	// docs
		Doc("Current user").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	return webService
}

// GET /users/current
// {"username": "test_test", "email": "test@test.loc", "role": "user"}
//
func (u *UsersResource) currentUser(request *restful.Request, response *restful.Response) {
	response.WriteEntity(new(User))
}

type User struct {
	Username string `json:"username" description:"User name"`
	Email string `json:"email" description:"User email"`
	Role string `json:"role" description:"User role"`
}