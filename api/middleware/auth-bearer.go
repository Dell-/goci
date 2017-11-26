package middleware

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"strings"
	"github.com/Dell-/goci/models"
	"time"
)

func BearerAuthenticate(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	token := GetAccessToken(request)

	// If the token is empty...
	if token == "" {
		// If we get here, the required token is missing
		response.WriteErrorString(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	if !validateToken(token) {
		response.WriteErrorString(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	chain.ProcessFilter(request, response)
}

func validateToken(token string) bool {
	accessToken := models.GetAccessTokenBySHA(token)
	return accessToken != nil && accessToken.ExpiredUnix > time.Now().Unix()
}

// Get token from the Authorization header
// format: Authorization: Bearer
func GetAccessToken(req *restful.Request) string {
	// Get token from the Authorization header
	// format: Authorization: Bearer
	token := req.Request.Header.Get("Authorization")
	if len(token) > 0 {
		return strings.TrimPrefix(token, "Bearer ")
	}

	return ""
}
