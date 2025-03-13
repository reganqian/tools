package excel

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/utils"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ExportExcel(dataList []map[string]interface{}, keys []string, names []string) (filename string, err error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	titleRow := sheet.AddRow()
	for _, name := range names {
		cell := titleRow.AddCell()
		cell.Value = name
	}

	for _, data := range dataList {
		row := sheet.AddRow()
		for _, key := range keys {
			cell := row.AddCell()
			fmt.Println(key)
			fmt.Println(data)
			val := data[key]
			if val == nil {
				val = ""
			}
			cell.Value = fmt.Sprintf("%v", val)
		}
	}
	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	// filePath := iniconf.String("sys::file_path")
	fileDir := iniconf.String("sys::sys_path")

	if !utils.FileExists(fileDir) {
		os.MkdirAll(fileDir, os.ModePerm)
	}
	filename = fileDir + "/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	//如果没错, 等待180s以后删除
	go RemoveExcel(filename)
	return filename, err
}

func RemoveExcel(fileName string) {
	time.Sleep(time.Second * 180)
	os.RemoveAll(fileName)
}

// filepath：对应url， isUseFirst, 是否读取第一行
func ReadExcelToMap(filePath string, isUseFirst bool) (dataList [][]string) {
	// res, err := http.Get("http://kkline.oss-ap-southeast-1.aliyuncs.com/txt/flb.xlsx")
	filePath = strings.ReplaceAll(filePath, "https://", "http://")

	res, err := http.Get(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := excelize.OpenReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	for idx, row := range rows {
		rowsData := []string{}
		if idx == 0 && !isUseFirst {
			continue
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
			rowsData = append(rowsData, colCell)
		}
		fmt.Println()
		dataList = append(dataList, rowsData)
	}
	return dataList
}
