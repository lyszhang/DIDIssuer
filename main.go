package main

import (
	"github.com/lyszhang/DIDIssuer/log"
	"github.com/lyszhang/DIDIssuer/router"
	"fmt"
	"net/http"

)

func init() {
	err := log.LogInit()
	if err != nil {
		fmt.Println("panic: log init error")
	}
}

func main ()  {
	router := router.NewRouter()
	log.Info("server is running...")
	log.ERROR.Fatal(http.ListenAndServe(":8000", router))
}
