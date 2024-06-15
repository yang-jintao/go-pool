package pool

import (
	"fmt"
	"testing"
	"time"
)

func UserInfo(name string, age int) {
	fmt.Println(name, age)
}

func TestPool(t *testing.T) {
	// 创建任务池
	pool, err := NewPool(10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		// 任务放入池中
		pool.Put(&Task{
			Handler: func() {
				UserInfo("yjt", i)
			},
		})
	}

	time.Sleep(3 * time.Second) // 等待执行
	fmt.Println(pool.runningWorkers)
	time.Sleep(3 * time.Second) // 等待执行
	pool.Close()
	time.Sleep(3 * time.Second) // 等待执行
	fmt.Println(pool.runningWorkers)
	time.Sleep(3 * time.Second) // 等待执行
}
