package simpleSite

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Save(*User) error
	GetUser(string) (*User, error)
}

func (svc *ServiceImpl) Save(u *User) error {
	user, _ := svc.Repository.FindUserByUsername(u.Username)

	if user.Id == 0 {
		id, err := svc.Repository.CreateUser(u)
		u.Id = id
		return err
	} else {
		u.Id = user.Id
		err := svc.Repository.UpdateUser(u)
		return err
	}
}

func (svc *ServiceImpl) GetUser(username string) (*User, error) {
	user, err := svc.Repository.FindUserByUsername(username)
	if err != nil {
		return &User{}, err
	}

	addresses, err := svc.Repository.GetAddressesByUserName(user.Username)
	if err != nil {
		return user, err
	}

	user.Addresses = *addresses

	return user, nil
}
