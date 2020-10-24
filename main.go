package main

import (
	"demoProject/dao"
	"demoProject/routers"
	"demoProject/setting"
	"fmt"
)

func main() {
	setting.Init()
	err := dao.InitMySQL(setting.Conf)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	r := routers.SetupRouter()
	r.Run()
}



