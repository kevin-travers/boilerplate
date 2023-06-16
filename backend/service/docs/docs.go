// Package classification awesome.
//
// Documentation of our awesome API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: some-url.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta

package docs

import (
	"service/models"

	_ "github.com/pdrum/swagger-automation/api"
)

// swagger:route GET /job
// Get job by id
// responses:
//   200: JobRequest

// This text will appear as description of your response body.
// swagger:response JobRequest
type JobRequestWrapper struct {
	// in:body
	Body models.JobRequest
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body models.JobRequest
}
