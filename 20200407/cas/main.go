package main
import (
    "fmt"
    "test/cas/atm"
    "time"
    "sync"
    "sync/atomic"
)

var (
    counter int32          //计数器
    wg      sync.WaitGroup //信号量
)

func main() {

    threadNum := 5

    //1. 五个信号量
    wg.Add(threadNum)

    //2.开启5个线程
    for i := 0; i < threadNum; i++ {
        go incCounter(i)
    }

    atm.SwapInt32(&threadNum, 4)
    //3.等待子线程结束
    wg.Wait()
    fmt.Println(counter)
}

func incCounter(index int) {
    defer wg.Done()

    spinNum := 0
    looptime := 0
    for {
        //2.1原子操作
        old := counter
	looptime++
       
    	fmt.Printf("thread,%d,looptime,%d\n",index,looptime)
	time.Sleep(time.Millisecond)
       
	ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
        if ok {
            break
        } else {
            spinNum++
        }
    }

    fmt.Printf("thread,%d,spinnum,%d\n",index,spinNum)

}
