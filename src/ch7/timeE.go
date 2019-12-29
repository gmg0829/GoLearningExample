package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now() //获取当前时间
    fmt.Printf("current time:%v\n", now)
    year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
	second := now.Second() //秒
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	
	var layout string = "2006-01-02 15:04:05"
    var timeStr string = "2019-12-12 15:22:12"
    timeObj1, _ := time.Parse(layout, timeStr)
    fmt.Println(timeObj1)

    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}