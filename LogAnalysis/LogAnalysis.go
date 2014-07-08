// TODO 全体的にテストを書く。

package main

import (
	"os"
	"regexp"
	"fmt"
	"log"
	"strings"
	"bufio"
	"time"

	"path/filepath"

	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Log struct {
	Id          int64
	Remotehost  string
	Fromidentd  string
	Remoteuser  string
	Datetime    string
	Httprequest string
	Httpstatus  string
	Databytes   string
	Refer       string
	Useragent   string
}

// ログを解析するための正規表現
// 今のところapache のログを解析するための正規表現を定義
// 将来的には解析用の関数を渡すようにする。
const (
	// WIP apacheLogFormatpattern   = `"^/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`
	apacheFormatSplitPattern = `"([^"]+)"|(\[[^\]]+\])|(\S+)`
)

//　日付フォーマット
//  指定の仕方はtime.Parseのconstを見て理解
const (
	timeformat = "02/Jan/2006:15:04:05 -0700"
)

var (
	Driver        = "sqlite3"
	HOME          = os.Getenv("HOME")
	DataSourceDir = ""
	DataSource    = ""
)

func init() {
	fmt.Println("calling init func.")

	if HOME == "" {
		HOME = "/tmp"
	}
	DataSourceDir = filepath.Join(HOME, "tmp", "db")
	if err := os.MkdirAll(DataSourceDir, 0755); err != nil {
		checkError("create failed dir ", err)
	}
	DataSource = filepath.Join(DataSourceDir, "log.db")
}

func initDb() *gorp.DbMap {
	db, err := sql.Open(Driver, DataSource)
	checkError(" sql.Open failed ", err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Log{}, "Log").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkError(" Create tables failed ", err)
	return dbmap
}

func extractLog(line string) Log {
	// 128.159.142.122 - - [28/Jun/2014:02:20:48 +0900] "GET /category/books HTTP/1.1" 200 97 "/item/finance/113" "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; YTB730; GTB7.2; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET4.0C; .NET4.0E; Media Center PC 6.0)"
	// TODO 本当は各項目を正規表現で取得したいけどうまく取れないので一時的にsplitすることとする。
	//	appacheLogFormatregex := regexp.MustCompile((apacheFormatPattern))
	apacheLogSplitRegex := regexp.MustCompile((apacheFormatSplitPattern))
	matched := apacheLogSplitRegex.FindAllString(line, -1)

	if matched == nil {
		log.Fatalln("does not much")
		return Log{}
	}

	log := Log {
		Remotehost : matched[0],
		Fromidentd : matched[1],
		Remoteuser : matched[2],
		Datetime : strings.Trim(strings.Trim(matched[3], "["), "]"),
		Httprequest : strings.Trim(matched[4], "\""),
		Httpstatus : matched[5],
		Databytes : matched[6],
		Refer : strings.Trim(matched[7], "\""),
		Useragent : strings.Trim(matched[8], "\""),
	}
	return log
}

func timeParse(datetime string) time.Time {

	// TODO ログを分解した時点で消す
	datetime = strings.Trim(datetime, "[\"")
	datetime = strings.Trim(datetime, "\"]")
	t, err := time.Parse(timeformat, datetime)
	checkError("parsing time formt error", err)
	return t
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		checkError("could not open file.", err)
		defer fp.Close()
	}

	begin := time.Now()
	scanner := bufio.NewScanner(fp)
	dbmap := initDb()
	defer dbmap.Db.Close()

	tx, _ := dbmap.Begin()
	for scanner.Scan() {
		line := scanner.Text()
		l := extractLog(line)
		tx.Insert(&l)
	}
	tx.Commit()

	fmt.Println("Elapsed time: ", time.Now().Sub(begin))
}
