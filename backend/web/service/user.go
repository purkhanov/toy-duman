package service

type UserService struct{}

func (s *UserService) GetUser() (string, error) {
	// fmt.Println("GetUser called")
	return "Nurlan", nil
}
