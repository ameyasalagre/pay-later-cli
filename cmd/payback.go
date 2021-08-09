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

// paybackCmd represents the payback command
var paybackCmd = &cobra.Command{
	Use:   "payback",
	Short: "PayBack and restore Credit Limit",
	Long: `PayBack and restore Credit Limit`,
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		credit, error := strconv.ParseFloat(args[1], 32)
		if error != nil {
			log.Fatal("UnknownError Occured", error)
		}
		components.CreatePayBack(userName,float32(credit))
	},
}

func init() {
	rootCmd.AddCommand(paybackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// paybackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// paybackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
