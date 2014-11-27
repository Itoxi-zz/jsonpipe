package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Request struct {
	Id     string `json:"reqId"`
	Action string `json:"action"`
	Data   Data   `json:"data"`
}

type Data struct {
	SomeData string `json:"someData"`
}

func main() {
	sendSomeData()
}

func sendSomeData() {

	conn, _ := net.Dial("tcp", ":8080")
	fmt.Println("Sending some data")

	dat := Data{SomeData: "Yuppers"}

	req := Request{
		Id:     "101",
		Action: "foo",
		Data:   dat,
	}

	bytes, err := json.Marshal(req)

	_, err = conn.Write(bytes)
	_, err = conn.Write([]byte("\n"))
	if err != nil {
		println("Write to server failed:", err.Error())
	}

	println("write to server = ", string(bytes))

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Read from server failed:", err.Error())
	}

	println("reply from server=", string(reply))
	conn.Close()
}
