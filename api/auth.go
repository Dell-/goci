package api

import (
	"net/http"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/Dell-/goci/models"
	"github.com/Dell-/goci/api/errors"
	"github.com/Dell-/goci/api/middleware"
)

// AuthResource is the REST layer to the Auth domain
type AuthResource struct {
	// normally one would use DAO (data access object)
}

// WebService creates a new service that can handle REST requests for User resources.
func (u AuthResource) WebService() *restful.WebService {
	webService := new(restful.WebService)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"POST", "GET"},
		CookiesAllowed: false}
	webService.Filter(cors.Filter)
	webService.
	Path("/auth").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"auth"}

	webService.Route(webService.POST("/login").To(u.loginUser).
	// docs
		Doc("Auth user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(LoginData{})) // from the request

	webService.Route(webService.POST("/logout").
		Filter(middleware.BearerAuthenticate).
		To(u.logoutUser).
	// docs
		Doc("Logout current user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(LoginData{})) // from the request

	return webService
}

// POST /auth/login
// {"email": "test@test.loc", "password": "123123q"}
//
func (u *AuthResource) loginUser(request *restful.Request, response *restful.Response) {
	data := new(LoginData)
	err := request.ReadEntity(&data)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusUnprocessableEntity, errors.LoginError())
		return
	}
	user := models.GetUserByEmail(data.Email)

	if user == nil {
		response.WriteHeaderAndEntity(http.StatusUnprocessableEntity, errors.LoginError())
		return
	}

	if !user.CheckPassword(data.Password) {
		response.WriteHeaderAndEntity(http.StatusUnprocessableEntity, errors.LoginError())
		return
	}

	token := &models.AccessToken{
		UID: user.ID,
	}

	if models.DeleteAccessTokenOfUserByID(user.ID) != nil || models.NewAccessToken(token) != nil {
		response.WriteHeaderAndEntity(http.StatusUnprocessableEntity, errors.LoginError())
		return
	}

	response.WriteHeaderAndEntity(http.StatusCreated, UserToken{Token: token.Token, Expired: token.ExpiredUnix})
}

// POST /auth/logout
func (u *AuthResource) logoutUser(request *restful.Request, response *restful.Response) {
	token := middleware.GetAccessToken(request)
	err := models.DeleteAccessTokenOfUser(token)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, errors.LogoutError())
		return
	}

	response.WriteHeader(http.StatusOK)
}

type LoginData struct {
	Email    string `json:"email" description:"User email"`
	Password string `json:"password" description:"User Password"`
}

type UserToken struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
