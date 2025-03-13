package excel

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/utils"
	"github.com/spf13/cast"
)

// 将账户信息导出为 txt 文件
func ExportStringsToTxt(dataList []string) (fileName string, err error) {

	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	// filePath := iniconf.String("sys::file_path")
	fileDir := iniconf.String("sys::sys_path")

	if !utils.FileExists(fileDir) {
		os.MkdirAll(fileDir, os.ModePerm)
	}

	fileName = fileDir + "/" + cast.ToString(time.Now().Unix()) + ".txt"

	// 创建或打开文件
	file, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("无法创建文件: %w", err)
	}
	defer file.Close()

	// 创建写入器
	writer := bufio.NewWriter(file)

	// // 写入表头
	// _, err = writer.WriteString("ID\tName\tBalance\tCurrency\n")
	// if err != nil {
	// 	return fmt.Errorf("写入表头失败: %w", err)
	// }

	// 写入数据
	for idx, data := range dataList {
		_, err := writer.WriteString(data)
		if idx < len(dataList)-1 {
			_, err = writer.WriteString("\n")
		}

		if err != nil {
			return "", fmt.Errorf("写入数据失败: %w", err)
		}
	}

	// 刷新缓冲区，确保数据写入文件
	err = writer.Flush()
	if err != nil {
		return "", fmt.Errorf("刷新缓冲区失败: %w", err)
	}

	// err = file.Save(filename)
	//如果没错, 等待180s以后删除
	go RemoveExcel(fileName)
	return fileName, nil
}
