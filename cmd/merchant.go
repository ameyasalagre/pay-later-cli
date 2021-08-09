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
	"strings"

	"github.com/spf13/cobra"
)

// merchantCmd represents the merchant command
var merchantCmd = &cobra.Command{
	Use:   "merchant",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Pay-Later is a CLI app`,
	Run: func(cmd *cobra.Command, args []string) {
		merchantName := args[0]
		discountPercentage := strings.Split(args[1], "%")
		discount, error := strconv.ParseFloat(discountPercentage[0], 32)
		if error != nil {
			log.Fatal("Error While Connection", error)
		}
		components.CreateMerchant(merchantName, float32(discount))
	},
}

var updateMerchantDiscountCmd = &cobra.Command{
	Use:   "merchant",
	Short: "Change Discount",
	Long:  `Change Discount provided by merchant`,
	Run: func(cmd *cobra.Command, args []string) {
		merchantName := args[0]
		discountPercentage := strings.Split(args[1], "%")
		discount, error := strconv.ParseFloat(discountPercentage[0], 32)
		if error != nil {
			log.Fatal("Error While Connection", error)
		}
		components.UpDateMerchantDiscountByName(merchantName, float32(discount))
	},
}

var merchantReport = &cobra.Command{
	Use:   "discount",
	Short: "Discount provided by Merchant",
	Long:  `Discount provided by Merchant`,
	Run: func(cmd *cobra.Command, args []string) {
		merchantName := args[0]
		components.GetDiscountOfferedByMerchant(merchantName)
	},
}

func init() {
	newCmd.AddCommand(merchantCmd)
	updateCmd.AddCommand(updateMerchantDiscountCmd)
	reportCmd.AddCommand(merchantReport)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// merchantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// merchantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
