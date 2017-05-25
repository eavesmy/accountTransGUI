package user

import (
	"bufio"
	"fmt"
	"os"
)

var passwd = ""

func InputPasswd() {
	fmt.Print("请输入密码：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		passwd = scanner.Text()
	}
}

func GetPasswd() string {
	return passwd
}
