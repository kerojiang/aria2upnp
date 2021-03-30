/*
 * Copyright (c)  2021  All Rights Reserved
 * 项目名称:arialib
 * 文件名称:main.go
 * 修改日期:2021/03/30 10:01:06
 * 作者:kerojiang
 */

package main

import (
	"flag"
	"fmt"

	"github.com/prestonTao/upnp"
)

func main() {
	//文件路径
	var filePath string

	flag.StringVar(&filePath, "f", "", "指定的aria2配置文件路径")

	flag.Parse()

	if filePath != "" {
		cf := NewConfFile(filePath)
		if cf != nil {
			btPort := cf.GetBTPort()
			dhcPort := cf.GetDHCPort()
			if btPort != 0 && dhcPort != 0 {
				//upnp端口映射
				mapping := new(upnp.Upnp)

				err := mapping.AddPortMapping(btPort, btPort, "TCP")
				if err != nil {
					if err.Error() == "未发现网关设备" {
						fmt.Println("bt端口映射异常,请检查防火墙设置")
					} else {
						fmt.Println("bt端口映射异常", err)
					}

				} else {
					fmt.Println("bt端口映射成功")
				}

				err = mapping.AddPortMapping(dhcPort, dhcPort, "UDP")
				if err != nil {
					if err.Error() == "未发现网关设备" {
						fmt.Println("dhc端口映射异常,请检查防火墙设置")
					} else {
						fmt.Println("dhc端口映射异常", err)
					}
				} else {
					fmt.Println("dhc端口映射成功")
				}

			} else {
				fmt.Println("读取配置文件异常")
			}

		} else {
			fmt.Println("读取配置文件异常")
		}
	} else {
		fmt.Println("没有指定配置文件路径,请输入-h参数查看帮助")
	}

}
