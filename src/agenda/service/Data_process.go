package service

import (
	"encoding/json"
	"os"
	"io"
	"bufio"
	"log"
	"fmt"
)

func read_meetings() []Meeting{
	file, err := os.Open("./service/data/Meeting.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var res []Meeting
	meetings := bufio.NewReader(file)
	for {
		meeting_json, err := meetings.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		res = append(res, meeting_decode([]byte(meeting_json)))
	}
	return res
}

func write_meetings(){
	file, err := os.OpenFile("./service/data/Meeting.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		defer log.Println("[Error] Write meeting error!")
		defer fmt.Println("[Error] Write meeting error!")
	}
	defer file.Close()
	for i := 0; i < len(meetings); i++{
		file.WriteString(string(meeting_encode(meetings[i])[:]))
		file.WriteString("\n")
	}
}

func meeting_decode(json_info []byte) Meeting{
	var res Meeting
	err := json.Unmarshal(json_info, &res)
	if err != nil {
		defer log.Println("[Error] Get meeting error!")
		defer fmt.Println("[Error] Get meeting error!")
	}
	return res
}

func meeting_encode(meeting Meeting)[]byte{
	json_info, err := json.Marshal(meeting)
	if err != nil{
		defer log.Println("[Error] Write meeting error!")
		defer fmt.Println("[Error] Write meeting error!")
	}
	return json_info
}

func read_users() []User{
	file, err := os.Open("./service/data/User.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var res []User
	users := bufio.NewReader(file)
	for {
		user_json, err := users.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		res = append(res, user_decode([]byte(user_json)))
	}
	return res
}

func write_users(){
	file, err := os.OpenFile("./service/data/User.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		defer log.Println("[Error] Write user error!")
		defer fmt.Println("[Error] Write user error!")
	}
	defer file.Close()
	for i := 0; i < len(users); i++{
		file.WriteString(string(user_encode(users[i])))
		file.WriteString("\n")
	}
}

func user_decode(json_info []byte) User{
	var res User
	err := json.Unmarshal(json_info, &res)
	if err != nil {
		defer log.Println("[Error] Get user error!")
		defer fmt.Println("[Error] Get user error!")
	}
	return res
}

func user_encode(user User)[]byte{
	json_info, err := json.Marshal(user)
	if err != nil{
		defer log.Println("[Error] Write user error!")
		defer fmt.Println("[Error] Write user error!")
	}
	return json_info
}