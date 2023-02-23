package app

import (
	"errors"
	"flag"
	"framework"
	"os"
	"path/filepath"
)

type LsfApp struct {
	container  framework.Container //服务容器
	baseFolder string              // 基础路径
}

func (app LsfApp) Version() string {
	return "0.0.1"
}

func (app *LsfApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	file, err := os.Getwd()
	if err == nil {
		return file + "/"
	}
	return ""
}

// ConfigFolder  表示配置文件地址
func (app LsfApp) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (app LsfApp) LogFolder() string {
	return filepath.Join(app.StorageFolder(), "log")
}

func (app LsfApp) HttpFolder() string {
	return filepath.Join(app.BaseFolder(), "http")
}

func (app LsfApp) ConsoleFolder() string {
	return filepath.Join(app.BaseFolder(), "console")
}

func (app LsfApp) StorageFolder() string {
	return filepath.Join(app.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (app LsfApp) ProviderFolder() string {
	return filepath.Join(app.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (app LsfApp) MiddlewareFolder() string {
	return filepath.Join(app.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (app LsfApp) CommandFolder() string {
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app LsfApp) RuntimeFolder() string {
	return filepath.Join(app.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (app LsfApp) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}

// 初始化HadeApp
func NewLsfApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &LsfApp{baseFolder: baseFolder, container: container}, nil
}
