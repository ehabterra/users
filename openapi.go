package usersapi

import (
	"log"
	openapi "users/gen/openapi"
)

const (
	roleBucket string = "ROLE"
	userBucket string = "USER"
)

// openapi service example implementation.
// The example methods log the requests and return zero values.
type openapisrvc struct {
	logger *log.Logger
}

// NewOpenapi returns the openapi service implementation.
func NewOpenapi(logger *log.Logger) openapi.Service {
	return &openapisrvc{logger}
}
