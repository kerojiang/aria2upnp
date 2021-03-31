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
	"strconv"
	"testing"
)

func TestConfFile_GetBTPort(t *testing.T) {

	file1 := "/home/jie/.config/aria2/aria2.conf"

	cf1 := NewConfFile(file1)
	if cf1 == nil {
		fmt.Println("测试1,初始化conffile失败")
		t.Fail()
	}

	fmt.Println(strconv.Itoa(cf1.GetBTPort()))
	//fmt.Println(strconv.Itoa(cf1.GetDHCPort()))

	file2 := "/home/jie/.config/aria2/aria2.sh"
	cf2 := NewConfFile(file2)

	if cf2 == nil {
		fmt.Println("测试2,初始化conffile失败")
		t.Fail()
	}
	if cf2.GetBTPort() != 0 {
		fmt.Println(strconv.Itoa(cf2.GetBTPort()))
	} else {
		t.Fail()
	}

}

func TestConfFile_GetDHCPort(t *testing.T) {

}
