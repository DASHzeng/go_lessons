package fifth

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	window         int     = 10
	bucketDuration int     = 10
	curCount       int32   = 0
	curFalse       int32   = 0
	totalCount     int32   = 0
	totalFalse     int32   = 0
	ring                   = ItemQueue{}
	limiter        int32   = 100
	errLimiter     float64 = 0.9
	totalLock      sync.Mutex
)

func main() {
	ring.New()
	for i := 0; i < window-1; i++ {
		bucket := [2]int32{0, 0}
		ring.Enqueue(bucket)
	}
	//每隔10秒更新一次
	go func() {
		timer := time.NewTicker(time.Duration(bucketDuration) * time.Second)
		for range timer.C {
			arr := ring.Dequeue()
			oldBucket := (*arr).([]int32)
			totalLock.Lock()
			totalCount = atomic.AddInt32(&totalCount, curCount-oldBucket[0])
			totalFalse = atomic.AddInt32(&totalFalse, curFalse-oldBucket[1])
			totalLock.Unlock()
			ring.Enqueue([2]int32{curCount, curFalse})
			curCount = atomic.AddInt32(&curCount, -curCount)
			curFalse = atomic.AddInt32(&curFalse, -curFalse)
		}
	}()

}
func getRequest() {
	//
	if totalCount > 0 {
		fal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(totalFalse)/float64(totalCount)), 64)
		if totalCount < limiter || fal < errLimiter {
			_, err := doSomeThing()
			if err == nil {
				curFalse = atomic.AddInt32(&curFalse, 1)
			}
		}
	}
	curCount = atomic.AddInt32(&curCount, 1)
}
func doSomeThing() (int, error) {
	return 1, errors.New("xx")

}
