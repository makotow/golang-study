package main

import (
	"os"
	"regexp"
	"fmt"
	"log"
	//	"strings"
	"bufio"
	"time"
)

type Log struct {
	remotehost  string
	fromidentd  string
	remoteuser  string
	datetime    string
	httprequest string
	httpstatus  string
	databytes   string
	refer       string
	useragent   string
}

// ログを解析するための正規表現今のところ
// apache のログを解析するための正規表現を定義
// 将来的には解析用の関数を渡すようにする。
const (
//	ippattern = `(\\d{1,3}\.\\d{1,3}\.\\d{1,3}\.\\d{1,3})`
//	apacheFormatPattern      = `^(?P<ip>\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}) (?P<fromidentd>\S+) (?P<remoteuser>\S+) [(?P<datetime>[\w:/]+\s[+\-]\d{4})] \"(?P<request>\S+)\" (?P<status>\d{3}) (?P<bytes>\d+) \"(?P<referer>[^\"].+)\" \"(?P<agent>[^\"].+)\"$`
	apacheFormatSplitPattern = `"([^"]+)"|(\[[^\]]+\])|(\S+)`
)

func extractLog(line string) Log {
	// 128.159.142.122 - - [28/Jun/2014:02:20:48 +0900] "GET /category/books HTTP/1.1" 200 97 "/item/finance/113" "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; YTB730; GTB7.2; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET4.0C; .NET4.0E; Media Center PC 6.0)"
	// TODO 本当は各項目を正規表現で取得したいけどうまく取れないので一時的にsplitすることとする。
	//	appacheLogFormatregex := regexp.MustCompile((apacheFormatPattern))
	apacheLogSplitRegex := regexp.MustCompile((apacheFormatSplitPattern))

	mathed := apacheLogSplitRegex.FindAllString(line, -1)

	if mathed == nil {
		fmt.Println("does not much")
		return Log{}
	}

	log := Log {
		remotehost : mathed[0],
		fromidentd : mathed[1],
		remoteuser : mathed[2],
		datetime : mathed[3],
		httprequest : mathed[4],
		httpstatus : mathed[5],
		databytes : mathed[6],
		refer : mathed[7],
		useragent : mathed[8],
	}
	return log
}

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal("could not open file.", err)
		}
		defer fp.Close()
	}

	begin := time.Now()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLog := extractLog(line)
		fmt.Println(parsedLog)
	}
	fmt.Println("Elapsed time: ", time.Now().Sub(begin))
}
