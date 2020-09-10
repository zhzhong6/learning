package main
import (
	"fmt"
	"net"
)
func main()  {
	// 建立客户端连接
	conn,err :=net.Dial("tcp","127.0.0.1:8090")
	if err !=nil{
		fmt.Println("建立连接120.0.0.1:8090出错",err)
		return    // 注意 如果连接失败，这里要失败
	}
	fmt.Println("连接120.0.0.1:8090成功！",err)
	defer conn.Close()
	// 发送信息
	whriteStr :="hello word,你好！"
	// 通过conn连接发送信息
	n,err :=conn.Write([]byte(whriteStr))
	if err!=nil{
		fmt.Println("conn whrite err=",err)
	}
	fmt.Println("客户端发送了",n,"个字节并退出")
}