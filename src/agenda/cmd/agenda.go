/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"agenda/service"
	"github.com/spf13/cobra"
)

// agendaCmd represents the agenda command
var agendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agenda called")
		user_name, _ := cmd.Flags().GetString("user_name")
		user_password, _ := cmd.Flags().GetString("user_password")
		user_email, _ := cmd.Flags().GetString("user_email")
		user_phone, _ := cmd.Flags().GetInt("user_telephone")
		register, _ := cmd.Flags().GetBool("register")
		login, _ := cmd.Flags().GetBool("login_in")
		if(register && login){
			log.Println("[Error] Do not set -l and -r true together!")
			fmt.Println("[Error] Do not set -l and -r true together!")
			return
		}
		if(register){
			service.Init() 
			if service.Create_user(user_name, user_password, user_email, user_phone){
				service.Login_in(user_name, user_password)
			}else{
				log.Println("[Error] Some info do not fit the standard!")
				fmt.Println("[Error] Some info do not fit the standard!")
			}
		}
		if(login){
			service.Init() 
			service.Login_in(user_name, user_password)
		}
	},
}

func init() {
	rootCmd.AddCommand(agendaCmd)
	agendaCmd.Flags().StringP("user_name", "u", "", "user's name")
	agendaCmd.Flags().StringP("user_password", "p", "", "user's password")
	agendaCmd.Flags().StringP("user_email", "e", "", "user's email")
	agendaCmd.Flags().IntP("user_telephone", "t", 0, "user's telephone number")
	agendaCmd.Flags().BoolP("register", "r", false, "register a new user and login in")
	agendaCmd.Flags().BoolP("login_in", "l", false, "login in with current user")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agendaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agendaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
