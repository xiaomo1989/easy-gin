package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// 定义 job 命令
var jobCmd = &cobra.Command{
	Use:   "test",
	Short: "运行定时任务",
	Long:  "这个命令可以执行一个定时任务，例如清理数据库、发送邮件等",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("定时任务开始执行...")
		time.Sleep(2 * time.Second) // 模拟任务执行
		fmt.Println("定时任务执行完成！")
	},
}

func init() {
	RootCmd.AddCommand(jobCmd)
}
