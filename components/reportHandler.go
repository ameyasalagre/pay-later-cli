package components

import (
	"fmt"
	"log"
)

func GetDiscountOfferedByMerchant(merchantName string) string{
	result,error:=getMerchantByName(merchantName)
	if(error!=""){
		return error
	}
	log.Println("Merchant : "+ result.MerchantName) 
	log.Println("Discount : " , float64(result.Discount))
	return fmt.Sprintln("Merchant :"+ result.MerchantName +  "is providing ") +
	fmt.Sprintln(float64(result.Discount))
}

func GetDueOfUser(userName string) string{
	result,error:=	getUserByUserName(userName)
	if(error!=""){
		return error
	}
	due := result.Limit - result.Credit
	log.Println("User : "+ result.UserName) 
	log.Println("Due : " , float64(due))
	return fmt.Sprintln("User :"+ result.UserName ) +
	fmt.Sprintln("Due Amount is",float64(due))
}

func GetDuesOfAllUsers() string {
	_,err:=getAllUsers()
	if(err!=""){
		return err
	}

	return "All Clients Fetched Successfully"
}