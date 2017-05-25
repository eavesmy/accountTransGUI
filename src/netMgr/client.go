package netMgr

import (
	"fmt"
	"net/http"
)

var URL_STATUS = false

type SERVERINFO struct {
	clientServer string
	targetServer string
}

var ServerInfo SERVERINFO

func GetUrl(url string) bool {

	if ServerInfo.clientServer == "" {
		ServerInfo.clientServer = url

		return URL_STATUS
	}
	ServerInfo.targetServer = url

	URL_STATUS = true

	return URL_STATUS
}

func CheckURL() bool {
	return ServerInfo.clientServer != "" && ServerInfo.targetServer != ""
}

func GetServer() SERVERINFO {
	return ServerInfo
}

func Client(path string) *http.Response {
	fmt.Println(path)
	res, err := http.Get(path)

	if err != nil {
		fmt.Println("网络错误：", err)
	}

	return res
}
