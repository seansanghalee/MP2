# MP2

By Socially Distancing Group

# How to Run

### server.go

To start the server, first type on the command line,

```
go run server.go
```

The program will then prompt user to provide the port number to listen to.

```
go run server.go
Enter the port number: 
```

Type in the port number, and the server is ready to go.


### client.go

To start the client, first type on the command line,

```
go run client.go
```

The program will then prompt user to provide the IP address of the server, the port number to listen to, and the username.

```
Enter the host address:
Enter the port number:
Enter your username:
```

Provide those three fields, and the client is ready to go.

### Sending and receiving messages

To send a messasge, simply type the username to send the message to when prompted. 

```
To:
```

Then, the system will ask user to type in the message.

```
Enter your message:
```

If the client attempts to send a message to another client that does not exist, they will recieve an error message from the server.

```
//client
To: billybob
Enter your message: hello

SERVER: No Such User [billybob]
```
If the message has been sent successfully, it will appear on the receiving user's command line as,

```
JohnDoe: hello
```

### Exiting the program

A client can exit the program by typing EXIT into their terminal window. This sends a message to the server informing the server which then removes the client from the clients map. The client then automatically quits the program.
```
//client "billybob"
EXIT
//server
billybob is exiting the server...
```
The server can also exit the program by typing EXIT into their terminal window. This sends a message to all clients informing them that the server is exiting. The server and clients then automatically quit the program.
```
//server 
EXIT
//client
SERVER: Server exiting...
```


# Architecture and Design

The design includes one centralized server connecting to multiple clients using TCP.

### Message
There is a struct of type Message that stores the destination user, the originating user, and the contents of a message.  It also has a method to construct a type message struct and display a message to the terminal window in a readable format.
```
type Message struct{
    to, from, content String    
}
```
### Client
Client package includes methods that allow the client to constantly listen for incoming messages, process them, and respond accordingly.  It also contains a login method that takes user input upon initial login and sends that information to the server.

### Server
server.go includes the main function that first starts the tcp server.  It also includes two goroutines, waitForExit(): which waits for the server EXIT command, and serve(): which handles interclient communication.

# Notes
We added lots of additional features to improve user experience, such as personalized messages for each user when listing server participants!

Added mutex lock and unlock to prevent race condition in func serve() in server.go