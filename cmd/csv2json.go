/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"csv2json/cmd/handle"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// csv2jsonCmd represents the csv2json command
var csv2jsonCmd = &cobra.Command{
	Use:   "csv2json",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reslut := handle.HandleCsvFile(args[0])

		filename := strings.Split(args[0], ".")[0] + ".json"

		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("open file is failed, err: ", err)
		}
		// 延迟关闭
		defer file.Close()

		// 写入UTF-8 BOM，防止中文乱码
		file.WriteString(reslut)
		//生成 json
		// fmt.Println(jsonize.V(v, false))
		// w.Write([]string{strconv.Itoa(i), jsonize.V(v, false)})
		fmt.Println("生成请求文件成功，请查看工具当前目录，文件名为：", filename)
	},
}

func init() {
	rootCmd.AddCommand(csv2jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csv2jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csv2jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
