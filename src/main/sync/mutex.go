package main

import "sync"

func main() {
	// todo:
}

var mutex sync.Mutex
var rwMutex sync.RWMutex

// Mutex 锁
func Mutex() {
	mutex.Lock()
	defer mutex.Unlock()

	// 代码
}

// RwMutex 读写锁
func RwMutex() {
	// 读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()
}

// Failed1 不可重入
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 如果只有一个 goroutine 会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

// Failed2 不可升级
func Failed2() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 读锁之后不可使用写锁
	// 如果只有一个 goroutine 会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}
