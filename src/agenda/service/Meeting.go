package service

import (
	"fmt"
	"log"
)

type Meeting struct{
	Participators []string //just record user's name
	Sponsor,Title string
	Start_time, End_time Date
}	

func delete_meeting(meeting_idx int) bool{
	meetings[meeting_idx].Participators = append(meetings[meeting_idx].Participators, meetings[meeting_idx].Sponsor)
	for _, user := range meetings[meeting_idx].Participators { //update all influenced user info
		delete_user_meeting(meetings[meeting_idx], user)
	}
	meetings = append(meetings[:meeting_idx], meetings[meeting_idx + 1 : ]...)
	defer log.Println("[Success] Delete meeting successful!")
	defer fmt.Println("[Success] Delete meeting successful!")
	return true
}

func delete_meeting_par(meeting_idx int, user_name string) bool {//delete participantor form meetings
	if !delete_user_meeting(meetings[meeting_idx], user_name) {
		return false
	}
	for index, par := range meetings[meeting_idx].Participators{ 
		if par == user_name {
			meetings[meeting_idx].Participators = append(meetings[meeting_idx].Participators[:index], meetings[meeting_idx].Participators[index + 1 :]...)
			if len(meetings[meeting_idx].Participators) == 0{ //participators == 0 then delete the whole meeting
				delete_meeting(meeting_idx)
			}
			break
		}
	}
	defer log.Println("[Success] Delete meeting's participator successful!")
	defer fmt.Println("[Success] Delete meeting's participator successful!")
	return true
}

func delete_user_meeting(meeting Meeting, user_name string) bool{//delete meeting from users
	user_idx := -1
	for index, user_tmp := range users{
		if user_tmp.User_name == user_name{
			user_idx = index
			break
		}
	}
	if user_idx == -1{
		defer log.Println("[Error] User does not partipant current meeting!")
		defer fmt.Println("[Error] User does not partipant current meeting!")
		return false
	}
	for index, meeting_tmp := range users[user_idx].User_meeting{ 
		if meeting_tmp.Title == meeting.Title {
			users[user_idx].User_meeting = append(users[user_idx].User_meeting[:index], users[user_idx].User_meeting[index + 1 :]...)
			break
		}
	}
	return true
}
func create_meeting(participators []string, sponsor, title string, start_time, end_time Date) bool{
	res := Meeting{participators, sponsor, title, start_time, end_time}
	participators = append(participators, sponsor)
	if !check_meeting(res, participators, sponsor, title, start_time, end_time){
		defer log.Println("[Error] Create meeting failed!")
		defer fmt.Println("[Error] Create meeting failed!")
		return false
	}
	meetings = append(meetings, res)
	for _, par := range participators{ //update participants meeting info
		for index, user := range users{
			if user.User_name == par {
				users[index].User_meeting = append(users[index].User_meeting, res)
				break
			}
		}
	}
	defer log.Println("[Success] Create meeting successful!")
	defer fmt.Println("[Success] Create meeting successful!")
	return true
}

func check_meeting(res Meeting, participators []string, sponsor, title string, start_time, end_time Date) bool{
	if !isValid(start_time) {
		defer log.Println("[Error] Error start time!")
		defer fmt.Println("[Error] Error start time!")
		return false
	}
	if !isValid(end_time) {
		defer log.Println("[Error] Error end time!")
		defer fmt.Println("[Error] Error end time!")
		return false
	}
	if compare_date(start_time, end_time) { // check date
		defer log.Println("[Error] start time should earily than end time!")
		defer fmt.Println("[Error] start time should earily than end time!")
		return false
	}
	for _, meeting := range meetings{ // check title
		if meeting.Title == title {
			defer log.Println("[Error] Meeting's title has exist!")
			defer fmt.Println("[Error] Meeting's title has exist!")
			return false
		}
	}
	for _, par := range participators{
		user_exist := false
		for _, user := range users{
			if user.User_name == par{ //check user existence
				user_exist = true
				for _, user_meeting := range user.User_meeting{ // check whether user meetings overlap 
					if isOverlap(res, user_meeting){
						defer log.Println("[Error] Some users'meetings may overlap!")
						defer fmt.Println("[Error] Some users'meetings may overlap!")
						return false
					}
				}
			}
		}
		if !user_exist{
			defer log.Println("[Error] Some users may not exist")
			defer fmt.Println("[Error] Some users may not exist")
			return false
		} 
	}
	for i := 0; i < len(participators); i++{ // check whether users duplicate
		for j := i + 1; j < len(participators); j++{
			if participators[i] == participators[j]{
				defer log.Println("[Error] Some users may not duplicate")
				defer fmt.Println("[Error] Some users may not duplicate")
				return false
			}
		}
	}
	return true
}

func isOverlap(meeting1, meeting2 Meeting) bool{ // false mean not overlap, input is one person's two meetings
	if compare_date(meeting1.Start_time, meeting2.End_time) || compare_date(meeting2.Start_time, meeting1.End_time){
		return false
	}else{
		return true
	}
}

