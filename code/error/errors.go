package error

import (
	"errors"
	"monitor/config"
	"monitor/model"
	"net/http"
)

// our ready errors we always define the errors here
//the names are explanatory

var ErrMissingFields = errors.New("some fields are missing")
var ErrMaxNumLink = errors.New("maximum number of links have been reached")

var ErrNotFound = model.ErrResponse{
	Status:    "failed",
	ErrorMsg:  "the resource you are looking for cannot be found",
	ErrorCode: http.StatusNotFound,
	Help:      config.Help,
}

var ErrBadRequestBadFields = model.ErrResponse{
	Status:    "failed",
	ErrorMsg:  "some of the required fields are wrong or missing",
	ErrorCode: http.StatusBadRequest,
	Help:      config.Help,
}

var ErrInternal = model.ErrResponse{
	Status:    "failed",
	ErrorMsg:  "something went wrong with our service",
	ErrorCode: http.StatusInternalServerError,
	Help:      config.Help,
}

var ErrNoPermission = model.ErrResponse{
	Status:    "failed",
	ErrorMsg:  "not enough permission",
	ErrorCode: http.StatusForbidden,
	Help:      config.Help,
}
