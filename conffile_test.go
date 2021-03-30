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
	"testing"
)

func TestConfFile_GetBTPort(t *testing.T) {

	file1 := "/home/jie/.config/aria2/aria2.conf"

	cf1 := NewConfFile(file1)
	if cf1 == nil {
		fmt.Println("测试1,初始化conffile失败")
		t.Fail()
	}

	fmt.Fprintln(cf1.GetBTPort())
	fmt.Fprintln(cf1.Get())

	file2 := "/home/jie/.config/aria2/aria2.sh"
	cf2 := NewConfFile(file2)

	if cf2 == nil {
		fmt.Println("测试2,初始化conffile失败")
		t.Fail()
	}

}

func TestConfFile_GetDHCPort(t *testing.T) {
	type fields struct {
		_filepath string
		_content  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := &ConfFile{
				_filepath: tt.fields._filepath,
				_content:  tt.fields._content,
			}
			if got := cf.GetDHCPort(); got != tt.want {
				t.Errorf("ConfFile.GetDHCPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
