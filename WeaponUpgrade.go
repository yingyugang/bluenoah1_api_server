package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)
var grade1 int = 20
var grade2 int = 80
var grade3 int = 120
var grade4 int = 220
var grade5 int = 400
var grade6 int = 999
func WeaponUpgrade(w http.ResponseWriter, r *http.Request)  {
	var uuid = r.Header.Get("uuid")
	var weapon = r.Header.Get("weapon")
	weaponId,err :=  strconv.Atoi(weapon)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var column string
	switch weaponId {
	case 0:
		column = "ak47_lvl"
		break
	case 1:
		column = "m16_lvl"
		break
	case 2:
		column = "scatter_lvl"
		break
	case 3:
		column = "laser_lvl"
		break
	case 4:
		column = "firegun_lvl"
		break
	case 5:
		column = "rpg_lvl"
		break
	case 6:
		column = "laserx_lvl"
		break
	case 7:
		column = "awp_lvl"
		break
	}

	var lvl int
	var item2,item3,item4,item5,item6 int64
	var item1 float64
	rows,err := db1.Query("select " + column + ",item1,item2,item3,item4,item5,item6 from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&lvl,&item1,&item2,&item3,&item4,&item5,&item6)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	var upgradeData = getUpgradeData(lvl)
	if upgradeData.Coin <= item1 && upgradeData.GreenGear <= item2 && upgradeData.BlueGear <= item3 && upgradeData.PurpleGear <= item4 && upgradeData.OrangeGear <= item5 && upgradeData.SupperGear <=item6{
		item1 -= upgradeData.Coin
		item2 -= upgradeData.GreenGear
		item3 -= upgradeData.BlueGear
		item4 -= upgradeData.PurpleGear
		item5 -= upgradeData.OrangeGear
		item6 -= upgradeData.SupperGear
		db1.Exec("update user_info  set " + column + " = ?,item1 = ?,item2 = ?,item3 = ?,item4 = ?,item5 = ?,item6 = ? where device_id = ?",lvl + 1,item1,item2,item3,item4,item5,item6,uuid)
	}
	returnUser(w,uuid)
}



func WeaponUpgradeBulk(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var weapon = r.Header.Get("weapon")
	weaponId,err :=  strconv.Atoi(weapon)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var column string
	switch weaponId {
	case 0:
		column = "ak47_lvl"
		break
	case 1:
		column = "m16_lvl"
		break
	case 2:
		column = "scatter_lvl"
		break
	case 3:
		column = "laser_lvl"
		break
	case 4:
		column = "firegun_lvl"
		break
	case 5:
		column = "rpg_lvl"
		break
	case 6:
		column = "laserx_lvl"
		break
	case 7:
		column = "awp_lvl"
		break
	}
	var lvl int
	var item2,item3,item4,item5,item6 int64
	var item1 float64
	rows,err := db1.Query("select " + column + ",item1,item2,item3,item4,item5,item6 from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&lvl,&item1,&item2,&item3,&item4,&item5,&item6)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	if lvl == grade1 || lvl == grade2 || lvl == grade3 || lvl == grade4 || lvl == grade5 {
		returnUser(w,uuid)
		return
	}
	var maxLvl = getMaxLevel(lvl)
	var totalCost float64 = 0
	var nextLvl = lvl
	for i := lvl + 1;i <= maxLvl;i++  {
		var upgradeData = getUpgradeData(i)
		if upgradeData.Coin + totalCost <= item1{
			totalCost += upgradeData.Coin
			nextLvl = i
		}else{
			break
		}
	}

	if nextLvl > lvl {
		db1.Exec("update user_info  set " + column + " = ?,item1 = ? where device_id = ?",nextLvl,item1 - totalCost,uuid)
	}
	returnUser(w,uuid)
}
func getMaxLevel(currentLvl int)(maxLevl int){
	if currentLvl < grade1{
		return grade1
	}
	if currentLvl < grade2{
		return grade2
	}
	if currentLvl < grade3{
	return grade3
	}
	if currentLvl < grade4{
	return grade4
	}
	if currentLvl < grade5{
	return grade5
	}
	if currentLvl > grade5{
		return grade6
	}
	return grade6
}
var LvlUpGold float64 = 70
var CommonMulti float64 = 1.14
func getUpgradeData(level int)(upgradeData UpgradeData) {
	var coin float64
	coin = LvlUpGold * math.Pow(CommonMulti,float64(level - 1))
	upgradeData = UpgradeData{Coin:coin }
	if level == grade1 {
		upgradeData.GreenGear = 10
	}else if level == grade2 {
		upgradeData.GreenGear = 25
		upgradeData.BlueGear = 10
	}else if level == grade3 {
		upgradeData.GreenGear = 95
		upgradeData.BlueGear = 70
		upgradeData.PurpleGear = 25
	} else if level == grade4 {
		upgradeData.BlueGear = 135
		upgradeData.PurpleGear = 80
		upgradeData.OrangeGear = 30
	}else if level == grade5 {
		upgradeData.PurpleGear = 250
		upgradeData.OrangeGear = 160
		upgradeData.SupperGear = 15
	}
	return upgradeData
}



