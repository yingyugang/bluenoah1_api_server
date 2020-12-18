package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func inherenceUpgrade(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var inherence =  r.Header.Get("inherence")
	var diamondStr =  r.Header.Get("diamond")
	var coinStr = r.Header.Get("coin")
	coinfloat,err := strconv.ParseFloat(coinStr,64)
	diamondint,err := strconv.Atoi(diamondStr)
	fmt.Printf("inherence: [%s]",inherence)
	fmt.Printf("diamondStr: [%s]",diamondStr)

	var column string
	var max int = 35
	//var plus int
	switch inherence {
	case "0":
		column = "atk_up"
		//max = 10000 //100倍
		//plus = 15 //15%
		break
	case "1":
		column = "atk_boss_up"
		//max = 10000 //100倍
		//plus = 15 //15%
		break
	case "2":
		column = "speed_up"
		//max = 10000 //100倍
		//plus = 15 //15%
		break
	case "3":
		column = "buff_up"
		//max = 1000 //1倍
		//plus = 10 //1%
		break
	case "4":
		column = "dodge_up"
		//max = 1000
		//plus = 10 //1%
		break
	case "5":
		column = "atk_speed_up"
		//max = 1000 //1倍
		//plus = 10 //1%
		break
	case "6":
		column = "coin_up"
		//max = 10000
		//plus = 15//15%
		break
	case "7":
		column = "critical_up"
		//max = 5
		//plus = 1
		break
	case "8":
		column = "hp_up"
		//max = 1000
		//plus = 10//1%
		break
	}

	rows,err := db1.Query("select " + column +",item1,diamond_count from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var inherenceV int
	var item1 float64
	var diamond int
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&inherenceV,&item1,&diamond)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	if inherenceV < max && item1 >= coinfloat && diamond >= diamondint {
		inherenceV += 1
	}
	db1.Exec("update user_info  set " + column + " = ?,item1 = ?,diamond_count = ? where device_id = ?",inherenceV,item1-coinfloat,diamond-diamondint,uuid)
	returnUser(w,uuid)
}
