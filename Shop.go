package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func LookAdsWithShop(w http.ResponseWriter, r *http.Request)  {
	var uuid = r.Header.Get("uuid")
	var item = r.Header.Get("item")
	itemId,err := strconv.Atoi(item)
	fmt.Printf("[%s]",itemId)
	if err != nil{
		fmt.Printf("[%s]",err)
	}
	rows1,err1 := db1.Query("select shop_item_1,shop_item_2,item1,item2,item3,item4,item5,diamond_count from user_info where device_id = ?",uuid)
	if err1 != nil{
		fmt.Printf("[%s]",err1)
	}
	var shop_item_1,shop_item_2,diamond_count int
	var item2,item3,item4,item5 int64
	var item1 float64
	for rows1.Next(){
		rows1.Columns()
		err := rows1.Scan(&shop_item_1,&shop_item_2,&item1,&item2,&item3,&item4,&item5,&diamond_count)
		if itemId == 1{
			if shop_item_1 > 0 {
				db1.Exec("update user_info set shop_item_1 = ?,diamond_count = ? where device_id = ?",shop_item_1 -1,diamond_count + 8,uuid)
			}
		}else if itemId == 2{
			if shop_item_2 > 0 {
				//TODO
			}
		}
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}

	returnUser(w,uuid)


	//if diamond_count >= 10{
	//	db1.Exec("update user_info  set diamond_count = ? where device_id = ?",diamond_count - 10,uuid)
	//	returnUser(w,uuid)
	//}else{
	//	returnNull(w)
	//}
}
