package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsersFromBd()
}

// CreateUser inserts a new user.
func (s *Service) CreateUser(user *User) error {
	return s.repo.CreateUserFromBd(user)
}

// GetUserByID retrieves a user by ID.
func (s *Service) GetUserByID(id int) (*User, error) {
	return s.repo.GetUserByIDFromBd(id)
}

// DeleteUser deletes a user by ID.
func (s *Service) DeleteUser(id int) error {
	return s.repo.DeleteUserFromBd(id)
}

// UpdateUser updates a user with new values.
func (s *Service) UpdateUser(id int, updates *User) error {
	return s.repo.UpdateUserFromBd(id, updates)
}

// func GetTasksForUser(userID uint) ([]models.Task, error) {
// 	return GetTasksForUserFromBd(userID)
// }
