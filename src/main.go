package main

import (
	"./netMgr"
	"./user"
	"bufio"
	"fmt"
	"io"
	"os"
)

const CONFIG_PATH = "./config.txt"
const ERROR_CONFIGERROR = "配置文件格式不正确"

func init() {
	file, err := os.Open(CONFIG_PATH)

	if err != nil {
		fmt.Println("未找到配置文件文件")
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

	user.InputPasswd()

	fmt.Println("GO")
}
