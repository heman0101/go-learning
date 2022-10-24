package main

import (
	"bufio"
	"flag" //进行命令行解析
	"fmt"
	"go-learning/homework10-17/subject1/core"
	"os"
)

var version = flag.Bool("v", false, "client version")

const (
	statement = "1" // 一言语录
	dogDiary  = "2" // 舔狗日记
)

func main() {
	// 调用flag.Parse()解析命令行参数到定义的flag
	// 非flag命令行参数是指不满足命令行语法的参数，
	// 如命令行参数为cmd --flag=true abc则第一个非flag命令行参数为“abc”
	flag.Parse()
	if *version {
		fmt.Println("version: 0.0.1")
		return
	}
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("键入指令: ")
		if !s.Scan() {
			return
		}

		var (
			result string
			err    error
		)
		// 执行命令
		switch s.Text() {
		case statement:
			result, err = core.Statement()
		case dogDiary:
			result, err = core.Diary()
		case "quit", "exit", "3":
			fmt.Println("\n再见")
			return
		default:
			fmt.Println("命令不存在")
		}

		// 处理 err
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	}
}
