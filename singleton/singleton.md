#  单例模式

## 概述

单例，是指一个对象至多只有一个实例，且单例的实现不应该由人来控制，而应该由代码来强制单例

##  应用场景

1.  网站的计数器，若多例计数则难以同步
2.  Windows 的任务管理器以及回收站
3.  应用程序的日志应用，一般都用单例模式实现，多例内容不好追加
4.  多线程的线程池，单例可以方便对线程进行统一控制
5.  数据库的连接池，使用单例模式能节省打开或关闭数据库连接所引起的效率损耗

### 总结：

优势：可以避免对象频繁创建带来的性能损耗

## 实现(仅列出关键代码)

### 懒汉式

```go
func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
```

懒汉式缺点在于，在并发条件下，恰好第一个 goroutine 执行到 m = &singleton{} 这句话之前，第二个 goroutine 也来获取实例，第二个 goroutine 去判断 m 是不是 nil,因为 m = &singleton{}还没有来得及执行，所以 m 还是 nil，所以会导致出现两个实例

### 带线程锁

```go
func GetInstance() *singleton {
    lock.Lock()
    defer lock.Unlock()
    if m == nil {
        m = &singleton {}
    }
    return m
}
```

带线程锁就是在创建实例时加锁， 解决并发时出现的竞争问题， 缺点在于每一次获取该实例时都会进行加锁， 在频繁调用时降低了效率

### 双重锁

```go
func GetInstance() *singleton {
    if m == nil {
        lock.Lock()
        defer lock.Unlock()
        if m == nil {
            m = &singleton {}
        }
    }
    return m
}
```

 这样就可以在并发条件下  高效的保证单例了。

### once

然而在 go 的 sync 包中有一个 once 方法，可以保证函数只执行一次，而且是线程安全的，所以  在 go 中单例模式的正确打开方式为

```go
func getInstace() *singleton {
	once.Do(func() {
		instance = &singleton{""}
	})
	return instance
}
```
