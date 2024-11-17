package cenvvm

import (
	"errors"
)

var (
	ErrNoCoreServiceRunning = errors.New("no Core service instance was found on the host")
	ErrSocketConnUpgrade    = errors.New("the socket connection could not be upgraded")
	ErrEmptyBasePath        = errors.New("basePath cannot be empty")
	ErrSocketNotFound       = errors.New("socket file not found")
	ErrInvalidSocket        = errors.New("invalid socket path")
	ErrPermission           = errors.New("permission denied")
)
