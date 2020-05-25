package business

import (
	"context"
	"fmt"
	"gomicro/server/model"
	proto "gomicro/server/protobuf"
	"time"
)

type UserService struct{}

func (u UserService) GetListUsers(c context.Context, request *proto.EmptyRequest, response *proto.GetListUserResponse) error {
	fmt.Print(-time.Now().Nanosecond())
	fmt.Print("+")
	users, err := model.GetUsers()
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	var list []*proto.User
	for _, u := range users {
		list = append(list, &proto.User{
			Name: u.Name,
			Age:  u.Age,
		})
	}
	response.User = list
	fmt.Print(time.Now().Nanosecond())
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
