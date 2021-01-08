package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Purchase(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var diamondStr =  r.Header.Get("diamond")
	diamondint,err := strconv.Atoi(diamondStr)
	rows,err := db1.Query("select diamond_count  from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	var diamond_count int
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&diamond_count)
		if err != nil{
			fmt.Printf("returnUser:get user info error [%s]",err)
		}
		break
	}
	rows.Close()
	db1.Exec("update user_info  set diamond_count = ? where device_id = ?",diamond_count + diamondint ,uuid)
	returnUser(w,uuid)
}