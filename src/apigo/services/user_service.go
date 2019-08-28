package services

import (
	"../domains"
	"../utils"
	"strconv"
	"fmt"
)

func GetUser(userId string) (*domains.User,*utils.ApiError) {
	userIdInt, err := strconv.Atoi(userId)
	//userIdInt, err := strconv.ParseInt(userId,10,64)
	if err == nil {

		fmt.Println(userIdInt)
	}
	user := domains.User{
		ID: userIdInt,
	}
	if err := user.Get(); err != nil {
		return nil,err
	}
	return &user, nil

}


