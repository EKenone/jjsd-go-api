package main

import (
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	prefix = flag.String("f", "", "表前缀名称")
	table  = flag.String("t", "sdf", "数据表名称")
	path   = flag.String("p", "", "存放路径")
	cover  = flag.Bool("c", false, "覆盖性生成模型")
)

func main() {
	flag.Parse()

	if *path == "" {
		log.Panicln("未定义存放路径")
	}

	paths := strings.Split(*path, "/")
	pathsLen := len(paths)

	packName := paths[pathsLen-1]
	name := strings.TrimPrefix(*table, *prefix)
	modelName := ToCamelCase(name)
	fileName := *path + "/" + name + ".go"

	_, err := os.Stat(fileName)
	if err == nil && !*cover {
		log.Panicln("文件已存在")
	}

	content, _ := ioutil.ReadFile("tools/build_model/model_demo.tmp")
	fileContent := strings.Replace(string(content), "[pack]", packName, 1)
	fileContent = strings.Replace(fileContent, "[table]", *table, 1)
	fileContent = strings.Replace(fileContent, "[model_name]", modelName, 2)

	fileContent = strings.Replace(fileContent, "[field]", strings.Trim(mysqlModel(), "\r\n"), 1)

	f, err := os.Create(fileName)
	defer f.Close()

	f.WriteString(fileContent)
}

//func TypeChange(s string) string {
//
//}

func mysqlModel() string {
	var st []struct {
		Field   string
		Type    string
		Null    string
		Key     string
		Default string
		Extra   string
	}

	dsn := "jjsd:123456@tcp(11.11.11.114:3306)/jjsd?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	res := db.Raw("DESC " + *table).Scan(&st)
	if res.Error != nil {
		log.Panicln(err)
	}

	base := ""

	for _, v := range st {
		field := v.Field
		if v.Field == "id" {
			field = "ID"
		}

		base += "	" + ToCamelCase(field) + "	"

		if strings.HasPrefix(v.Type, "bigint") {
			if strings.HasSuffix(v.Type, "unsigned") {
				base += "uint64"
			} else {
				base += "int64"
			}
		} else if strings.Index(v.Type, "int") > -1 && !strings.HasPrefix(v.Type, "bigint") {
			if strings.HasSuffix(v.Type, "unsigned") {
				base += "uint"
			} else {
				base += "int"
			}
		} else if strings.HasPrefix(v.Type, "decimal") {
			base += "float64"
		} else {
			base += "string"
		}

		base += "	" + "`gorm:\"column:" + v.Field

		if v.Key == "PRI" {
			base += ";primaryKey"
		}

		base += "\"` \r\n"
	}

	return base
}

func ToCamelCase(str string) string {
	temp := strings.Split(str, "_")
	for i, r := range temp {
		temp[i] = strings.Title(r)
	}

	return strings.Join(temp, "")
}
