package default_router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func validateIfKey1False(value map[string]bool) bool {
	time.Sleep(1 * time.Second)
	key, ok := value["1"]
	if !ok {
		return false
	}
	return key
}

func validateIfKey1FalseChannel(value map[string]bool, channel chan bool, exit chan bool) {
	time.Sleep(1 * time.Second)
	key, ok := value["1"]
	if !ok {
		exit <- false
	}
	if !key {
		exit <- false
	}

	channel <- true
}

func validateIfKey2False(value map[string]bool) bool {
	time.Sleep(1 * time.Second)
	key, ok := value["2"]
	if !ok {
		return false
	}
	return key
}

func validateIfKey2FalseChannel(value map[string]bool, channel chan bool, exit chan bool) {
	time.Sleep(1 * time.Second)
	key, ok := value["2"]
	if !ok {
		exit <- false
	}
	if !key {
		exit <- false
	}
	channel <- true
}

func validateIfKey3False(value map[string]bool) bool {
	time.Sleep(1 * time.Second)
	key, ok := value["3"]
	if !ok {
		return false
	}
	return key
}

func validateIfKey3FalseChannel(value map[string]bool, channel chan bool, exit chan bool) {
	time.Sleep(1 * time.Second)
	key, ok := value["3"]
	if !ok {
		exit <- false
	}
	if !key {
		exit <- false
	}
	channel <- true
}

func validateIfKey4False(value map[string]bool) bool {
	time.Sleep(1 * time.Second)
	key, ok := value["4"]
	if !ok {
		return false
	}
	return key
}

func validateIfKey4FalseChannel(value map[string]bool, channel chan bool, exit chan bool) {
	time.Sleep(4 * time.Second)
	key, ok := value["4"]
	if !ok {
		exit <- false
	}
	if !key {
		exit <- false
	}
	channel <- true
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		goRoutine := c.Request.URL.Query().Get("goroutine") == "true"
		if goRoutine {
			result := true
			var channel = make(chan bool)
			var exit = make(chan bool)
			var wg sync.WaitGroup
			wg.Add(1)
			go Goroutine(c, channel, exit, &result, &wg)
			wg.Wait()

			if result {
				c.JSON(http.StatusAccepted, gin.H{"message": "success"})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}

		} else {
			nonGoroutine(c)
			return
		}

	}
}

func Goroutine(c *gin.Context, channel chan bool, exit chan bool, result *bool, wg *sync.WaitGroup) {
	var hashMap map[string]bool
	json.NewDecoder(c.Request.Body).Decode(&hashMap)

	go validateIfKey1FalseChannel(hashMap, channel, exit)
	go validateIfKey2FalseChannel(hashMap, channel, exit)
	go validateIfKey3FalseChannel(hashMap, channel, exit)
	go validateIfKey4FalseChannel(hashMap, channel, exit)
	count := 0

	go func(channel chan bool, exit chan bool, result *bool, wg *sync.WaitGroup) {

		for {
			select {
			case success, _ := <-channel:

				if success {
					count += 1
				}
				if count == 4 {
					close(exit)
				}

			case _, status := <-exit:
				fmt.Println(status)
				if count != 4 {
					*result = false
				}
				wg.Done()
				return

			}
		}

	}(channel, exit, result, wg)

}

func nonGoroutine(c *gin.Context) {

	var hashMap map[string]bool
	json.NewDecoder(c.Request.Body).Decode(&hashMap)
	if validateIfKey1False(hashMap) && validateIfKey2False(hashMap) && validateIfKey3False(hashMap) && validateIfKey4False(hashMap) {
		c.JSON(http.StatusAccepted, hashMap)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
}
