package main

import "fmt"

func main() {
	 username  := "hrushi"
	 password  := "2020BCS022"
	fmt.Println("Authorization: Basic", username+":"+password)

	smsSendingLimit:=0
	costPerSMS:=0.0
	hasPermission:= false
	name:=""

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, name)

	const new_name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", new_name, openRate)

	fmt.Print(msg)
}
