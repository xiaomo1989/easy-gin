package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

// 定义 job 命令
var jobCmd = &cobra.Command{
	Use:   "test",
	Short: "运行定时任务",
	Long:  "这个命令可以执行一个定时任务，例如清理数据库、发送邮件等",
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(jobCmd)
}

func run(cmd *cobra.Command, args []string) {

	// 获取传递的参数 "interval"
	intervalStr, _ := cmd.Flags().GetString("interval")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		fmt.Println("无效的间隔时间！使用默认值2秒")
		interval = 2
	}

	// 输出间隔时间并模拟任务执行
	fmt.Printf("定时任务将在 %d 秒后开始执行...\n", interval)
	time.Sleep(time.Duration(interval) * time.Second) // 模拟任务执行
	fmt.Println("定时任务执行完成！")

}
