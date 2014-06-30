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
	"encoding/csv"
)

type Log struct {
	remotehost  string
	fromidentd  string
	remoteuser  string
	datetime    time.Time
	httprequest string
	httpstatus  string
	databytes   string
	refer       string
	useragent   string
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

func (l *Log) show() {
	fmt.Printf("%v %v %v %v %v %v %v %v %v\n",
		l.remotehost,
		l.fromidentd,
		l.remoteuser,
		l.datetime.String(),
		l.httprequest,
		l.httpstatus,
		l.databytes,
		l.refer,
		l.useragent)
}

// WIP ファイル書き出し用
//		とりあえずテスト用に作る
func (l *Log) output(w *csv.Writer) {
	w.Write([]string {
	l.remotehost,
	l.fromidentd,
	l.remoteuser,
	l.datetime.String(),
	l.httprequest,
	l.httpstatus,
	l.databytes,
	l.refer,
	l.useragent})
	w.Flush()
}

func extractLog(line string) Log {
	// 128.159.142.122 - - [28/Jun/2014:02:20:48 +0900] "GET /category/books HTTP/1.1" 200 97 "/item/finance/113" "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; YTB730; GTB7.2; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET4.0C; .NET4.0E; Media Center PC 6.0)"
	// TODO 本当は各項目を正規表現で取得したいけどうまく取れないので一時的にsplitすることとする。
	//	appacheLogFormatregex := regexp.MustCompile((apacheFormatPattern))
	apacheLogSplitRegex := regexp.MustCompile((apacheFormatSplitPattern))
	matched := apacheLogSplitRegex.FindAllString(line, -1)

	if matched == nil {
		fmt.Println("does not much")
		return Log{}
	}

	log := Log {
		remotehost : matched[0],
		fromidentd : matched[1],
		remoteuser : matched[2],
		datetime : timeParse(matched[3]),
		httprequest : matched[4],
		httpstatus : matched[5],
		databytes : matched[6],
		refer : matched[7],
		useragent : matched[8],
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
	var fop *os.File
	var err error

	if len(os.Args) < 3 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		checkError("could not open file.", err)
		defer fp.Close()

		// ファイルを書き込みモードでオープン(ファイルがなかったら作成する)
		fop, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0666)
		checkError("could not open outputfile", err)
		defer fop.Close()
	}

	begin := time.Now()
	scanner := bufio.NewScanner(fp)
	// Writerを書き込みモードでオープン
	new_headers := []string { "remotehost", "fromidentd", "remoteuser", "datetime", "httprequest", "httpstatus", "databytes" , "refer", "useragent"}
	writer := csv.NewWriter(fop)
	err = writer.Write(new_headers)
	checkError("could not write outputfile", err)

	for scanner.Scan() {
		line := scanner.Text()
		l := extractLog(line)
		writer.Write([]string {l.remotehost,l.fromidentd,l.remoteuser,l.datetime.String(),l.httprequest,l.httpstatus,l.databytes,l.refer,l.useragent})
		writer.Flush()
		//		parsedLog.output(writer)
		//		parsedLog.show()
	}
	fmt.Println("Elapsed time: ", time.Now().Sub(begin))
}
