package main

import "fmt"

func getDict(name string) (iDict, error) {
	if name == "baidu" {
		return &BaiduDict{}, nil
	} else if name == "caiyun" {
		return &CaiyunDict{}, nil
	}
	return nil, fmt.Errorf("wrong dictName")
}
