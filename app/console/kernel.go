package console

import (
	"context"
	"framework"
	"framework/command"
	"github.com/spf13/cobra"
)

func RunCommand(container framework.Container) error {
	rootCmd := &cobra.Command{
		// 定义根命令的关键字
		Use: "lsf",
		// 简短介绍
		Short: "lsf 命令",
		// 根命令的详细介绍
		Long: "lsf 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	ctx := context.WithValue(context.Background(), "container", container)
	rootCmd.SetContext(ctx)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)
	// 执行RootCommand
	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	//  demo 例子
	//	rootCmd.AddCommand(demo.InitFoo())
}
