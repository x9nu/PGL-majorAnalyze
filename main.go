/*
	author：x9nu
	date：2023.5.7
	description：JSON 字符串解析

	步骤：
	1.定义结构体类型，用于存储从JSON解析出来的数据
	2.使用 json.Unmarshal() 函数将JSON数据解析成Go对象。
		- 过程中，可以使用 `tagName` 这样的tag指定JSON键名与Go结构体字段名之间的对应关系

*/
package main

import (
	"MSI_ANALYZE/models"
	"MSI_ANALYZE/utils/excel"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./race.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	bytes, ioErr := ioutil.ReadAll(file)
	if ioErr != nil {
		panic(err.Error())
	}
	// 解析JSON
	var res models.Res
	parseErr := json.Unmarshal([]byte(bytes), &res)
	if parseErr != nil {
		fmt.Println("Error parsing JSON:", parseErr)
	}

	// 输出到xlsx文件中
	excel.WriteToFile("2022PGLmajor.xlsx", res.Result[0].Matchs) // json中，result虽然是个数组，但是只有一个Result[0]
}
