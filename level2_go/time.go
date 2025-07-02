package main

import (
	"fmt"
	"time"
)

func main() {
	var timestamp int64 = 1587880013             //时间戳
	t := time.Unix(timestamp, 0)                 //日期对象
	fmt.Println(t.Format("2006-01-02 03:04:05")) //日期格式化输出
}
