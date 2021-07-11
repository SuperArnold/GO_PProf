package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

var datas []string

type LookupType int8

const (
	LookupGoruotine LookupType = iota
	LookupThreadcreate
	LookupHeap
	LookupAllocs
	LookupBlock
	LookupMutex
)

func init() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
}
func main() {
	http.HandleFunc("/lookup/heap", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupHeap, os.Stdout)
	})
	http.HandleFunc("/lookup/threadcreate", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupThreadcreate, os.Stdout)
	})
	http.HandleFunc("/lookup/block", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupBlock, os.Stdout)
	})
	http.HandleFunc("/lookup/goroutine", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupGoruotine, os.Stdout)
	})
	_ = http.ListenAndServe("0.0.0.0:6060", nil)

}

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

func pprofLookup(lookupType LookupType, w io.Writer) error {
	var err error
	switch lookupType {
	case LookupGoruotine:
		p := pprof.Lookup("goroutine")
		err = p.WriteTo(w, 2)
	case LookupThreadcreate:
		p := pprof.Lookup("threadcreate")
		err = p.WriteTo(w, 2)
	case LookupHeap:
		p := pprof.Lookup("heap")
		err = p.WriteTo(w, 2)
	case LookupAllocs:
		p := pprof.Lookup("allocs")
		err = p.WriteTo(w, 2)
	case LookupBlock:
		p := pprof.Lookup("block")
		err = p.WriteTo(w, 2)
	case LookupMutex:
		p := pprof.Lookup("mutex")
		err = p.WriteTo(w, 2)
	}

	return err
}
