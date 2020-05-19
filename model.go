package main

//whr-helen 2019

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	_ "strings"
	"io/ioutil"
	"flag"
	"os"
	"strings"

	"time"
)

//用户输入配置
var (
	db_host    string //数据库地址
	db_port    string //数据库端口
	db_name    string //数据库名称
	db_account string //数据库账号
	db_pwd     string //数据库密码
	db_fix     string //表前缀
	path       string //结构体保存路径
	tables     string //表名称
)
//表类型
type Dbtabels struct {
	Name string `json:"name"`
}

//数据字段类型
type Column struct {
	Columname     string `json:"columname"`
	Columntype    string `json:"columntype"`
	Datatype      string `json:"datatype"`
	Columncomment string `json:"columncomment"`
	Columnkey     string `json:"columnkey"`
	Extra         string `json:"extra"`
}

func init() {
	flag.StringVar(&db_account, "acc", "root", "# Database account")
	flag.StringVar(&db_pwd, "pwd", "508416398599e2d7", "# Database password")
	flag.StringVar(&db_host, "host", "127.0.0.1", "# Database host")
	flag.StringVar(&db_port, "port", "3306", "# Database port")
	flag.StringVar(&db_name, "d", "", "# Database name")
	flag.StringVar(&db_fix, "f", "", "# Database name")
	flag.StringVar(&path, "path", "./models", "# Structure preservation path")
	flag.StringVar(&tables, "t", "all", "# Table name formats such as - t user, rule, config")
}

func main() {
	flag.Parse()
	if db_name == "" {
		flag.Usage()
		return
	}
	tables = convtables(tables) //转换

	fmt.Println("数据库:", db_name)
	fmt.Println("数据库账号:", db_account)
	fmt.Println("数据库密码:", db_pwd)
	fmt.Println("结构体保存路径:", path)
	fmt.Println("结构体保存路径:", db_fix)

	fmt.Println("指定数生成据表:", tables)
	fmt.Println("Automatic Struct Start ...")
	if checkpath(path) {
		fmt.Println("paPath irregularityth  Example：'./models'")
		return
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_account, db_pwd, db_host, db_port, db_name)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println("mysql connect err:", err)
		return
	}
	//获取所有表
	var sqls = "select table_name from information_schema.tables where table_schema=?"
	if tables != "all" {
		sqls = fmt.Sprintf("select table_name from information_schema.tables where table_schema=? and table_name in (%s)", tables)
	}
	rows, err2 := db.Query(sqls, db_name)
	if err2 != nil {
		fmt.Println("mysql query err:", err2)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	table := Dbtabels{}
	for rows.Next() {
		err := rows.Scan(&table.Name)
		fmt.Println("正在生成表结构：", table.Name)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		//获取单个表所有字段
		//`COLUMN_KEY`,`COLUMN_TYPE`
		cloumn_sql := fmt.Sprintf("select column_name columnName,column_type,data_type dataType, column_comment columnComment, column_key columnKey,extra from information_schema.columns where table_name = '%s' and table_schema = (select database()) order by ordinal_position", table.Name)
		cloumns, err3 := db.Query(cloumn_sql)
		if err3 != nil {
			fmt.Println("mysql 获取单个表所有字段 err:", err3)
		}
		defer func() {
			if cloumns != nil {
				cloumns.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		struct_str := fmt.Sprintf("type %s struct { \n", Capitalize(table.Name))
		column := Column{}
		func_str := ""

		themTableName := replaceNth(table.Name, db_fix, "", 1)

		for cloumns.Next() {
			err := cloumns.Scan(&column.Columname, &column.Columntype, &column.Datatype, &column.Columncomment, &column.Columnkey, &column.Extra)
			//类型判断
			//开始拼接字符串
			if err != nil {
				fmt.Printf("Scan failed,err:%v", err)
				return
			}
			struct_str += "    " + Capitalize(column.Columname)

			if column.Datatype == "int" || column.Datatype == "tinyint" {

				struct_str += " *uint "

			} else if column.Datatype == "decimal" {

				struct_str += " float64 "

			} else if column.Datatype == "timestamp" {
				struct_str += " zone.Time "
			} else {

				struct_str += " string "
			}

			//`gorm:"column:id;primary_key;auto_increment"`
			if column.Extra == "auto_increment" {

				struct_str += fmt.Sprintf("`gorm:\"column:%s;primary_key;auto_increment\"`\n", column.Columname)

			} else if column.Datatype == "timestamp" {

				struct_str += fmt.Sprintf("`gorm:\"column:%s;\"` \n", column.Columname)
			} else {
				//`gorm:"column:uid;type:int(11)"`
				struct_str += fmt.Sprintf("`gorm:\"column:%s;type:%s;\"` \n", column.Columname, column.Columntype)
			}

			setAuthor := fmt.Sprintf("/*设置 %s 值*/\n", Capitalize(column.Columname))

			func_str += fmt.Sprintf("%sfunc (%s *%s) Set%sAttribute(value interface{}) interface{} {\n ", setAuthor, themTableName, Capitalize(table.Name), Capitalize(column.Columname))

			func_str += fmt.Sprintf("return %s.%s \n", themTableName, Capitalize(column.Columname))

			func_str += "}\n"

			getAuthor := fmt.Sprintf("/**获取 %s 值\n*/\n", Capitalize(column.Columname))

			func_str += fmt.Sprintf("%sfunc (%s *%s) Get%sAttribute(value interface{}) interface{} {\n ", getAuthor, themTableName, Capitalize(table.Name), Capitalize(column.Columname))

			func_str += fmt.Sprintf("return %s.%s \n", themTableName, Capitalize(column.Columname))

			func_str += "}\n\n"
		}

		//模型
		struct_str += "    model.BaseModel \n}\n\n"

		struct_str += fmt.Sprintf("func (%s *%s) TableName() string{ \n", themTableName, Capitalize(table.Name))

		struct_str += fmt.Sprintf("     return %s.SetTableName(\"%s\") \n", themTableName, themTableName)

		struct_str += "}\n"
		struct_str += func_str

		//TableName end

		Author := fmt.Sprintf("/*\nAuthor:pangxianfei\n*DATE:%s \n*contact:421339244@qq.com\n*/\n", time.Now().Format("2020-05-01 15:04:05"))
		head := fmt.Sprintf("%spackage models \n\n import (\n", Author)

		s1 := "     \"github.com/totoval/framework/model\"\n"
		s2 := "     \"github.com/totoval/framework/helpers/zone\"\n"
		s3 := "  )\n\n"
		model_head := head + s1 + s2 + s3

		//导出文件
		body := model_head + struct_str
		tableName := replaceNth(table.Name, db_fix, "", 1)

		filename := fmt.Sprintf("%s/%s.go", path, tableName)

		//创建文件夹
		error2 := os.MkdirAll(path, os.ModePerm)
		if error2 != nil {
			fmt.Println("midkr path err:", error2)
		}
		err4 := ioutil.WriteFile(filename, []byte(body), 0666)
		if err4 != nil {
			fmt.Println("写入文件错误:", err4)
		}
	}

	fmt.Println("End  SUCCESS")
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

//首字母大写 驼峰命名
func Capitalize(name string) string {
	if name == "" {
		return ""
	}
	//去除表前缀
	name = replaceNth(name, db_fix, "", 1)
	temp := strings.Split(name, "_")

	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	return s
}

//判断用户输入
func checkpath(path string) bool {
	a := strings.Split(path, "/")
	if len(a) >= 2 && a[1] != "" {
		return false
	} else {
		return true
	}
}

//用户输入转换
func convtables(tab string) string {
	if tab == "all" {
		return tab
	} else {
		str_arr := strings.Split(tab, ",")
		var tabs string
		for _, v := range str_arr {
			if v != "" {
				item := fmt.Sprintf("'%s',", v)
				tabs += item
			}
		}
		return tabs[0 : len(tabs)-1]
	}
}
