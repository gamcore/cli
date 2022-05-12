package api

import (
	"errors"
)

var (
	// Core schema
	ErrIsNotInstalled     = errors.New("application is not installed")
	ErrAlreadyInstalled   = errors.New("application is already installed")
	ErrIsUpToDate         = errors.New("application is up to date")
	ErrAlreadyInitialized = errors.New("already initialized! use --force flag")
	ErrNotInitialized     = errors.New("app manager is not initialized. See documentation")
	ErrRepoNotExist       = errors.New("repository is not exist")
	ErrRepoAlreadyExist   = errors.New("repository is already exist")
	ErrNoUpdates          = errors.New("no updates")
	// manifest / schema
	ErrManifestUrlEmpty           = errors.New("manifest: empty url somewhere")
	ErrManifestDescriptionEmpty   = errors.New("manifest: description is empty")
	ErrManifestRequiredAnyExec    = errors.New("manifest: must defined one execute path")
	ErrManifestLicenseNameEmpty   = errors.New("manifest: empty license name")
	ErrManifestUrlPatternRequired = errors.New("manifest: must provide url pattern")
	ErrSchemaIsNotSupportedYet    = errors.New("manifest: this schema is not supported yet")
	ErrSchemaInvalidType          = errors.New("manifest: invalid update schema type")
	ErrSchemaInvalidURL           = errors.New("manifest: invalid github repository url")
	ErrSchemaInvalidPattern       = errors.New("manifest: invalid version pattern")
	ErrSchemaInvalidPath          = errors.New("manifest: invalid path or value is not a string")
	ErrSchemaPathRequired         = errors.New("manifest: path is required")
	ErrSchemaPatternRequired      = errors.New("manifest: regex pattern is required")
	// download / extract
	ErrArchiveInvalidPath = errors.New("extract: invalid file path")
	ErrUnsupportedArchive = errors.New("extract: unsupported archive")
	// inputs
	ErrInputAlreadyExist = errors.New("already exist")
	ErrInputEmpty        = errors.New("must be not empty")
)
