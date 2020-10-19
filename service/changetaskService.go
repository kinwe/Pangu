package service

import (
	"fmt"
	"sync"
	"test/utils"
)

func Changetask(url string,wg *sync.WaitGroup) {
	fmt.Println(utils.Get(url))
	wg.Done()
}
