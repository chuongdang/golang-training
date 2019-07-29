package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const TOKEN_ADMIN = "admin"
const ROLE_ADMIN = "admin"

const TOKEN_USER = "user"
const ROLE_USER = "user"

type MapTokenRole map[string]string

var mapTokenRole MapTokenRole

type MapResourceRole map[string]MapTokenRole

var mapResourceRole MapResourceRole

func init() {
	mapTokenRole = MapTokenRole{
		TOKEN_ADMIN: ROLE_ADMIN,
		TOKEN_USER: ROLE_USER,
	}

	mapResourceRole = MapResourceRole{
		"POST /user": MapTokenRole{
			ROLE_ADMIN: ROLE_ADMIN,
		},
		"GET /user": MapTokenRole{
			ROLE_ADMIN: ROLE_ADMIN,
			ROLE_USER: ROLE_USER,
		},
	}
}

func Auth(c *gin.Context) {
	tokenHeader := c.Request.Header.Get("X-Nordic-Token")
	if token := mapTokenRole[tokenHeader]; token == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Wrong token, please try again"))
		return
	}

	role := mapTokenRole[tokenHeader]
	method := strings.ToUpper(c.Request.Method)
	route := strings.ToLower(c.Request.URL.Path)

	methodAndRoute := method + " " + route

	if valid := mapResourceRole[methodAndRoute]; valid == nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("You could not access this path"))
		return
	}

	allowedRoleList := mapResourceRole[methodAndRoute]
	if a := allowedRoleList[role]; a == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Please ask admin for permission"))
	}

	c.Next()
}