package main

import "fmt"

func main(){

	//Anonymous Structs
	myCar := struct {
		Make string
		Model string
	} {
		Make: "tesla",
		Model: "model 3",
	}

	fmt.Printf(myCar.Model+"\n")


	//Embedded Structs
	type user struct {
		name   string
		number int
	}
	type sender struct {
		user
		rateLimit int
	}
	

	
	s := sender{
    user: user{
        name:   "John Doe",
        number: 12345,
    },
    rateLimit: 10,
}
fmt.Printf(s.name+"\n")

//Structs Methods in go 	
r := rect{
	width: 5,
	height: 10,
}

fmt.Println(r.area())
}
type rect struct {
	width int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}
//nested structs
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name=="" || mToSend.sender.number==0 || mToSend.recipient.name=="" || mToSend.recipient.number==0 {
		return false
	}
	
	return true

}