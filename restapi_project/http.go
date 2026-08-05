package main

import (
	"fmt"
	"log"
	"net/http"
)

func main1() {

	//localhost:3000
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello Server")
	})

	const serverAddress string = ":3000" //if :3000 port only, go auto knows is localhost

	fmt.Println("Server Listening on port :", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

}
