package internal

import (
	"errors"
)

var (
	ErrIsNotInstalled         = errors.New("application is not installed")
	ErrAlreadyInstalled       = errors.New("application is already installed")
	ErrIsUpToDate             = errors.New("application is up to date")
	ErrAlreadyInitialized     = errors.New("already initialized! use --force flag")
	ErrNotInitialized         = errors.New("app manager is not initialized. See documentation")
	ErrRepositoryNotExist     = errors.New("repository is not exist")
	ErrRepositoryAlreadyExist = errors.New("repository is already exist")
)
