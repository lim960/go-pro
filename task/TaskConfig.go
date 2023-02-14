package task

import (
	"github.com/robfig/cron/v3"
	"pro/middleware"
)

var Crons = cron.New(cron.WithSeconds())

func Start() {
	//开启定时任务
	Crons.Start()
	//添加工作任务
	//TesJob()
}

// cron格式		秒 分 小时 日 月 星期
// @hourly 		每小时	-- 	0 0 * * * *
// @daily 		每天		-- 	0 0 0 * * *
// @every 10s	每十秒钟

func TesJob() {
	Crons.AddFunc("*/10 * * * * *", func() {
		//Crons.AddFunc("@every 10s", func() {
		middleware.Info("定时任务输出")
	})
}
