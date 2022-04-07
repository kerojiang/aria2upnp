/*
 * Copyright (c)  2021  All Rights Reserved
 * 项目名称:arialib
 * 文件名称:file.go
 * 修改日期:2021/03/30 10:01:51
 * 作者:kerojiang
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type ConfFile struct {
	_filepath string
	_content  string
}

// @name:file构造函数
//
// @description:
//
// @param:
//
// @return:
func NewConfFile(filepath string) *ConfFile {
	_, err := os.Stat(filepath)
	//文件不存在
	if err != nil {
		return nil
	} else {
		//文件扩展名不是conf
		if path.Ext(path.Base(filepath)) != ".conf" {
			return nil
		} else {
			return &ConfFile{
				_filepath: filepath,
			}
		}
	}
}

func (cf *ConfFile) readFile() {
	if cf._content == "" {
		f, _ := os.Open(cf._filepath)
		defer f.Close()
		r, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("读取配置文件异常,", err)
		} else {
			cf._content = string(r[:])
		}

	}

}

// @name:获取bt端口
//
// @description:
//
// @param:
//
// @return:
func (cf *ConfFile) GetBTPort() int {

	return cf.findPort("listen-port")

}

// @name:找到指定端口号
//
// @description:
//
// @param:
//
// @return:
func (cf *ConfFile) findPort(portName string) int {
	cf.readFile()

	flag := "\n"
	if strings.Contains(cf._content, "\r\n") {
		flag = "\r\n"
	}

	for _, line := range strings.Split(cf._content, flag) {

		result := strings.Trim(line, " ")

		if strings.HasPrefix(result, portName) {
			//找到指定端口
			port, err := strconv.Atoi(strings.Split(result, "=")[1])
			if err != nil {
				fmt.Println("配置文件端口异常", err)
				return 0
			}

			return port
		}
	}
	return 0
}

// @name:获取dhc端口
//
// @description:
//
// @param:
//
// @return:
func (cf *ConfFile) GetDHCPort() int {
	return cf.findPort("dht-listen-port")
}
