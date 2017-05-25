package user

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

var passwd = ""

func InputPasswd() {
	fmt.Print("请输入密码：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		passwd = scanner.Text()

		passwdEncrypt()
		return
	}
}

func GetPasswd() string {
	return passwd
}

func passwdEncrypt() {
	c := md5.New()

	c.Write([]byte(passwd))

	passwdHexed := c.Sum(nil)

	passwd = hex.EncodeToString(passwdHexed)
}
