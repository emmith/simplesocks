package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

var baidu iDict
var caiyun iDict

func init() {
	baidu, _ = getDict("baidu")
	caiyun, _ = getDict("caiyun")
}

// 只要一个返回，就输出
func searchAny(word string) {
	log.Printf("----------------------Print the fast search result-----------------------")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		baidu.searchWord(word)
		cancel()
	}()
	go func() {
		caiyun.searchWord(word)
		cancel()
	}()
	<-ctx.Done()
}

// 全部完成后，输出
func searchAll(word string) {
	log.Printf("--------------------Print all search result--------------------")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		baidu.searchWord(word)
		wg.Done()
	}()
	go func() {
		caiyun.searchWord(word)
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	var word string
	fmt.Scanf("%s\n", &word)
	searchAll(word)
	searchAny(word)
}
