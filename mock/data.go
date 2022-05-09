package mock

import (
    userpb "github.com/jha929/lego-grpc/protos/user"
)

/*
user 서비스에는 총 5명의 유저가 존재하고, 
유저들이 갖고 있는 값들은 user.proto에서 정의한 값들과 같다.
*/
var Users = []*userpb.UserMessage{
	{
		UserId: "1",
		Name: "Henry",
		PhoneNumber: "01012341234",
		Age: 22,
	},
	{
		UserId: "2",
		Name: "Michael",
		PhoneNumber: "01098128734",
		Age: 55,
	},
	{
		UserId: "3",
		Name: "Jessie",
		PhoneNumber: "01056785678",
		Age: 15,
	},
	{
		UserId: "4",
		Name: "Max",
		PhoneNumber: "01099999999",
		Age: 37,
	},
	{
		UserId: "5",
		Name: "Tony",
		PhoneNumber: "01012344321",
		Age: 25,
	},
}
