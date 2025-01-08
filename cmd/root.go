package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shadowcloud",                              //命令的名称
	Short: "shadowcloud is a distributed file system", //命令的简短描述
	Long:  `shadowcloud is a distributed file system`, //命令的详细描述

	//Run 属性是一个函数，当执行命令时会调用此函数。
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shadowcloud is running")
	},
}

// rootCmd.Execute() 是命令的执行入口，其内部会解析 os.Args[1:] 参数列表（默认情况下是这样，也可以通过 Command.SetArgs 方法设置参数），然后遍历命令树，为命令找到合适的匹配项和对应的标志。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
