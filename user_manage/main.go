package main

import (
	"log"
	"net/http"

	um "./user_manager"
)

func main() {
	http.HandleFunc("/query", um.QueryUserRouter)
	http.HandleFunc("/delete", um.DeleteUserRouter)
	http.HandleFunc("/update", um.UpdateUserRouter)
	http.HandleFunc("/insert", um.InsertUserRouter)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln("listen and server failed.")
	}
}
