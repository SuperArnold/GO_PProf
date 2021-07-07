# 透過瀏覽器存取
`http://localhost:6060/debug/pprof/`

# 透過互動式終端使用
* 直接透過命令列完成對正在執行的應用程式PProf進行抓取和分析，可以用top 10來得到結果。 <br>
`go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=30`

* 可以快速的得到結果，因為它不像CPU Profiling那樣需要作取樣等待 <br>
`go tool pprof http://localhost:6060/debug/pprof/heap`
  * inuse_space: 分析應用程式常駐記憶體的佔用情況。
  * alloc_object: 分析應用程式的記憶體臨分時配情況。

* 分別對應檢視每個函數的物件數量和分配的記憶體大小，可以用traces指令來印出對應的所有呼叫堆疊以及指標資訊。<br>
`go tool pprof http://localhost:6060/debug/pprof/goroutine`

* 分別對應檢視每個函數的物件數量和分配的記憶體大小，可以用traces指令來印出對應的所有呼叫堆疊以及指標資訊。<br>
