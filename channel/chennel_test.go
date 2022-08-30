package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan []int, 0)
	go func() {
		ch <- []int{1}
		close(ch)
	}()

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	ch = nil
	time.Sleep(time.Second)
	_, ok := <-ch
	fmt.Println(ok)
	fmt.Println(<-ch)

}
