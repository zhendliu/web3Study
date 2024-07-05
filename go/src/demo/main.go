package main

import (
	"fmt"
	"task-manage/router"
)

func main() {

	r := router.SetupRouter()

	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", 8000)); err != nil {
		fmt.Println("启动服务失败:", err)
	}

}
