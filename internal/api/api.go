// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /login)
	GetLogin(c *gin.Context)

	// (POST /login)
	PostLogin(c *gin.Context)

	// (POST /password/{token})
	PostPasswordToken(c *gin.Context, token string)

	// (POST /space/)
	PostSpace(c *gin.Context)

	// (DELETE /space/{space-id})
	DeleteSpaceSpaceId(c *gin.Context, spaceId int)

	// (GET /space/{space-id})
	GetSpaceSpaceId(c *gin.Context, spaceId int)

	// (PUT /space/{space-id})
	PutSpaceSpaceId(c *gin.Context, spaceId int)

	// (GET /spaces/)
	GetSpaces(c *gin.Context)

	// (POST /user/)
	PostUser(c *gin.Context)

	// (DELETE /user/{user-id})
	DeleteUserUserId(c *gin.Context, userId int)

	// (GET /user/{user-id})
	GetUserUserId(c *gin.Context, userId int)

	// (PUT /user/{user-id})
	PutUserUserId(c *gin.Context, userId int)

	// (GET /users/)
	GetUsers(c *gin.Context)

	// (GET /webhooks/)
	GetWebhooks(c *gin.Context)

	// (POST /webhooks/)
	PostWebhooks(c *gin.Context)

	// (PUT /webhooks/{name})
	PutWebhooksName(c *gin.Context, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetLogin operation middleware
func (siw *ServerInterfaceWrapper) GetLogin(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetLogin(c)
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostLogin(c)
}

// PostPasswordToken operation middleware
func (siw *ServerInterfaceWrapper) PostPasswordToken(c *gin.Context) {

	var err error

	// ------------- Path parameter "token" -------------
	var token string

	err = runtime.BindStyledParameter("simple", false, "token", c.Param("token"), &token)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter token: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostPasswordToken(c, token)
}

// PostSpace operation middleware
func (siw *ServerInterfaceWrapper) PostSpace(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSpace(c)
}

// DeleteSpaceSpaceId operation middleware
func (siw *ServerInterfaceWrapper) DeleteSpaceSpaceId(c *gin.Context) {

	var err error

	// ------------- Path parameter "space-id" -------------
	var spaceId int

	err = runtime.BindStyledParameter("simple", false, "space-id", c.Param("space-id"), &spaceId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter space-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteSpaceSpaceId(c, spaceId)
}

// GetSpaceSpaceId operation middleware
func (siw *ServerInterfaceWrapper) GetSpaceSpaceId(c *gin.Context) {

	var err error

	// ------------- Path parameter "space-id" -------------
	var spaceId int

	err = runtime.BindStyledParameter("simple", false, "space-id", c.Param("space-id"), &spaceId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter space-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSpaceSpaceId(c, spaceId)
}

// PutSpaceSpaceId operation middleware
func (siw *ServerInterfaceWrapper) PutSpaceSpaceId(c *gin.Context) {

	var err error

	// ------------- Path parameter "space-id" -------------
	var spaceId int

	err = runtime.BindStyledParameter("simple", false, "space-id", c.Param("space-id"), &spaceId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter space-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutSpaceSpaceId(c, spaceId)
}

// GetSpaces operation middleware
func (siw *ServerInterfaceWrapper) GetSpaces(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSpaces(c)
}

// PostUser operation middleware
func (siw *ServerInterfaceWrapper) PostUser(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUser(c)
}

// DeleteUserUserId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user-id" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "user-id", c.Param("user-id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserUserId(c, userId)
}

// GetUserUserId operation middleware
func (siw *ServerInterfaceWrapper) GetUserUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user-id" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "user-id", c.Param("user-id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserUserId(c, userId)
}

// PutUserUserId operation middleware
func (siw *ServerInterfaceWrapper) PutUserUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user-id" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "user-id", c.Param("user-id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user-id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutUserUserId(c, userId)
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsers(c)
}

// GetWebhooks operation middleware
func (siw *ServerInterfaceWrapper) GetWebhooks(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetWebhooks(c)
}

// PostWebhooks operation middleware
func (siw *ServerInterfaceWrapper) PostWebhooks(c *gin.Context) {

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostWebhooks(c)
}

// PutWebhooksName operation middleware
func (siw *ServerInterfaceWrapper) PutWebhooksName(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", c.Param("name"), &name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(CookieAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutWebhooksName(c, name)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/login", wrapper.GetLogin)
	router.POST(options.BaseURL+"/login", wrapper.PostLogin)
	router.POST(options.BaseURL+"/password/:token", wrapper.PostPasswordToken)
	router.POST(options.BaseURL+"/space/", wrapper.PostSpace)
	router.DELETE(options.BaseURL+"/space/:space-id", wrapper.DeleteSpaceSpaceId)
	router.GET(options.BaseURL+"/space/:space-id", wrapper.GetSpaceSpaceId)
	router.PUT(options.BaseURL+"/space/:space-id", wrapper.PutSpaceSpaceId)
	router.GET(options.BaseURL+"/spaces/", wrapper.GetSpaces)
	router.POST(options.BaseURL+"/user/", wrapper.PostUser)
	router.DELETE(options.BaseURL+"/user/:user-id", wrapper.DeleteUserUserId)
	router.GET(options.BaseURL+"/user/:user-id", wrapper.GetUserUserId)
	router.PUT(options.BaseURL+"/user/:user-id", wrapper.PutUserUserId)
	router.GET(options.BaseURL+"/users/", wrapper.GetUsers)
	router.GET(options.BaseURL+"/webhooks/", wrapper.GetWebhooks)
	router.POST(options.BaseURL+"/webhooks/", wrapper.PostWebhooks)
	router.PUT(options.BaseURL+"/webhooks/:name", wrapper.PutWebhooksName)
}
