package test

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/fs"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )

// var rootDir, separator string

// var jsonMap map[string]any

// var jsonName = "dir.json"

// func ParseJson(jsonMap map[string]any, wd string) {
// 	for k, v := range jsonMap {
// 		switch v.(type) {
// 		case string:
// 			{
// 				if wd == "" {
// 					wd = rootDir
// 				}
// 				path := wd + separator + v.(string)
// 				CreateDirectory(path)
// 				if k == "text" {
// 					wd = path
// 				}
// 			}
// 		case []any:
// 			{
// 				ParseArray(v.([]any), wd)
// 			}
// 		}
// 	}
// }

// func CreateDirectory(path string) {
// 	fmt.Println(path)
// 	os.MkdirAll(path, fs.ModePerm)
// }

// func ParseArray(array []any, wd string) {
// 	for _, v := range array {
// 		ParseJson(v.(map[string]any), wd)
// 	}
// }

// func loadJson() {
// 	curDir, _ := os.Getwd()
// 	separator = string(filepath.Separator)

// 	rootDir = curDir[:strings.LastIndex(curDir, separator)]
// 	fmt.Println(rootDir)
// 	JsonBytes, _ := os.ReadFile(curDir + separator + jsonName)

// 	err := json.Unmarshal(JsonBytes, &jsonMap)

// 	if err != nil {
// 		panic("json.Unmarshal failed" + err.Error())
// 	}
// }

// func TestGenerateDir(test *testing.T) {
// 	loadJson()
// 	ParseJson(jsonMap, rootDir)
// }
