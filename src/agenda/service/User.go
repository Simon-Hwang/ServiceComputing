package service

import(
	"regexp"
	"strconv"
	"log"
	"fmt"
)

type User struct{
	User_name string
	User_password string
	User_email string
	User_phone int
	User_meeting []Meeting 
}

func create_user(name, password, email string, phone int) bool{
	if !isPhone(phone) || !isEmail(email){
		return false
	}
	for _,user := range users{
		if user.User_name == name {
			defer log.Println("[Error] User's name has exist!")
			defer fmt.Println("[Error] User's name has exist!")
			return false
		}
	}
	empty := make([]Meeting, 0)
	user := User{name, password, email, phone, empty}
	users = append(users, user)
	defer log.Println("[Success] Create user successful!")
	defer fmt.Println("[Success] Create user successful!")
	return true
}

func isPhone(phone int) bool{
	res, _ := regexp.MatchString("^1[0-9]{10}$", strconv.Itoa(phone))
	if !res{
		defer log.Println("[Error] Phone's format is not correct!")
		defer fmt.Println("[Error] Phone's format is not correct!")
	}
	return res
}

func isEmail(email string) bool{
	res, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	if !res{
		defer log.Println("[Error] Email's format is not correct!")
		defer fmt.Println("[Error] Email's format is not correct!")
	}
	return res
}