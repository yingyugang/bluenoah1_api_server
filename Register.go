package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	var uuid = checkSignin(r)
	returnUser(w,uuid)
}


func checkSignin(r *http.Request)(uuidResult string)  {
	var uuid = r.Header.Get("uuid")
	var ios = r.Header.Get("iosUser")
	var hasIos = r.Header.Get("hasIosUser")
	var android = r.Header.Get("androidUser")
	var hasAndroid = r.Header.Get("hasAndroidUser")
	var user = ""
	if len(r.URL.Query()["userName"]) > 0{
		user = r.URL.Query()["userName"][0]
	}
	if len(user) == 0{
		user = "New User"
	}
	hasIosInt,err := strconv.Atoi(hasIos)
	hasAndroidInt,err := strconv.Atoi(hasAndroid)
	var mapUser map[string]int
	mapUser = make(map[string]int)
	if ios == "1000"{
		hasIosInt = 0
	}
	if android == "1000"{
		hasAndroidInt = 0
	}

	if len(uuid)==0{
		//TODO to verify the user id.
		if hasIosInt==1{
			rows,err := db1.Query("select id,user_name,has_ios_account,ios_account,device_id from user_info where ios_account = ?",ios)
			if err != nil{
				fmt.Printf("select fail [%s]",err)
			}
			var id,has_ios_account int
			var username,ios_account,device_id string
			for rows.Next(){
				rows.Columns()
				err := rows.Scan(&id,&username,&has_ios_account,&ios_account,&device_id)
				if err != nil{
					fmt.Printf("get user info error [%s]",err)
				}
				mapUser[username] = id
				uuid = device_id
				break
			}
			rows.Close()
		}else if hasAndroidInt==1 {
			rows,err := db1.Query("select id,user_name,has_android_account,android_account,device_id from user_info where android_account = ?",android)
			if err != nil{
				fmt.Printf("select fail [%s]",err)
			}
			var id,has_android_account int
			var username,android_account,device_id string
			for rows.Next(){
				rows.Columns()
				err := rows.Scan(&id,&username,&has_android_account,&android_account,&device_id)
				if err != nil{
					fmt.Printf("get user info error [%s]",err)
				}
				mapUser[username] = id
				uuid = device_id
				break
			}
			rows.Close()
		}
	}else{
		rows,err := db1.Query("select id,user_name from user_info where device_id = ?",uuid)
		if err != nil{
			fmt.Printf("select fail [%s]",err)
		}
		var id int
		var username string
		for rows.Next(){
			rows.Columns()
			err := rows.Scan(&id,&username)
			if err != nil{
				fmt.Printf("get user info error [%s]",err)
			}
			mapUser[username] = id
			break
		}
		rows.Close()
	}

	if len(mapUser) == 0 {
		var newuuid = createUUID()
		r1, err1 := db1.Exec("insert into user_info (user_name,device_id,has_ios_account,ios_account,has_android_account,android_account) values (?,?,?,?,?,?)",user,newuuid,hasIosInt,ios,hasAndroidInt,android)
		id, err1 := r1.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err1)
		}
		r2, err1 := db1.Exec("insert into hero_info (hero_name,user_id) values (?,?)","New hero",id)
		r2.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err1)
			return
		}
		uuid = newuuid
	}

	rows1,err1 := db1.Query("select lastday,loginday,bonus from user_info where device_id = ?",uuid)
	if err1 != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var lastday,loginday,bonus int
	for rows1.Next(){
		rows1.Columns()
		err := rows1.Scan(&lastday,&loginday,&bonus)
		var day = time.Now().YearDay()
		if lastday != day{
			db1.Exec("update user_info  set lastday = ?,loginday = ?,shop_item_1 = 5,shop_item_2 = 5 ,bonus = 0 where device_id = ?",day,loginday + 1,uuid)
		}
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	rows1.Close()
	return uuid
}