package netMgr

import (
	"fmt"
	"net/http"
)

var URL_STATUS = false

var ClientServer = ""
var TargetServer = ""

func GetUrl(url string) bool {

	if ClientServer == "" {
		ClientServer = url

		return URL_STATUS
	}
	TargetServer = url

	URL_STATUS = true

	return URL_STATUS
}

func CheckURL() bool {
	return ClientServer != "" && TargetServer != ""
}

func Client(path string) *http.Response {
	res, err := http.Get(path)

	if err != nil {
		fmt.Println("网络错误：", err)
	}

	return res
}
