# 透過瀏覽器存取
`http://localhost:6060/debug/pprof/`

# 透過互動式終端使用
* 直接透過命令列完成對正在執行的應用程式PProf進行抓取和分析，可以用`top 10`來得到結果。 <br>
`go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=30`

* 可以快速的得到結果，因為它不像CPU Profiling那樣需要作取樣等待 <br>
`go tool pprof http://localhost:6060/debug/pprof/heap`
  * inuse_space: 分析應用程式常駐記憶體的佔用情況。
  * alloc_object: 分析應用程式的記憶體臨分時配情況。

* 分別對應檢視每個函數的物件數量和分配的記憶體大小，可以用`traces`指令來印出對應的所有呼叫堆疊以及指標資訊。<br>
`go tool pprof http://localhost:6060/debug/pprof/goroutine`

* Mutex Profiling發生互斥鎖時，可以透過下面方法找出互斥量的分析。<br>
需設定runtime.SetMutexProfileFraction(1)<br>
`go tool pprof http://localhost:6061/debug/pprof/mutex`
  * `top`: 檢視互斥量的排名。
  * `list main`: 檢視指定函數的程式情況。

* block Profiling發生互斥鎖時，可以透過下面方法找出互斥量的分析。<br>
需設定runtime.SetBlockProfileRate(1)<br>
`go tool pprof http://localhost:6061/debug/pprof/block`
  * `top`: 檢視互斥量的排名。
  * `list main`: 檢視指定函數的程式情況。

# 檢視視覺化介面
`wget http://localhost:6060/debug/pprof/profile`
執行完畢後會在目錄下產生profile檔案，針對視覺化介面可使用<br>
`go tool pprof profile`
 使用`web`指令
  * 如發生failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH，則需要安裝Graphviz
  `brew install graphviz`
  * 可以查看產出的`pprof001.svg`Graph的整體函數呼叫流程
    * 框越大、線越粗、框顏色越鮮豔，則代表佔用的時間越久。


  
