package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 根
var root = cobra.Command{
	Short: "excel to db",
	Long:  "read excel file to mysql db database",
}

// 当前工具版本号
var versionCmd = cobra.Command{
	Use:   "version",
	Short: "get excel to db tool version",
	Long:  "get excel to db tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

// 设置excel文档转成数据库数据参数
var cfgParam = cobra.Command{
	Use:   "cfgParam",
	Short: "set excel to db tool param",
	Long:  "Configure Excel as the parameter of the db tool. The parameter information can be found in https://github.com/miajio/excel_to_db/blob/master/README.md",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// 读取excel文件并转换成数据库数据
var readExcel = cobra.Command{
	Use:   "readExcel",
	Short: "read excel to db database",
	Long:  "read excel to db database",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// 读取数据库并将数据写出为excel
var readDb = cobra.Command{
	Use:   "readDb",
	Short: "read db to excel",
	Long:  "read db to excel",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Run 启动
func Run() error {
	root.AddCommand(&versionCmd, &cfgParam, &readExcel, &readDb)
	return root.Execute()
}
