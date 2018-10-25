package main

import (
	"fmt"
	"runtime"
  "time"
  "net/http"

	"github.com/carlescere/scheduler"
)

func main() {
	// tick := func() {
	// 	fmt.Println("Tick! @", time.Now().UTC())
  // }

  access := func() {
    resp, err := http.Get("https://www.baidu.com")
    if err != nil {
      fmt.Println("Access baidu failed", time.Now().UTC())
    } else {
      fmt.Println("Access baidu successfully @ ", time.Now().UTC(), "status: ", resp.Status)
    }
  }

  // Run every 2 seconds but not now.
  scheduler.Every(5).Seconds().NotImmediately().Run(access)

  scheduler.Every().Day().At("8:30").Run(access)

  // Keep the program from not exiting.
	runtime.Goexit()
}
