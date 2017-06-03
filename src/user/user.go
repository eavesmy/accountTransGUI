package user

import (
	"bufio"
	"fmt"
	"os"
)

func Correct() bool {
	fmt.Print("确定要转移吗？(y/n) ")
	scanner := bufio.NewScanner(os.Stdin)

	_b := true

	for scanner.Scan() {
		result := scanner.Text()

		if result != "y" {
			_b = false
		}

		return _b
	}

	return _b
}
