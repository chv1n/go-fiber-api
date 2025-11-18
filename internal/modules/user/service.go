package user

type Service interface {
	GetUsers() ([]User, error)
	CreateUser(u *User) error
}

type userService struct {
	repo Repository
}

func NewUserService(r Repository) Service {
	return &userService{r}
}

func (s *userService) GetUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *userService) CreateUser(u *User) error {
	return s.repo.Create(u)
}
