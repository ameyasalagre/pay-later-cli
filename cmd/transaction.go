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
	"log"
	"pay-later/components"
	"strconv"

	"github.com/spf13/cobra"
)

// merchantCmd represents the merchant command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Pay-Later is a CLI app`,
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		merchant := args[1]
		transactionAmt,error :=  strconv.ParseFloat(args[2], 32)
		if error != nil {
			log.Fatal("Error While Connection", error)
		}
		components.InitiatePayout(userName,merchant,float32(transactionAmt))
	
	},
}

func init() {
	newCmd.AddCommand(transactionCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// merchantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// merchantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
