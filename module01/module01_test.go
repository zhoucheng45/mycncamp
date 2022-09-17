/*
第一次作业
*/
package module01

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
编写一个小程序：
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func Test01(t *testing.T) {
	strArr := [5]string{"I", "am", "stupid", "and", "weak"}
	for index, val := range strArr {
		fmt.Printf("索引:%v, 值:%v\n", index, val)
		if val == "stupid" {
			strArr[index] = "smart"
		} else if val == "weak" {
			strArr[index] = "strong"
		}
	}
	fmt.Printf("修改后的数组值：%v", strArr)
}

/**
基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func Test02(t *testing.T) {
	intChan := make(chan int, 10)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			intn := rand.Intn(10)
			println("准备放数据", intn)
			intChan <- intn
		}
	}()
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case val := <-intChan:
			println("从channel中获取到值:", val)
		default:
			now := time.Now()
			fmt.Printf("%v,没有数据\n", now)
		}
	}

}
