package main

import (
	"fmt"
	"github.com/ibuilding-x/driver-box/driverbox"
	"github.com/ibuilding-x/driver-box/driverbox/export"
	"github.com/ibuilding-x/driver-box/driverbox/helper"
	"github.com/ibuilding-x/driver-box/driverbox/plugin"
	"os"
	"time"
)

func main() {
	//关闭driver-box日志
	os.Setenv("LOG_LEVEL", "error")

	//启动driver-box服务
	driverbox.Start([]export.Export{})

	go func() {
		// 创建一个 Ticker，每隔 1 秒执行一次任务
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// 获取设备影子开关状态
				v, _ := helper.DeviceShadow.GetDevicePoint("switch-1", "onOff")
				fmt.Printf("开关状态：%v", v)
				fmt.Println()

				//切换开关状态
				if v == int64(0) {
					_ = driverbox.WritePoint("switch-1", plugin.PointData{
						PointName: "onOff",
						Value:     1,
					})
				} else {
					_ = driverbox.WritePoint("switch-1", plugin.PointData{
						PointName: "onOff",
						Value:     0,
					})
				}
			}
		}
	}()
	select {}
}
