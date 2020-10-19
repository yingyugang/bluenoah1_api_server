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
	Critical int
	Current_w int
	Ak47_lvl int
	M16_lvl int
	Scatter_lvl int
	Firegun_lvl int
	Rpg_lvl int
	Laserx_lvl int
	Awp_lvl int
}

type BattleResult struct {
	Stage int
	Item1 int64
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
	var item1 int64
	rows,err := db1.Query("select id,stage,item1,current_w,ak47_lvl,m16_lvl,scatter_lvl,firegun_lvl,rpg_lvl,laserx_lvl,awp_lvl from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&id,&stage,&item1,&current_w,&ak47_lvl,&m16_lvl,&scatter_lvl,&firegun_lvl,&rpg_lvl,&laserx_lvl,&awp_lvl)
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
	user := User{Uuid:uuid,HeroId:heroId,Atk:atk,Def:def,MaxHp:maxHp,Hp:hp,MaxSp:maxSp,Sp:sp,Level:level,Stage:stage,Item1:item1,Critical: critical,Current_w:current_w,Ak47_lvl:ak47_lvl,M16_lvl:m16_lvl,Scatter_lvl:scatter_lvl,Firegun_lvl:firegun_lvl,Rpg_lvl:rpg_lvl,Laserx_lvl:laserx_lvl,Awp_lvl:awp_lvl }
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
	var coin = int64(math.Floor (70 * math.Pow(1.14,	float64(lvl + 1))))
	fmt.Printf("coin:[%s]\n",math.Pow(1.14,	float64(lvl + 1)))
	fmt.Printf("select [%s]\n",lvl)
	if item1 >= coin {
		item1-=coin
		db1.Exec("update user_info  set " + column + " = ?,item1 = ? where device_id = ?",lvl + 1,item1,uuid)
		returnUser(w,uuid)
	}
}

func stageClear(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	//TODO
	var result = r.Header.Get("result")
	fmt.Printf(result)
	var battleResult BattleResult
	json.Unmarshal([]byte(result),&battleResult)

	rows,err := db1.Query("select stage,item1 from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var stage int
	var item1 int64
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&stage,&item1)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	db1.Exec("update user_info  set stage = ?,item1 = ? where device_id = ?",stage + 1,battleResult.Item1 + item1,uuid)
	returnUser(w,uuid)
	//fmt.Fprintf(w,strconv.Itoa(stageint))
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}