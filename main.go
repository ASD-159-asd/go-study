package main

import (
	"encoding/base64"
	"fmt"
)

func bianma() {
	var s string
	for {
		fmt.Println("请输入编码的字符串 返回用back")
		fmt.Scanf("%s", &s)
		if s == "back" {
			break
		}
		encodeString := base64.StdEncoding.EncodeToString([]byte(s))
		fmt.Println(encodeString)

	}
}

func jiema() {
	var s string
	for {
		fmt.Println("请输入解码的字符串 返回用back")
		fmt.Scanf("%s", &s)
		if s == "back" {
			break
		}
		s, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			fmt.Println("解码错误")
			continue
		}
		fmt.Println(string(s))
	}
}

func main() {
	S := 'v'
	for {
		fmt.Println("编码请输入1 解码请输入2 结束请输入e ")
		fmt.Scanf("%c", &S)
		if S == '1' {
			bianma()
		} else if S == '2' {
			jiema()
		} else if S == 'e' {
			break
		} else {
			fmt.Println("请正确输入!")

		}

	}
}
