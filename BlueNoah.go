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
	"math/rand"
	"net/http"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}

type UpgradeData struct {
	Coin float64
	Diamond int64
	GreenGear int64
	BlueGear int64
	PurpleGear int64
	OrangeGear int64
	SupperGear int64
}

//注意，首字母一定要大写
type User struct {
	Uuid string
	UserName string
	HeroId int
	Atk int
	Def int
	MaxHp int
	Hp int
	MaxSp int
	Sp int
	Level int
	Stage int
	Item1 float64
	Item2 int64
	Item3 int64
	Item4 int64
	Item5 int64
	Item6 int64
	Rescue int64
	Critical int
	Current_w int
	Ak47_lvl int
	M16_lvl int
	Scatter_lvl int
	Laser_lvl int
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
	Coin_up int
	Buff_up int
	Diamond_count int
	Dodge_up int
	Lastday int
	Loginday int
	Bonus int
	Shop_item_1 int
	Shop_item_2 int
}

type BattleResult struct {
	Stage int
	Item1 float64
	Item2 int64
	Item3 int64
	Item4 int64
	Item5 int64
	Item6 int64
	Rescue int64
	Clear int
}

var db1 *sql.DB

var tokenMap map[string]string

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

func returnUser(w http.ResponseWriter,uuid string){
	var id,stage,heroId,atk,def,maxHp,hp,maxSp,sp,level,critical,current_w,ak47_lvl,m16_lvl,scatter_lvl,laser_lvl,firegun_lvl,rpg_lvl,laserx_lvl,awp_lvl int
	var atk_up,atk_speed_up,critical_up,speed_up,atk_boss_up,hp_up,diamond_count,dodge_up,coin_up,buff_up int
	var lastday,loginday,bonus,shop_item_1,shop_item_2 int
    var userName string
	var item2,item3,item4,item5,item6,rescue int64
	var item1 float64
	rows,err := db1.Query("select id,user_name,stage,item1,item2,item3,item4,item5,current_w,ak47_lvl,m16_lvl,scatter_lvl,firegun_lvl,rpg_lvl,laserx_lvl,awp_lvl,atk_up,atk_speed_up,critical_up,speed_up,atk_boss_up,hp_up,diamond_count,dodge_up,lastday,loginday,bonus,shop_item_1,shop_item_2,item6,laser_lvl,coin_up,buff_up,rescue  from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&id,&userName,&stage,&item1,&item2,&item3,&item4,&item5,&current_w,&ak47_lvl,&m16_lvl,&scatter_lvl,&firegun_lvl,&rpg_lvl,&laserx_lvl,&awp_lvl,&atk_up,&atk_speed_up,&critical_up,&speed_up,&atk_boss_up,&hp_up,&diamond_count,&dodge_up,&lastday,&loginday,&bonus,&shop_item_1,&shop_item_2,&item6,&laser_lvl,&coin_up,&buff_up,&rescue)
		if err != nil{
			fmt.Printf("returnUser:get user info error [%s]",err)
		}
		break
	}
	rows.Close()
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
	rows1.Close()
	user := User{Uuid:uuid,UserName:userName,HeroId:heroId,Atk:atk,Def:def,MaxHp:maxHp,Hp:hp,MaxSp:maxSp,Sp:sp,Level:level,Stage:stage,Item1:item1,Item2:item2,Item3:item3,Item4:item4,Item5:item5,Critical: critical,Current_w:current_w,Ak47_lvl:ak47_lvl,M16_lvl:m16_lvl,Scatter_lvl:scatter_lvl,Firegun_lvl:firegun_lvl,Rpg_lvl:rpg_lvl,Laserx_lvl:laserx_lvl,Awp_lvl:awp_lvl,Atk_up:atk_up,Atk_speed_up:atk_speed_up,Critical_up:critical_up,Speed_up:speed_up,Atk_boss_up:atk_boss_up,Hp_up:hp_up,Diamond_count:diamond_count,Dodge_up:dodge_up,Lastday:lastday,Loginday:loginday,Bonus:bonus,Shop_item_1:shop_item_1,Shop_item_2:shop_item_2 ,Item6: item6,Laser_lvl: laser_lvl,Coin_up: coin_up,Buff_up: buff_up,Rescue: rescue}
	result,err := json.Marshal(user)
	fmt.Println(string(result) )
	fmt.Println(db1.Stats().OpenConnections)
	w.Write(result)

}


func returnNull(w http.ResponseWriter){
	fmt.Fprintf(w,"")
}

func StageClear(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	//TODO
	var result = r.Header.Get("result")
	fmt.Printf(result)
	var battleResult BattleResult
	json.Unmarshal([]byte(result),&battleResult)

	rows,err := db1.Query("select stage,item1,item2,item3,item4,item5,item6,rescue from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}
	var stage int
	var item2,item3,item4,item5,item6,rescue int64
	var item1 float64
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&stage,&item1,&item2,&item3,&item4,&item5,&item6,&rescue)
		if err != nil{
			fmt.Printf("StageClear [%s]",err)
		}
		break
	}
	rows.Close()
	if battleResult.Stage == stage && battleResult.Clear == 1{
		db1.Exec("update user_info  set stage = ?,item1 = ?,item2 = ?,item3 = ?,item4 = ?,item5 = ?,item6 = ?,rescue = ? where device_id = ?",stage + 1,battleResult.Item1 + item1,battleResult.Item2 + item2,battleResult.Item3 + item3,battleResult.Item4 + item4,battleResult.Item5 + item5,battleResult.Item6 + item6,battleResult.Rescue + rescue,uuid)
	}else{
		db1.Exec("update user_info  set item1 = ?,item2 = ?,item3 = ?,item4 = ?,item5 = ?,item6 = ?,rescue = ? where device_id = ?",battleResult.Item1 + item1,battleResult.Item2 + item2,battleResult.Item3 + item3,battleResult.Item4 + item4,battleResult.Item5 + item5,battleResult.Item6 + item6,battleResult.Rescue + rescue,uuid)
	}
	returnUser(w,uuid)
}

func SetCurrentWeapon(w http.ResponseWriter, r *http.Request){
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

func Revive(w http.ResponseWriter, r *http.Request){
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
			fmt.Printf("Revive [%s]",err)
		}
		break
	}
	rows.Close()
	if diamond_count >= 10{
		db1.Exec("update user_info  set diamond_count = ? where device_id = ?",diamond_count - 10,uuid)

		returnUser(w,uuid)
	}else{
		returnNull(w)
	}
}

func LoginBonusObtain(w http.ResponseWriter, r *http.Request){
	var uuid = r.Header.Get("uuid")
	var ads = r.Header.Get("ads")
	var lastday,loginday,bonus int
	var item2,item3,item4,item5,diamond_count int64
	var item1 float64
	rows,err := db1.Query("select item1,item2,item3,item4,item5,diamond_count,lastday,loginday,bonus  from user_info where device_id = ?",uuid)
	if err != nil{
		fmt.Printf("returnUser:select fail [%s]",err)
	}
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&item1,&item2,&item3,&item4,&item5,&diamond_count,&lastday,&loginday,&bonus)
		if err != nil{
			fmt.Printf("LoginBonusObtain:get user info error [%s]",err)
		}
		break
	}
	rows.Close()
	if bonus == 0{
		var day = (loginday -1) % 7
		var week = (loginday -1) / 7
		var multi int64 = 1
		if ads == "1"{
			multi = 2
		}
		switch day {
		case 0:
			var coin = math.Floor (1000000 * math.Pow(1.14,float64(week)))
			db1.Exec("update user_info set item1 = ?,bonus = ? where device_id = ?",item1 + coin * float64(multi),1,uuid)
			fmt.Printf(uuid)
			break
		case 1:
			var diamond int64 = 11
			db1.Exec("update user_info  set diamond_count = ?,bonus = 1 where device_id = ?",diamond_count + diamond * multi,uuid)
			break
		case 2:
			var coin = math.Floor (1000000 * math.Pow(1.14,	float64(week))) * 2
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * float64(multi),uuid)
			break
		case 3:
			var coin = math.Floor (1000000 * math.Pow(1.14,	float64(week))) * 3
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * float64(multi),uuid)
			break
		case 4:
			var diamond int64 = 20
			db1.Exec("update user_info  set diamond_count = ?,bonus = 1 where device_id = ?",diamond_count + diamond * multi,uuid)
			break
		case 5:
			var coin = math.Floor (1000000 * math.Pow(1.14,	float64(week)))* 5
			db1.Exec("update user_info  set item1 = ?,bonus = 1 where device_id = ?",item1 + coin * float64(multi),uuid)
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

var blueNoahRand *rand.Rand

func main() {
	//db,err := sql.Open("mysql","root:810412@tcp(35.187.200.112:3306)/BlueNoah?charset=utf8")
	db,err := sql.Open("mysql","admin:810412612@tcp(bluenoah.cdcgxd165efz.ap-northeast-1.rds.amazonaws.com:3306)/BlueNoah?charset=utf8")
	//db,err := sql.Open("mysql","admin:Yingyugang2017@tcp(bluenoah-dev.cdcgxd165efz.ap-northeast-1.rds.amazonaws.com:3306)/BlueNoah?charset=utf8")

	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}
	db1 = db
	db.SetMaxOpenConns(65)
	db.SetMaxIdleConns(65)
	tokenMap = make(map[string]string)
	blueNoahRand = rand.New(rand.NewSource(99))
	http.HandleFunc("/login", LoginViewHandler)
	http.HandleFunc("/stage_clear", StageClear)
	http.HandleFunc("/weapon_upgrade", WeaponUpgrade)
	http.HandleFunc("/weapon_upgrade_bulk", WeaponUpgradeBulk)
	http.HandleFunc("/weapon_set", SetCurrentWeapon)
	http.HandleFunc("/inherence_upgrade", InherenceUpgrade)
	http.HandleFunc("/revive",Revive)
	http.HandleFunc("/login_bonus_obtain", LoginBonusObtain)
	http.HandleFunc("/look_ads_with_shop", LookAdsWithShop)
	http.HandleFunc("/purchase",Purchase)
	http.HandleFunc("/change_name",ChangeUserName)
	http.HandleFunc("/registerToken",RegisterPushToken)
	http.HandleFunc("/push",PushNotification)

	log.Fatal(http.ListenAndServe(":8080", nil))
}