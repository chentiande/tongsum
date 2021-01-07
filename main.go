package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var showVer = flag.Bool("v", false, "show build version")
	var confstr = flag.String("conf", "link.cfg", "汇总配置文件路径，默认为link.cfg")
	var starttime = flag.String("starttime", "", "开始时间，sql中可使用#starttime#代替")
	var endtime = flag.String("endtime", "", "结束时间，sql中可使用#endtime#代替")
	var sumlevel = flag.String("sumlevel", "0", "汇总维度 0-小时，1-天，2-周，3-月，sql中可以使用#sumlevel#代替")
	flag.Parse()

	dbsql, sqlerr := getsql(*confstr)
	if sqlerr != nil {
		fmt.Println("get "+*confstr+" err:", sqlerr)
	}
	dbconf, dberr := getconf("db.conf")
	if dberr != nil {
		fmt.Println("get db.conf err:", dberr)
	}
	db, err := sql.Open("mysql", dbconf.Dbuser+":"+dbconf.Dbpwd+"@tcp("+dbconf.Dbip+":"+dbconf.Dbport+")/"+dbconf.Dbname+"?charset=utf8&allowOldPasswords=1")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if *showVer {

		fmt.Printf("build name:\t%s\n", "tongsum for mysql")
		fmt.Printf("build ver:\t%s\n", "20210106")
		fmt.Printf("build author:\t%s\n", "chentiande")

		os.Exit(0)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)

	for x := range dbsql.Sqls {
		sqlname := dbsql.Sqls[x].Sqlname
		strsql := dbsql.Sqls[x].Sql
		//layout := "2006-01-02 15:04:05"

		if *starttime == "" {
			*starttime = string(time.Now().AddDate(0, 0, -1).Format("2006-01-02")) + " 00:00:00"
		}
		if *endtime == "" {
			*endtime = string(time.Now().AddDate(0, 0, -1).Format("2006-01-02")) + " 23:59:59"
		}
		//time.Sleep(time.Duration(5) * time.Second)
		// just one second

		//替换变量
		strsql = strings.Replace(strsql, "#starttime#", *starttime, -1)
		strsql = strings.Replace(strsql, "#endtime#", *endtime, -1)
		strsql = strings.Replace(strsql, "#sumlevel#", *sumlevel, -1)

		_, err := db.Exec(strsql)
		if err != nil {
			log.Printf("执行sql失败：(%s),sql=%s\n", err, strsql)
			os.Exit(0)
		}
		log.Printf("执行成功：(%s),sql=%s\n", sqlname, strsql)
	}

}
