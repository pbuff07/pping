package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	osType := runtime.GOOS
	var cmd *exec.Cmd

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Error: No Address provided!")
		os.Exit(1)
	}

	target := string(args[0])
	target = strings.Replace(target, "https://", "", -1)
	target = strings.Replace(target, "http://", "", -1)
	target = replaceSpecialCharacter(target)

	switch osType {
	case "windows":
		cmd = exec.Command("ping", "-n", "4", target)
	case "linux", "darwin":
		cmd = exec.Command("ping", "-c", "4", target)
	default:
		fmt.Println("Unsupported operating system.")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	fmt.Println(string(output))
}

// 函数定义确定了返回值也会执行到return再返回最后结果
func replaceSpecialCharacter(Address string) (target string) {
	target = Address // 默认返回值
	ColonIndex := strings.Index(Address, ":")
	HyphenIndex := strings.Index(Address, "/")
	// 首先处理: 因为端口号在路径之前
	if ColonIndex != -1 {
		return Address[:ColonIndex]
	}
	if HyphenIndex != -1 {
		return Address[:HyphenIndex]
	}
	return
}
