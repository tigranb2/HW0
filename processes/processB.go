package main

import (
	"../email"
	"bufio"
	"fmt"
	"net"
	"os"
)

//creates tcp client, integrated from linode.com
func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	//open server on port
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close() //ends process

	//server can interact with clients only if call is successful
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	var recievedEmail email.Email //struct which contains email fields provided in assignment description

	//loop, runs through inputs and assigns them to correct struct fields
	//for works well here because user inputs are always entered in order
	for i := 0; i < 5; i++{
		netData, err := bufio.NewReader(c).ReadString('\n') //read client input
		if err != nil {
			fmt.Println(err)
			return
		}

		switch i { //depending on iteration, set struct fields to client input
		case 0: recievedEmail.To = string(netData)
		case 1: recievedEmail.From = string(netData)
		case 2: recievedEmail.Date = string(netData)
		case 3: recievedEmail.Title = string(netData)
		case 4: recievedEmail.Content = string(netData)
		}

		if i == 4 { //don't know how to have switch case with multiple arguments
			fmt.Printf("%+v\n", recievedEmail) //%+v adds field names, integrated from golang.org
			c.Write([]byte("Email Recieved")) //AWK
			return
		}
	}
}