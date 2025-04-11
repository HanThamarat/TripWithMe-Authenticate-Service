package core

type UserService interface {
	Save(user User) (*User, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
};

func (s *userServiceImpl) Save(user User) (*User, error) {
	savedUser, err := s.repo.Save(user);

	if err != nil {
		return nil, err
	}

	return savedUser, nil;
}