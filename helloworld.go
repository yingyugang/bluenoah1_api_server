package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

var test1 int = 10
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

	test1++
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "dddd", test1)
	//fmt.Fprint(w,"asdasdasd",r)
}

func signinViewHandler(w http.ResponseWriter, r *http.Request) {
	 //var userID = r.Form.Get("user_id")

}

func loginViewHandler(w http.ResponseWriter, r *http.Request) {
	var uuid = r.URL.Query().Get("uuid")
	fmt.Printf("select fail [%s]",uuid)
	rows,err := db1.Query("select id,user_name from user_info where device_id = ?",uuid);
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
	if len(mapUser) > 0 {
		fmt.Fprintf(w, "1")
	}else{
		var newuuid = createUUID()
		r1, err1 := db1.Exec("insert into user_info (user_name,device_id) values (?,?)","New user",newuuid)
		id, err1 := r1.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err1)
			//return
		}
		fmt.Println("new id: ", id)
		r2, err1 := db1.Exec("insert into hero_info (hero_name,user_id) values (?,?)","New hero",id)
		r3, err1 := r2.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err1)
			return
		}
		fmt.Println("new id: ", r3)

		fmt.Fprintf(w, newuuid)
	}
}

func main() {
	db,err := sql.Open("mysql","root:Yyg810412@tcp(127.0.0.1:3306)/BlueNoah?charset=utf8")
	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}
	db1 = db
	http.HandleFunc("/signin", signinViewHandler)
	http.HandleFunc("/login", loginViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}