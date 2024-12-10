package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	BaseUrl  = "https://www.dida365.com"
	ApiUrl   = BaseUrl + "/api/v2/task"
	LoginUrl = BaseUrl + "/api/v2/user/signon?wc=true&remember=true"
	Password = "QzU4MjcxMzQxMCtsaQ=="
)

type UserConfig struct {
	UserName string
	PassWord string
}

type TickTickReq struct {
	Url    string
	Cookie string
}

func GetBaseRequest(url, body string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:36.0) Gecko/20100101 Firefox/36.0")
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.7,en;q=0.3")
	req.Header.Add("Referer", BaseUrl)
	req.Header.Add("DNT", "1")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept-Encoding", "deflate")

	return req, nil
}

func (t *TickTickReq) Login() {

	//req, err := GetBaseRequest(LoginUrl)

}

func main() {
	app := &cli.App{
		Name:  "TickTick",
		Usage: "use to create TickTick task",
		Action: func(context *cli.Context) error {
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
