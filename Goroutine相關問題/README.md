
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

- Goexit 退出当前执行的goroutine，但是defer函数还会继续调用

- Gosched 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

- NumCPU 返回 CPU 核数量

- NumGoroutine 返回正在执行和排队的任务总数

- GOMAXPROCS 用来设置可以并行计算的CPU核数的最大值，并返回之前的值

- https://www.jianshu.com/p/e45cea3e1723

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md

- https://blog.csdn.net/skh2015java/article/details/60330875

# Mutex/RWMutex

- https://blog.csdn.net/skh2015java/article/details/60334437

Mutex

func (m *Mutex) Lock()

func (m *Mutex) Unlock()

RWMutex(讀寫鎖, 用在讀多於寫的情境)

func (rw *RWMutex) Lock()

func (rw *RWMutex) RLock()

func (rw *RWMutex) RLocker() Locker

func (rw *RWMutex) RUnlock()

func (rw *RWMutex) Unlock()