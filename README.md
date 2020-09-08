# HW0


# To Run: 
1. Open the directory /processes within terminal in 2 terminal instances
2. Run processB with the following command: go run processB.go #### (input a port in place of the pound signs)
3. Run processA with the following command in the other terminal instance: go run processA.go 127.0.0.1:#### (enter the same port as entered above)
4. Input your desired values into the prompt, pressing enter each time
5. Confirm that the input fields printed in the processB instances, and that processA printed the AWK "Email Recieved"

# Citations 
PKG knowledge and general go knowledge from https://golang.org/

Project design heavily influenced from Darius Russell Kish: https://github.com/Dariusrussellkish/Example-MP-Solution 

TCP Server and Client code integrated from https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

# Documentation
The layout is copied over from Darius', with a file for the email struct and two files for the processes. 

ProcessA contains the TCP client code from linode.com. It starts of by connecting to the user provided ip and port. Then, it has a loop that is constantly reading user inputs and sending them to the server. The code that sends the user input is from linode.com too, but I changed the prompts and intitialized variables.

ProcessB features the TCP server code from linode.com. It creates a server from the user inputed port. Then, I created an instances of the email struct, so that processB could reconstruct the message to print. In a for loop, it reads the messages sent from the client. A for loop is useful here becasue the user inputs will always be in the order: To, From, Title, Content. Therefore, I can assign the struct parameter values based on the loop interation. Finally, on the last iteration, the process prints the entire struct, sends ACK to processA, and closes. When processA recieves the ACK, it too close.
