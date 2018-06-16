
# goroutines

- https://blog.csdn.net/skh2015java/article/details/60330785

# channel/buffered channels

- https://blog.csdn.net/skh2015java/article/details/60330785

# select

- https://blog.csdn.net/skh2015java/article/details/60330975

- https://blog.csdn.net/skh2015java/article/details/74818854

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md

# Range/Close

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md
  
# time.After

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md

# runtime包

Goexit/Gosched/NumCPU/NumGoroutine/GOMAXPROCS

- https://www.jianshu.com/p/e45cea3e1723

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md

- https://blog.csdn.net/skh2015java/article/details/60330875

# Mutex/RWMutex

Mutex/RWMutex(讀寫鎖, 用在讀多於寫的情境)

- https://blog.csdn.net/skh2015java/article/details/60334437

# sync.Once

- http://shanks.leanote.com/post/Untitled-55ca439338f41148cd000759-12

<!-- 
TODO

https://golang.org/pkg/sync/
把裡面還沒用過的在詳細確認

多個 goroutine時如何讓某些 goroutine優先執行, 例如用來收數據的 goroutine就需要優先權較高

https://gocn.vip/question/1378
https://www.cnblogs.com/dearplain/p/8276138.html

goroutine和 內部怎處理的?為什麼不能稱之為Thread?
https://github.com/k2huang/blogpost/blob/master/golang/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6/Go%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6.md

func sync.NewCond(l lock) * Cond
Wait
Signal
Broadcast

sync/atomic包如何運用

sync.WaitGroup

sync.Pool

sync/Map 併發的集合?

-->