package execmachine

import (
	"fmt"
	"os/exec"
)

func Execpython() string{
	cmd := exec.Command("python", "test.py")
	cmd.Dir = "/home/mengya/go/src/aiplat/ppp"
	out, err := cmd.CombinedOutput()
	//var out bytes.Buffer
	//cmd.Stdout = &out
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	return string(out)
}
