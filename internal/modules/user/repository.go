package user

import (
    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

// NewRepository creates a new repository for user
func NewRepository(db *gorm.DB) *Repository {
    return &Repository{DB: db}
}

// TODO: Implement repository methods for user
// CreateUser creates a new User record in the database
func (repository *Repository) CreateUser(params *User) (*User, error) {
    user := User{
        // Initialize fields if necessary
    }
    
    if err := repository.DB.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

// GetUser retrieves a User record by ID
func (repository *Repository) GetUser(id int) (*User, error) {
    user := User{}
    if err := repository.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Get all User list retrieves all User records and returns paginated results


// UpdateUser updates an existing User record
func (repository *Repository) UpdateUser(id int, params *User) (*User, error) {
    user := User{}
    if err := repository.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    // Update fields as necessary
    if err := repository.DB.Save(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

