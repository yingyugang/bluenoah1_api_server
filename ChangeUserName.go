package main

import (
	"fmt"
	"net/http"
)

func ChangeUserName(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var userName = r.URL.Query()["userName"][0]
	fmt.Printf("userName [%s]",userName)
	db1.Exec("update user_info  set user_name = ? where device_id = ?",userName,uuid)
	var data= []byte(userName)
	w.Write(data)
}