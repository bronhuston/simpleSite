package simpleSite

type SaveToDBService struct {
	Repository *Repository
}

type Service interface {
	save(*User) error
	getUser(string) (*User, error)
}

func (svc *SaveToDBService) save(u *User) error {
	user, _ := svc.Repository.findUserByUsername(u.Username)

	if user.Id == 0 {
		id, err := svc.Repository.createUser(u)
		u.Id = id
		return err
	} else {
		u.Id = user.Id
		err := svc.Repository.updateUser(u)
		return err
	}
}

func (svc *SaveToDBService) getUser(username string) (*User, error) {
	return svc.Repository.findUserByUsername(username)
}
