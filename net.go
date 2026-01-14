package main

import (
	"fmt"
	"net"
	"os"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("已经连接:", clientAddr)

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Printf("[%s] 捕获数据: %s\n", clientAddr, string(buf[:n]))
}

func main() {
	var role string
	fmt.Print("输入’z‘抓包 ‘f’发包): ")
	fmt.Scanln(&role)

	switch role {
	case "z":
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println("抓包失败:", err)
			return
		}
		defer listener.Close()
		fmt.Println("抓包启动，监听 8080 端口")
		fmt.Println("输入 quit 并回车即可退出抓包")

		go func() {
			var input string
			for {
				fmt.Scanln(&input)
				if input == "quit" {
					fmt.Println("正在关闭抓包...")
					listener.Close()
					os.Exit(0)
				} else {
					fmt.Println("输入错误，请输入 quit 退出")
				}
			}
		}()

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("抓包已停止")
				return
			}
			go handleClient(conn)
		}

	case "f":
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			fmt.Println("连接服务端失败:", err)
			return
		}
		defer conn.Close()

		data := []byte("Hello Concurrent TCP!")
		conn.Write(data)
		fmt.Println("发送数据:", string(data))

	default:
		fmt.Println("输入错误，输入 z 或 f")
	}
}
