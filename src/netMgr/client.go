package netMgr

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var URL_STATUS = false

type SERVERINFO struct {
	ClientServer string
	TargetServer string
}

var ServerInfo SERVERINFO

func GetUrl(url string) bool {

	url = strings.TrimSpace(url)

	if ServerInfo.ClientServer == "" {
		ServerInfo.ClientServer = url

		return URL_STATUS
	}

	ServerInfo.TargetServer = url

	URL_STATUS = true

	return URL_STATUS
}

func CheckURL() bool {
	return ServerInfo.ClientServer != "" && ServerInfo.TargetServer != ""
}

func GetServer() SERVERINFO {
	return ServerInfo
}

func Client(path string) *http.Response {

	values := url.Values{}

	values.Set("a", ServerInfo.ClientServer)
	values.Set("b", ServerInfo.TargetServer)

	fmt.Println("正在处理...")

	res, err := http.PostForm(path, values)

	if err != nil {
		fmt.Println("网络错误： 请检查IP")
	}
	return res

}
