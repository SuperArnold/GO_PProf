package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("GO PProf demo"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)

}

// func init() {
// 	runtime.SetBlockProfileRate(1)
// 	runtime.SetMutexProfileFraction(1)
// }
// func main() {
// 	var m sync.Mutex
// 	var datas = make(map[int]struct{})
// 	for i := 0; i < 999; i++ {
// 		go func(i int) {
// 			m.Lock()
// 			defer m.Unlock()
// 			datas[i] = struct{}{}
// 		}(i)
// 	}
// 	_ = http.ListenAndServe("0.0.0.0:6061", nil)

// }
