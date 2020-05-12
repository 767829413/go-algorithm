package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type demo struct {
	name string
	age  int
}

const (
	NUM1 = iota << 5
	NUM2
	NUM3
)

func main() {
	fmt.Println(NUM1)
	fmt.Println(NUM2)
	fmt.Println(NUM3)
}

func ParseInputDidString(did string) (string, []string) {
	data := make([]string, 0)
	if strings.Contains(did, "#") {
		data = strings.Split(did, "#")
	}
	if len(data) > 1 {
		data = data[1:]
	}
	return did, data
}

func test1() {
	t := time.NewTicker(100 * time.Millisecond)
	ctx := context.Background()
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			_, cancel := context.WithCancel(ctx)
			defer wg.Done()
			defer cancel()
			select {
			case <-t.C:
				fmt.Println("ok" + "-" + strconv.Itoa(i))
			}
		}(i, ctx)
	}
	wg.Wait()
}

func test2() {
	data := make(map[interface{}]interface{})
	k1 := "sdsds"
	k2 := demo{"sdsd", 45,}
	k3 := &demo{"wqeqwr", 36,}
	data[k1] = 1111
	data[k2] = 2222
	data[k3] = 3333

	fmt.Println(data[k1])
	fmt.Println(data[k2])
	fmt.Println(data[k3])
	fmt.Println(data[demo{"sdsd", 45,}])
	fmt.Println(data[&demo{"wqeqwr", 36,}])
}
