package main

import (
	"fmt"
	"sync"
)

type Application struct {
	CfgProvider ConfigProvider
}
// app 必须要使用 InitApplication 来创建
var App *Application
var AppOnce sync.Once

type AppOption func(app *Application) error

func InitApplication(opts... AppOption) error{
	var err error
	AppOnce.Do(func() {
		App = &Application{
			CfgProvider: NewInMemoryConfigProvider(),  // 默认从内存中加载配置
		}
		for _,opt := range opts {
			err = opt(App)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	return err
}
func WithCfgProvider(cfg ConfigProvider) AppOption{
	return func(app *Application) error {
		app.CfgProvider = cfg   // 修改app默认配置为传入的设置
		return nil
	}
}

