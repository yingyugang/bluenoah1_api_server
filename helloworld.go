package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}

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
	var id,stage,heroId,atk,def,maxHp,hp,maxSp,sp,level int

	rows,err := db1.Query("select id,stage from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&id,&stage)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}
	rows1, err := db1.Query("select id,atk,def,maxhp,hp,maxsp,sp,level from hero_info where user_id = ?",id)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	for rows1.Next(){
		rows1.Columns()
		err := rows1.Scan(&heroId,&atk,&def,&maxHp,&hp,&maxSp,&sp,&level)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		break
	}


	user := User{Uuid:uuid,HeroId:heroId,Atk:atk,Def:def,MaxHp:maxHp,Hp:hp,MaxSp:maxSp,Sp:sp,Level:level,Stage:stage}
	result,err := json.Marshal(user)
	w.Write(result)
	//fmt.Fprintf(w,uuid)
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

func stageClear(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var stage = r.Header.Get("stage")
	stageint,_ := strconv.Atoi(stage)
	stageint += 1
	db1.Exec("update user_info  set stage = ? where device_id = ?",stageint,uuid)
	fmt.Fprintf(w,strconv.Itoa(stageint))
}

func main() {
	db,err := sql.Open("mysql","root:Yyg810412@tcp(127.0.0.1:3306)/BlueNoah?charset=utf8")
	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}
	db1 = db
	http.HandleFunc("/login", loginViewHandler)
	http.HandleFunc("/stage_clear",stageClear)
	log.Fatal(http.ListenAndServe(":8080", nil))
}