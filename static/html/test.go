package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 读取eth.txt文件的内容
	ethContent, err := os.ReadFile("eth.txt")
	if err != nil {
		panic(err)
	}

	// 读取index.js文件的内容
	indexContent, err := os.ReadFile("index.js")
	if err != nil {
		panic(err)
	}

	// 将字节切片转换为字符串
	indexContentStr := string(indexContent)

	// 找到splitData()函数中的数据
	startIndex := strings.Index(indexContentStr, "splitData([")
	endIndex := strings.Index(indexContentStr, "])")
	if startIndex == -1 || endIndex == -1 {
		panic("splitData data not found")
	}
	startIndex += len("splitData([")

	// 替换splitData()函数中的数据为eth.txt文件的内容
	indexcs := indexContentStr[:startIndex]
	indexce := indexContentStr[endIndex:]
	newContent := indexcs[1:] + string(ethContent) + indexce[:len(indexce)-1]
	fmt.Println(newContent)
	// 将修改后的内容写回index.js文件
	//err = os.WriteFile("index.js", []byte(newContent), os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
}
