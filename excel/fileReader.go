package excel

import (
	"fmt"
	"io"
	"os"

	// "bufio"
	// "strings"
	"io/ioutil"
	"net/http"
	"strings"
	// "github.com/tealeg/xlsx"
	// log "github.com/sirupsen/logrus"
	// "strconv"
	// "github.com/Luxurioust/excelize"
)

// 通过URl读取文件
func ReadFileByUrl(filePath string) (string, error) {
	res, err := http.Get(filePath)

	fd, err := ioutil.ReadAll(res.Body)

	return string(fd), err
}

// 下载文件
func DownloadFile(url, fileName string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error")
		return
	}
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file error")
		return
	}
	io.Copy(f, res.Body)
}

// 获取文本文件中数据, 按照行读取(\n 或\r换行的数据)
func GetTxtDataList(filePath string) []string {
	data, _ := ReadFileByUrl(filePath)
	datas := strings.Split(data, "\n")
	var dataList []string
	for _, dataStr := range datas {
		datas2 := strings.Split(dataStr, "\r")
		for _, dataStr1 := range datas2 {
			re := strings.Replace(dataStr1, " ", "", -1)
			re = strings.Replace(re, "\r", "", -1)
			if re != "" {
				dataList = append(dataList, re)
			}
		}
	}
	return dataList
}
