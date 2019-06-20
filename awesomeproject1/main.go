package  main

import (
	"../dataBasePackage"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func server() {
	dbColunms := [7]string{"ID","DATE","CATEGORY","PRODUCT","VALUE","Expenses_talbe", ""}
	var db = dataBasePackage.NewDB(dbColunms)
	db.CreatDB()
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, db)
	}
}

func handleRequest(conn net.Conn, db *dataBasePackage.DataBaseHandler) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	var msg string
	msg = string(buf)
	fmt.Println("Messege is : " + msg)
	for string(buf) != "Bye" {
		strArray := strings.Split(string(buf), ",")
		fmt.Println(strArray[0])
		switch strArray[0] {
		case "GET":

			break
		case "DELETE":
			fmt.Println("Deleting ...." + strArray[1])
			id := strings.Split(string(strArray[1]), "\n")
			db.Delete(id[0])
			fmt.Println("Finish Deleting !")
			break
		case "INSERT":
			fmt.Println("Inserting ....")
			var id= db.Insert(strArray[1], strArray[2], strArray[3],strings.Split(string(strArray[4]), "\n")[0])
			fmt.Println(id)
			_, err = conn.Write([]byte(id+"\n"))
			fmt.Println("Sending respons ....")
			if err != nil {
				panic(err)
			}
			break
		default:
			break;
		}
		var i int
		for i = 0; i < 1024; i++{
			buf[i] = 0
		}

		// Read the incoming connection into the buffer.
		_, err := conn.Read(buf)
		msg = string(buf)
		fmt.Println("Messege is : " + msg)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
	}
	// Close the connection when you're done with it.
	conn.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error reading....")
		panic(err)
	}
}

func main() {
	server()
}
