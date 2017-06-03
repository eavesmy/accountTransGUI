package main

import (
	"./netMgr"
	"./user"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const CONFIG_PATH = "./config.txt"
const ERROR_CONFIGERROR = "配置文件格式不正确"
const ERROR_NOCONFIG = "未找到配置文件文件"
const ERROR_NET = "网络发生问题，请联系管理员查看"
const INFO_DONE = "转移完成"
const URL_PATH = "/sendAll"
const PORT = ":9090"
const HTTPHEAD = "http://"

func init() {
	file, err := os.Open(CONFIG_PATH)

	if err != nil {
		fmt.Println(ERROR_NOCONFIG)
	}

	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println(ERROR_CONFIGERROR)
		}

		URL_STATUS := netMgr.GetUrl(line)

		if URL_STATUS == true {
			return
		}
	}
}

func main() {
	isOK := netMgr.CheckURL()

	if isOK != true {
		fmt.Println(ERROR_CONFIGERROR)
		return
	}

	result := user.Correct()

	if !result {
		return
	}

	ServerInfo := netMgr.GetServer()

	res := netMgr.Client(HTTPHEAD + ServerInfo.ClientServer + PORT + URL_PATH)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(ERROR_NET)
		return
	}

	fmt.Println(string(body))

	defer res.Body.Close()
	return
}
