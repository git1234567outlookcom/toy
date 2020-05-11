package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"server/dao"
	"server/model"
)

type UserService struct {
	Dao *dao.UserDao
}

func GetUserService() *UserService {
	return &UserService{Dao: dao.GetUserDao()}
}
func (s *UserService) Save(u *model.User) (string, error) {
	return s.Dao.Save(u)
}

func (s *UserService) DeleteById(u *model.User) error {
	filter := bson.M{"_id": u.ObjectId}
	return s.Dao.DeleteByFilter(filter)
}

func (s *UserService) FindList(page *model.Page) ([]*model.User, error) {
	return s.Dao.FindList(page)
}
func (s *UserService) FindPhoneOrEmail(m *model.User) (*model.User, error) {

	var filter bson.M
	if m.Phone != "" {
		filter = bson.M{"phone": m.Phone}
	}

	if m.Email != "" {
		filter = bson.M{"email": m.Email}
	}

	return s.Dao.FindOneByFilter(filter)
}

func (s *UserService) Update(u *model.User) (*model.User, error) {

	update := s.Dao.Client.FindOneAndReplace(nil, bson.M{"_id": u.ObjectId}, u)
	if update.Err() != nil {
		return nil, update.Err()
	}
	return u, nil
}
