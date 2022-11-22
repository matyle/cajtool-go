package handle

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/pingcap/log"
	"go.uber.org/zap"
)

func HandleCsvFile(path string) string {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open file %s failed: %v", path, err))
	}
	defer file.Close()
	r := csv.NewReader(file)
	header, err := r.Read()
	if err != nil {
		log.Error("read header failed", zap.Error(err))
	}
	if len(header) <= 0 {
		log.Error("header is nil")
		return ""
	}
	log.Debug("header", zap.Strings("header", header))

	// 读取数据
	data := make([][]string, 0, len(header))
	jsonData := make([]string, 0, len(header))

	for {
		record, err := r.Read() //读取一行,有多列
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if len(record) != len(header) {
			log.Error("invalid record", zap.Strings("record", record))
		}
		data = append(data, record)
		//生成json字符串
		var json = "{"

		for i, v := range record {
			if i != len(record)-1 {
				json += fmt.Sprintf("\"%s\":\"%v\",", header[i], v)
			} else {
				json += fmt.Sprintf("\"%s\":\"%v\"", header[i], v)
			}
		}
		json += "}"
		jsonData = append(jsonData, json)
	}

	if len(jsonData) <= 0 {
		log.Error("data is nil")
		return ""
	}

	jsonLen := len(jsonData)
	lastStr := jsonData[jsonLen-1]
	lastStr = strings.TrimSuffix(lastStr, ",")
	jsonData[jsonLen-1] = lastStr
	if len(jsonData) > 1 {
		//开头加上[，最后逗号替换为]
		jsonData[0] = "[" + jsonData[0]
		jsonData[jsonLen-1] = jsonData[jsonLen-1] + "]"
	}

	result := strings.Join(jsonData, ",")
	return result
}
