#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <module_name>"
  exit 1
fi

MODULE=$1
MODULE_PATH="internal/modules/$MODULE"
MODULE_UPPER="$(tr '[:lower:]' '[:upper:]' <<< ${MODULE:0:1})${MODULE:1}" # This will be the upper cased module name

mkdir -p "$MODULE_PATH"

cat > "$MODULE_PATH/model.go" <<EOF
package $MODULE

import (
    "time"
)

// TODO: Define your $MODULE_UPPER model(s) here
type $MODULE_UPPER struct {
    // Add fields for the $MODULE_UPPER model
    ID   int    \`json:"id"\`
    CreatedAt     time.Time \`gorm:"autoCreateTime"\`
	UpdatedAt     time.Time \`gorm:"autoUpdateTime"\`
}
EOF

cat > "$MODULE_PATH/repository.go" <<EOF
package $MODULE

import (
    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

// NewRepository creates a new repository for $MODULE
func NewRepository(db *gorm.DB) *Repository {
    return &Repository{DB: db}
}

// TODO: Implement repository methods for $MODULE
// Create$MODULE_UPPER creates a new $MODULE_UPPER record in the database
func (repository *Repository) Create$MODULE_UPPER(params *$MODULE_UPPER) (*$MODULE_UPPER, error) {
    $MODULE := $MODULE_UPPER{
        // Initialize fields if necessary
    }
    
    if err := repository.DB.Create(&$MODULE).Error; err != nil {
        return nil, err
    }

    return &$MODULE, nil
}

// Get$MODULE_UPPER retrieves a $MODULE_UPPER record by ID
func (repository *Repository) Get$MODULE_UPPER(id int) (*$MODULE_UPPER, error) {
    $MODULE := $MODULE_UPPER{}
    if err := repository.DB.First(&$MODULE, id).Error; err != nil {
        return nil, err
    }
    return &$MODULE, nil
}

// Get all $MODULE_UPPER list retrieves all $MODULE_UPPER records and returns paginated results


// Update$MODULE_UPPER updates an existing $MODULE_UPPER record
func (repository *Repository) Update$MODULE_UPPER(id int, params *$MODULE_UPPER) (*$MODULE_UPPER, error) {
    $MODULE := $MODULE_UPPER{}
    if err := repository.DB.First(&$MODULE, id).Error; err != nil {
        return nil, err
    }
    // Update fields as necessary
    if err := repository.DB.Save(&$MODULE).Error; err != nil {
        return nil, err
    }
    return &$MODULE, nil
}

EOF

cat > "$MODULE_PATH/service.go" <<EOF
package $MODULE

import (
    "calorie_deficit/internal/utils"

    "errors"

    "gorm.io/gorm"
)

type Service struct {
    repository *Repository
}

func NewService(repository *Repository) *Service {
    return &Service{repository: repository}
}

// TODO: Implement service logic for $MODULE
// Create$MODULE_UPPER creates a new $MODULE_UPPER record
func (service *Service) Create$MODULE_UPPER(params *$MODULE_UPPER) (*$MODULE_UPPER, error) {
    createResponse, err := service.repository.Create$MODULE_UPPER(params)
    if err != nil {
        return nil, err
    }
    return createResponse, nil
}

// Get$MODULE_UPPER retrieves a $MODULE_UPPER record by ID
func (service *Service) Get$MODULE_UPPER(id int) (*$MODULE_UPPER, error) {
    response, err := service.repository.Get$MODULE_UPPER(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, utils.HandleNotFoundError(err)
        }
        return nil, err
    }
    return response, nil
}

EOF

cat > "$MODULE_PATH/handler.go" <<EOF
package $MODULE

import (
    "calorie_deficit/internal/constants"
    "calorie_deficit/internal/types"
    //"github.com/gin-gonic/gin"
)

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

var (
	// ErrInvalidRequest is an error for invalid request
	errInvalidRequest = types.ErrorResponse{
		Code:    400,
		Message: constants.LogMessages.General.InvalidRequest,
	}
)

// TODO: Implement HTTP handlers for $MODULE
// Create$MODULE_UPPER handles the creation of a new $MODULE_UPPER record

// Get$MODULE_UPPER handles the retrieval of a $MODULE_UPPER record by ID

// Get${MODULE_UPPER}List handles the retrieval of all $MODULE_UPPER records

// Update$MODULE_UPPER handles the update of an existing $MODULE_UPPER record

EOF

echo "Module '$MODULE' created at $MODULE_PATH"