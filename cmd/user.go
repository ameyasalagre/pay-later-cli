/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"pay-later/components"
	"strconv"

	// "strconv"
	// "strings"
	// "text/template/parse"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long:  `pay-later create user bob`,
	Run: func(cmd *cobra.Command, args []string) {

		userName := args[0]
		emailId := args[1]
		credit, error := strconv.ParseFloat(args[2], 32)
		if error != nil {
			log.Fatal("Error While Connection", error)
		}
		fmt.Println("UserName : " + userName)
		fmt.Println("emailId : " + emailId)
		components.CreateUser(userName, emailId, float32(credit), float32(credit))
	},
}

var updateUser = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long:  `pay-later create user bob`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var getUserDueCmd = &cobra.Command{
	Use:   "dues",
	Short: "A brief description of your command",
	Long:  `pay-later create user bob`,
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		components.GetDueOfUser(userName)
	},
}

var getAllUsersDueCmd = &cobra.Command{
	Use:   "total-dues",
	Short: "A brief description of your command",
	Long:  `pay-later create user bob`,
	Run: func(cmd *cobra.Command, args []string) {
		components.GetDuesOfAllUsers()
	},
}

func init() {
	newCmd.AddCommand(userCmd)
	updateCmd.AddCommand(updateUser)
	reportCmd.AddCommand(getUserDueCmd)
	reportCmd.AddCommand(getAllUsersDueCmd)

}
