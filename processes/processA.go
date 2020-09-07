package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type ProcessA struct {
	To string
	From string
	Date time.Time
	Title string
	Content string
}

//created TCP client, integrated from linode.com
func main() {

	//collect cmd line arguments
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	//connects to specified server
	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}


	for { //assign user input to variables and send to server
		reader := bufio.NewReader(os.Stdin) //read the user input

		fmt.Print(">>To ")
		inputTo, _ := reader.ReadString('\n')
		fmt.Fprintf(c, inputTo+"\n")

		fmt.Print(">>From ")
		inputFrom, _ := reader.ReadString('\n')
		fmt.Fprintf(c, inputFrom+"\n")

		time := time.Now()
		fmt.Fprintf(c, time.String()+"\n") //time pulled from package rather than user input

		fmt.Print(">>Title ")
		inputTitle, _ := reader.ReadString('\n')
		fmt.Fprintf(c, inputTitle+"\n")

		fmt.Print(">>Content ")
		inputContent, _ := reader.ReadString('\n')
		fmt.Fprintf(c, inputContent+"\n")


		message, _ := bufio.NewReader(c).ReadString('\n') //read server AWK
		fmt.Print("->: " + message)
		return
	}
}


