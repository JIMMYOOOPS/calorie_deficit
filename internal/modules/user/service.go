package user

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

// TODO: Implement service logic for user
// CreateUser creates a new User record
func (service *Service) CreateUser(params *User) (*User, error) {
    createResponse, err := service.repository.CreateUser(params)
    if err != nil {
        return nil, err
    }
    return createResponse, nil
}

// GetUser retrieves a User record by ID
func (service *Service) GetUser(id int) (*User, error) {
    response, err := service.repository.GetUser(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, utils.HandleNotFoundError(err)
        }
        return nil, err
    }
    return response, nil
}

