package main

import (
	"finally-main/mvc"
	"finally-main/runtime"
	"github.com/robfig/cron/v3"
)

func main() {

	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数

	//定时任务1m Getcashbal
	cashbalspec1 := "55 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	crontab.AddFunc(cashbalspec1, mvc.Getcashbal)

	//定时任务1m
	//spec1 := "55 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec1, runtime.Run1)

	//定时任务3m
	//spec3 := "30 2,5,8,11,14,17,20,23,26,29,32,35,38,41,44,47,50,53,56,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec3, runtime.Run3)

	//定时任务1m savecsv
	savecsv := "55 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	crontab.AddFunc(savecsv, runtime.Savecsvfinal)

	//定时任务5m
	//spec5 := "30 4,9,14,19,24,29,34,39,44,49,54,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec5, runtime.Run5)

	//定时任务15m
	spec15 := "0 14,28,43,58 * * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	crontab.AddFunc(spec15, runtime.Run15)

	//定时任务15m Getcashhistory
	cashhistoryspec15 := "30 0,15,30,45 * * * ?" //cron表达式，每15分钟一次
	// 添加定时任务,
	crontab.AddFunc(cashhistoryspec15, mvc.Getcashhistory)

	//定时任务5m GetuplRatio
	specuplRatio := "0 4,9,14,19,24,29,34,39,44,49,54,59 * * * ?" //cron表达式，每4小时一次
	// 添加定时任
	crontab.AddFunc(specuplRatio, mvc.GetuplRatio)

	//定时任务1H
	//spec1H := "55 5 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23 * * ?" //cron表达式，每1小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec1H, runtime.Run1H)

	//定时任务2H
	//spec2H := "55 55 0,2,4,6,8,10,12,14,16,18,20,22 * * ?" //cron表达式，每2小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec2H, runtime.Run2H)

	//定时任务4H
	//spec4H := "55 55 3,7,11,15,19,23 * * ?" //cron表达式，每4小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec4H, runtime.Run4H)

	//定时任务6H
	//spec6H := "55 10 0,6,12,18 * * ?" //cron表达式，每6小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec6H, runtime.Run6H)

	//定时任务12H
	//spec12H := "55 10 0,12 * * ?" //cron表达式，每12小时一次
	// 添加定时任务,
	//crontab.AddFunc(spec12H, runtime.Run12H)

	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer crontab.Stop()
	select {} //阻塞主线程停止

}
