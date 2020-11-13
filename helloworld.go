package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}
//注意，首字母一定要大写
type User struct {
	Uuid string
	HeroId int
	Atk int
	Def int
	MaxHp int
	Hp int
	MaxSp int
	Sp int
	Level int
	Stage int
	Item1 int64
	Item2 int64
	Item3 int64
	Item4 int64
	Item5 int64
	Critical int
	Current_w int
	Ak47_lvl int
	M16_lvl int
	Scatter_lvl int
	Firegun_lvl int
	Rpg_lvl int
	Laserx_lvl int
	Awp_lvl int
	Atk_up int
	Atk_speed_up int
	Critical_up int
	Speed_up int
	Atk_boss_up int
	Hp_up int
	Diamond_count int
	Dodge_up int
	Lastday int
	Loginday int
	Bonus int
}

type BattleResult struct {
	Stage int
	Item1 int64
	Item2 int64
	Item3 int64
	Item4 int64
	Item5 int64
	Clear int
}

var db1 *sql.DB

func createUUID() (uid string) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	return u.String()
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):]
	//p, _ := loadPage(title)
	//db,err := sql.Open("mysql","root:Yyg810412@tcp(127.0.0.1:3306)/BlueNoah?charset=utf8");
	//if err != nil{
	//	fmt.Printf("connect mysql fail ! [%s]",err)
	//}else{
	//	fmt.Println("connect to mysql success")
	//}
	rows,err := db1.Query("select id,user_name from user_info");
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}

	var mapUser map[string]int
	mapUser = make(map[string]int)

	for rows.Next(){
		var id int
		var username string
		rows.Columns()
		err := rows.Scan(&id,&username)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		mapUser[username] = id
	}

	for k,v := range mapUser{
		fmt.Println(k,v);
	}
	r1, err1 := db1.Exec("insert into user_info (user_name,device_id) values (?,?)","aaa","bbb")
	id, err1 := r1.LastInsertId()
	if err1 != nil {
		fmt.Println("exec failed, ", err1)
		return
	}
	fmt.Println("insert succ:", id)
}

func loginViewHandler(w http.ResponseWriter, r *http.Request) {
	var uuid = checkSignin(r)
	returnUser(w,uuid)
}

func returnUser(w http.ResponseWriter,uuid string){
	var id,stage,heroId,atk,def,maxHp,hp,maxSp,sp,level,critical,current_w,ak47_lvl,m16_lvl,scatter_lvl,firegun_lvl,rpg_lvl,laserx_lvl,awp_lvl int
	var atk_up,atk_speed_up,critical_up,speed_up,atk_boss_up,hp_up,diamond_count,dodge_up int
	var lastday,loginday,bonus int

	var item1,item2,item3,item4,item5 int64
	rows,err := db1.Query("select id,stage,item1,item2,item3,item4,item5,current_w,ak47_lvl,m16_lvl,scatter_lvl,firegun_lvl,rpg_lvl,laserx_lvl,awp_lvl,atk_up,atk_speed_up,critical_up,speed_up,atk_boss_up,hp_up,diamond_count,dodge_up,lastday,loginday,bonus  from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&id,&stage,&item1,&item2,&item3,&item4,&item5,&current_w,&ak47_lvl,&m16_lvl,&scatter_lvl,&firegun_lvl,&rpg_lvl,&laserx_lvl,&awp_lvl,&atk_up,&atk_speed_up,&critical_up,&speed_up,&atk_boss_up,&hp_up,&diamond_count,&dodge_up,&lastday,&loginday,&bonus)
		if err != nil{
			fmt.Printf("returnUser:get user info error [%s]",err)
		}
		break
	}
	rows1, err := db1.Query("select id,atk,def,maxhp,hp,maxsp,sp,level,critical from hero_info where user_id = ?",id)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows1.Next(){
		rows1.Columns()
		err := rows1.Scan(&heroId,&atk,&def,&maxHp,&hp,&maxSp,&sp,&level,&critical)
		if err != nil{
			fmt.Printf("returnUser:get hero info error [%s]",err)
		}
		break
	}
	user := User{Uuid:uuid,HeroId:heroId,Atk:atk,Def:def,MaxHp:maxHp,Hp:hp,MaxSp:maxSp,Sp:sp,Level:level,Stage:stage,Item1:item1,Item2:item2,Item3:item3,Item4:item4,Item5:item5,Critical: critical,Current_w:current_w,Ak47_lvl:ak47_lvl,M16_lvl:m16_lvl,Scatter_lvl:scatter_lvl,Firegun_lvl:firegun_lvl,Rpg_lvl:rpg_lvl,Laserx_lvl:laserx_lvl,Awp_lvl:awp_lvl,Atk_up:atk_up,Atk_speed_up:atk_speed_up,Critical_up:critical_up,Speed_up:speed_up,Atk_boss_up:atk_boss_up,Hp_up:hp_up,Diamond_count:diamond_count,Dodge_up:dodge_up,Lastday:lastday,Loginday:loginday,Bonus:bonus  }
	result,err := json.Marshal(user)
	fmt.Printf(string(result) )
	w.Write(result)
}

func checkSignin(r *http.Request)(uuidResult string)  {
	var uuid = r.Header.Get("uuid")
	rows,err := db1.Query("select id,user_name from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var mapUser map[string]int
	mapUser = make(map[string]int)

	for rows.Next(){
		var id int
		var username string
		rows.Columns()
		err := rows.Scan(&id,&username)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		mapUser[username] = id
		break
	}
	if len(mapUser) == 0 {
		var newuuid = createUUID()
		r1, err1 := db1.Exec("insert into user_info (user_name,device_id) values (?,?)","New user",newuuid)
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
			db1.Exec("update user_info  set lastday = ?,loginday = ? ,bonus = 0 where device_id = ?",day,loginday + 1,uuid)
		}
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	return uuid
}

func weaponUpgrade(w http.ResponseWriter, r *http.Request)  {
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
		column = "firegun_lvl"
		break
	case 4:
		column = "rpg_lvl"
		break
	case 5:
		column = "laserx_lvl"
		break
	case 6:
		column = "awp_lvl"
		break
	}

	var lvl int
	var item1 int64
	rows,err := db1.Query("select " + column + ",item1 from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&lvl,&item1)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	var coin = int64(math.Floor (70 * math.Pow(1.14,	float64(lvl-1))))
	fmt.Printf("coin:[%s]\n",math.Pow(1.14,	float64(lvl + 1)))
	fmt.Printf("select [%s]\n",lvl)
	if item1 >= coin {
		item1-=coin
		db1.Exec("update user_info  set " + column + " = ?,item1 = ? where device_id = ?",lvl + 1,item1,uuid)
		returnUser(w,uuid)
	}
}

func inherenceUpgrade(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var inherence =  r.Header.Get("inherence")
	var column string
	var max int
	var plus int
	switch inherence {
	case "0":
		column = "atk_up"
		max = 10000 //100倍
		plus = 15 //15%
		break
	case "1":
		column = "atk_speed_up"
		max = 1000 //1倍
		plus = 10 //1%
		break
	case "2":
		column = "critical_up"
		max = 1000
		plus = 10 //1%
		break
	case "3":
		column = "speed_up"
		max = 1000 //1倍
		plus = 10 //1%
		break
	case "4":
		column = "atk_boss_up"
		max = 10000
		plus = 15//15%
		break
	case "5":
		column = "hp_up"
		max = 5
		plus = 1
		break
	case "6":
		column = "dodge_up"
		max = 1000
		plus = 10//1%
		break
	}

	rows,err := db1.Query("select " + column +" from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var inherenceV int
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&inherenceV)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	if(inherenceV < max){
		inherenceV += plus
		if(inherenceV > max){
			inherenceV = max
		}
	}
	db1.Exec("update user_info  set " + column + " = ? where device_id = ?",inherenceV,uuid)
	returnUser(w,uuid)
}

func returnNull(w http.ResponseWriter){
	fmt.Fprintf(w,"")
}

func stageClear(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	//TODO
	var result = r.Header.Get("result")
	fmt.Printf(result)
	var battleResult BattleResult
	json.Unmarshal([]byte(result),&battleResult)

	rows,err := db1.Query("select stage,item1,item2,item3,item4,item5 from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var stage int
	var item1,item2,item3,item4,item5 int64
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&stage,&item1,&item2,&item3,&item4,&item5)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	if battleResult.Stage == stage && battleResult.Clear == 1{
		db1.Exec("update user_info  set stage = ?,item1 = ?,item2 = ?,item3 = ?,item4 = ?,item5 = ? where device_id = ?",stage + 1,battleResult.Item1 + item1,battleResult.Item2 + item2,battleResult.Item3 + item3,battleResult.Item4 + item4,battleResult.Item5 + item5,uuid)
	}else{
		db1.Exec("update user_info  set item1 = ?,item2 = ?,item3 = ?,item4 = ?,item5 = ? where device_id = ?",battleResult.Item1 + item1,battleResult.Item2 + item2,battleResult.Item3 + item3,battleResult.Item4 + item4,battleResult.Item5 + item5,uuid)
	}
	returnUser(w,uuid)
}

func setCurrentWeapon(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var weapon = r.Header.Get("weapon")
	fmt.Printf("select fail [%s]",weapon)
	weaponId,err :=  strconv.Atoi(weapon)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	db1.Exec("update user_info  set current_w = ? where device_id = ?",weaponId,uuid)
	returnUser(w,uuid)
}

func revive(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var diamond_count int
	rows,err := db1.Query("select diamond_count from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&diamond_count)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	if diamond_count >= 10{
		db1.Exec("update user_info  set diamond_count = ? where device_id = ?",diamond_count - 10,uuid)
		returnUser(w,uuid)
	}else{
		returnNull(w);
	}
}

func loginBonusObtain(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var ads = r.Header.Get("ads")
	var lastday,loginday,bonus int
	var item1,item2,item3,item4,item5,diamond_count int64
	rows,err := db1.Query("select item1,item2,item3,item4,item5,diamond_count,lastday,loginday,bonus  from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&item1,&item2,&item3,&item4,&item5,&diamond_count,&lastday,&loginday,&bonus)
		if err != nil{
			fmt.Printf("returnUser:get user info error [%s]",err)
		}
		break
	}
	if bonus == 0{
		var day = (loginday -1) % 7
		var week = (loginday -1) / 7
		var multi int64 = 1
		if ads == "1"{
			multi = 2
		}
		switch day {
		case 0:
			var coin = int64(math.Floor (1000000 * math.Pow(1.14,float64(week))))
			db1.Exec("update user_info set item1 = ?,bonus = ? where device_id = ?",item1 + coin * multi,1,uuid)
			fmt.Printf(uuid)
			break
		case 1:
			var diamond int64 = 11
			db1.Exec("update user_info  set diamond_count = ?,bonus = 1 where device_id = ?",diamond_count + diamond * multi,uuid)
			break
		case 2:
			var coin = int64(math.Floor (70 * math.Pow(1.14,	float64(week)))) * 2
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * multi,uuid)
			break
		case 3:
			var coin = int64(math.Floor (70 * math.Pow(1.14,	float64(week)))) * 3
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * multi,uuid)
			break
		case 4:
			var diamond int64 = 20
			db1.Exec("update user_info  set diamond_count = ?,bonus = 1 where device_id = ?",diamond_count + diamond * multi,uuid)
			break
		case 5:
			var coin = int64(math.Floor (70 * math.Pow(1.14,	float64(week))))* 5
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * multi,uuid)
			break
		case 6:
			var diamond int64 = 30
			db1.Exec("update user_info  set diamond_count = ?,bonus = 1 where device_id = ?",diamond_count + diamond * multi,uuid)
			break
		}
		returnUser(w,uuid)
	}else{
		returnNull(w)
	}
}

func main() {
	db,err := sql.Open("mysql","root:810412@tcp(35.187.200.112:3306)/BlueNoah?charset=utf8")
	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}
	db1 = db
	http.HandleFunc("/login", loginViewHandler)
	http.HandleFunc("/stage_clear",stageClear)
	http.HandleFunc("/weapon_upgrade",weaponUpgrade)
	http.HandleFunc("/weapon_set",setCurrentWeapon)
	http.HandleFunc("/inherence_upgrade",inherenceUpgrade)
	http.HandleFunc("/revive",revive)
	http.HandleFunc("/login_bonus_obtain",loginBonusObtain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}