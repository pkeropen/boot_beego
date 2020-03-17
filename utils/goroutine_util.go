package utils


import (
	"fmt"
	"runtime"
	"sync"

)

//OperateType 操作类型
type OperateType int

const (
	//AddOperate 加
	AddOperate OperateType = 1 + iota

	//ReduceOperate 减
	ReduceOperate
)


var (
	//goroutineMap go协程数量记录表
	goroutineMap = make(map[string]int, 1024)

	//协程锁
	mutex sync.Mutex
)

//Operate 操作map
func Operate(goroutineMapName string, operateType OperateType) {
	mutex.Lock()
	defer mutex.Unlock()

	count, ok := goroutineMap[goroutineMapName]
	if !ok {
		count = 0
	}

	//操作后的数量
	newCount := 0

	if operateType == AddOperate {
		newCount = count + 1
	}

	if operateType == ReduceOperate {
		if count > 0 {
			newCount = count - 1
		} else {
			newCount = 0
		}
	}

	goroutineMap[goroutineMapName] = newCount
}

//ToString 获取map的描述
func ToString() string {
	mutex.Lock()
	defer mutex.Unlock()

	//获取描述
	str := fmt.Sprintf("Goroutine Info:(%s,%d)", "Goroutine总数", runtime.NumGoroutine())
	for key, value := range goroutineMap {
		str += fmt.Sprintf("(%s,%d)", key, value)
	}

	return str
}
