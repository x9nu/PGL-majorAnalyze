package excel

import (
	"MSI_ANALYZE/models"
	"MSI_ANALYZE/utils/time2std"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func WriteToFile(filename string, data []models.Match) {
	fmt.Println("开始写入文件：", filename)
	// 创建新 Excel 文件
	f := excelize.NewFile()
	defer func() { // 以防
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建一个工作表
	sheetName := "比赛详情"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue(sheetName, "A1", "队伍A")
	f.SetCellValue(sheetName, "B1", "队伍B")
	f.SetCellValue(sheetName, "E1", "比赛开始时间")
	// 将数据写入工作表
	for i, row := range data {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row.Team1DTO.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row.Team2DTO.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), time2std.Time2std(row.StartTime))
		// fmt.Println(time2std.Time2std(row.StartTime))

	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	fmt.Println("写入成功，请查看")
}
