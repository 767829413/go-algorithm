package concurrency

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/767829413/go-algorithm/utilcom/randomstring"
)

func TestChannelUse(t *testing.T) {
	num := 10000000
	strs := make([]string, num)
	for i := 0; i < num; i++ {
		strs[i] = randomstring.RandStringBytesMaskImprSrc(3)
	}
	start := time.Now()
	toOp(strs)
	t.Log(time.Since(start))
}

func toOp(strs []string) {
	errs := make(chan error, len(strs))
	for _, v := range strs {
		go func(string) {
			_, err := op(v)
			errs <- err
		}(v)
	}
	var total = 0
	for range strs {
		if err := <-errs; err != nil {
			total++
			log.Println(err)
		}
	}
	log.Println(total)
}

func op(str string) (string, error) {
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	if result.Int64() > 50 {
		return "", errors.New("throw a error")
	}
	return strings.ToUpper(str), nil
}
