package service

import (
	"fmt"
	"log"
	"strconv"
)

var(
	users []User
	meetings []Meeting
	user_index int
	lf LogFile
)

func Init(){
	users = read_users()
	meetings = read_meetings()
	user_index = -1
	lf = LogFile{file: "agenda.log"}
	log.SetOutput(&lf)
}

func GetUsers() []User{
	return users
}

func GetMeetings() []Meeting{
	return meetings
}

func Create_user(name, password, email string, phone int) bool {
	return create_user(name, password, email, phone)
}

func Login_in(name, password string){
	for index, user := range users{
		if user.User_name == name && user.User_password == password{
			user_index = index
			break 
		}
	}
	if user_index == -1{
		log.Println("[Error] User not exist or input error info!")
		fmt.Println("[Error] User not exist or input error info!")
		return 
	}
	process()
}

func process(){
	var choice string
	for {
		fmt.Println("--q 	input q to quit agenda")
		fmt.Println("--cm 	create a new meeting")
		fmt.Println("--mr	remove a meeting")
		fmt.Println("--pr	remove a particopator from a meeting")
		fmt.Println("--qm	query current user's meeting")
		fmt.Scanln(&choice)
		if(choice == "q"){
			break
		}else if(choice == "cm"){
			cm_cmd()
		}else if(choice == "mr"){
			mr_cmd()
		}else if(choice == "pr"){
			pr_cmd()
		}else if(choice == "qm"){
			qm_cmd()
		}
	}
	write_users()
	write_meetings()
}

func qm_cmd(){
	for _, meeting := range users[user_index].User_meeting{
		fmt.Println(string(meeting_encode(meeting)))
	}
}
func pr_cmd(){
	var title, user_name string
	fmt.Println("Input the meeting title")
	fmt.Scanln(&title)
	fmt.Println("Input the participator's name")
	fmt.Scanln(&user_name)
	meeting_idx := -1
	for index, meeting_tmp := range meetings{
		if meeting_tmp.Title == title{
			meeting_idx = index
			break
		}
	}
	if meeting_idx == -1 {
		defer log.Println("[Error] Meeting not exist!")
		defer fmt.Println("[Error] Meeting not exist!")
		return
	}
	delete_meeting_par(meeting_idx, user_name)
}

func mr_cmd(){
	var title string
	fmt.Println("Input the meeting title you want to remove")
	fmt.Scanln(&title)
	meeting_idx := -1
	for index, meeting_tmp := range meetings{
		if meeting_tmp.Title == title{
			meeting_idx = index
			break
		}
	}
	if meeting_idx == -1 {
		defer log.Println("[Error] Meeting not exist!")
		defer fmt.Println("[Error] Meeting not exist!")
		return
	}
	delete_meeting(meeting_idx)
}

func cm_cmd(){
	var participators []string
	//sponsor is default to be current user
	var par, par_number, title, start_time, end_time string
	fmt.Println("Input your meeting's participators number: ")
	fmt.Scanln(&par_number)
	par_num, _ := strconv.Atoi(par_number)
	for i := 0; i < par_num ; i++{
		fmt.Println("Input number ", i, "participator's name")
		fmt.Scanln(&par)
		participators = append(participators, par)
	}
	fmt.Println("Input your meeting's title: ")
	fmt.Scanln(&title)
	fmt.Println("Input your meeting's start time: ")
	fmt.Scanln(&start_time)
	fmt.Println("Input your meeting's end time: ")
	fmt.Scanln(&end_time)
	create_meeting(participators, users[user_index].User_name, title, StringToDate(start_time), StringToDate(end_time))
}