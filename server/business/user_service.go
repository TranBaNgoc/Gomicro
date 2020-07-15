package business

import (
	"context"
	"gomicro/server/model"
	proto "gomicro/server/protobuf"
)

type UserService struct{}

func (u UserService) GetListUsers(c context.Context, request *proto.EmptyRequest, response *proto.GetListUserResponse) error {
	//users, err := model.GetUsers()
	//if err != nil {
	//	return err
	//}
	//
	//var list []*proto.User
	//for _, u := range users {
	//	list = append(list, &proto.User{
	//		Name: u.Name,
	//		Age:  u.Age,
	//	})
	//}
	//response.User = list
	response.User = []*proto.User{}
	return nil
}

func (u UserService) AddUser(context context.Context, request *proto.User, response *proto.AddUserResponse) error {
	newUser := &model.User{
		Name: request.Name,
		Age:  request.Age,
	}
	if err := model.AddUser(newUser); err != nil {
		return err
	}
	response.Message = "Them thanh cong"
	return nil
}
