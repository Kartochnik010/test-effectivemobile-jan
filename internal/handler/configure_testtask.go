// This file is safe to edit. Once it exists it will not be overwritten

package handler

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Kartochnik010/test-effectivemobile-jan/internal/handler/operations"
)

//go:generate swagger generate server --target ../../../test-effectivemobile-jan --name Testtask --spec ../../api/swagger.yaml --server-package internal/handler --principal interface{}

func configureFlags(api *operations.TesttaskAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TesttaskAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DeletePeopleIDHandler == nil {
		api.DeletePeopleIDHandler = operations.DeletePeopleIDHandlerFunc(func(params operations.DeletePeopleIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeletePeopleID has not yet been implemented")
		})
	}
	if api.GetPeopleHandler == nil {
		api.GetPeopleHandler = operations.GetPeopleHandlerFunc(func(params operations.GetPeopleParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPeople has not yet been implemented")
		})
	}
	if api.GetPeopleIDHandler == nil {
		api.GetPeopleIDHandler = operations.GetPeopleIDHandlerFunc(func(params operations.GetPeopleIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPeopleID has not yet been implemented")
		})
	}
	if api.PostPeopleHandler == nil {
		api.PostPeopleHandler = operations.PostPeopleHandlerFunc(func(params operations.PostPeopleParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostPeople has not yet been implemented")
		})
	}
	if api.PutPeopleIDHandler == nil {
		api.PutPeopleIDHandler = operations.PutPeopleIDHandlerFunc(func(params operations.PutPeopleIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutPeopleID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
