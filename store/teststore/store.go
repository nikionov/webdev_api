package teststoretype

import (
	"github.com/nikionov/job-backend-trainee/internal/app/store/sqlstore"
	"github.com/nikionov/job-backend-trainee/model"
)
Store struct {
userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() sqlstore.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
