package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	// 在 Go 语言中，_（下划线）通常用作空白标识符，表示忽略导入的包或变量，
	// 以满足 Go 语言的语法要求，同时避免编译器报未使用的变量或包的警告。

	// 在导入包时，如果你不使用导入的包中的任何函数、变量或方法，可以使用 _ 来代替标识符，
	// 示意你明确知道导入但不会使用。这样做的目的可能是为了执行包的初始化代码，
	// 而不实际使用包的导出内容。

	// 具体到问题中，_ "github.com/go-sql-driver/mysql" 
	// 表示导入了 github.com/go-sql-driver/mysql 包，
	// 但在代码中没有直接使用该包的任何内容。通常，这样的导入语句用于执行包的初始化代码，
	// 如注册数据库驱动。
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// 数据库表对应的结构体
type RemindDate struct {
	ID   int    `db:"id"`
	Event string `db:"event"`
	LunarDate string `db:"lunar_date"`
	SolarDate string `db:"solar_date"`
}

func main() {

	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 从环境变量中读取账号密码
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbIp := os.Getenv("DB_IP")
	dbName := os.Getenv("DB_NAME")

	// 构建数据库连接字符串
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbIp, dbName)

	// 连接到 MySQL 数据库
	db, err := sqlx.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 执行查询
	people := []RemindDate{}
	err = db.Select(&people, "SELECT id, event, lunar_date, solar_date FROM remind_date")
	if err != nil {
		panic(err)
	}

	// 打印查询结果
	for _, person := range people {
		fmt.Println(reflect.TypeOf(person))
		fmt.Printf("ID: %d, Event: %s, LunarDate: %s, SolarDate: %s\n", person.ID, person.Event, person.LunarDate, person.SolarDate)
	}
}

/*
docker exec go_c go run demo/mysql.go
*/
