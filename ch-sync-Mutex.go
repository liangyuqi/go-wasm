// +build ignore

/*
Go 程序可以使用通道进行多个 goroutine 间的数据交换，但这仅仅是数据同步中的一种方法。
通道内部的实现依然使用了各种锁，因此优雅代码的代价是性能。
在某些轻量级的场合，原子访问（atomic包）、互斥锁（sync.Mutex）以及等待组（sync.WaitGroup）能最大程度满足需求。
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 goroutine 可以访问共享资源。

type Mutex
    func (m *Mutex) Lock()
    func (m *Mutex) Unlock()
其中Mutex为互斥锁，Lock()加锁，Unlock()解锁，
使用Lock()加锁后，便不能再次对其进行加锁，直到利用Unlock()解锁对其解锁后，才能再次加锁．
适用于读写不确定场景，即读写次数没有明显的区别，并且只允许只有一个读或者写的场景，所以该锁也叫做全局锁．
func (m *Mutex) Unlock()用于解锁，如果在使用Unlock()前未加锁，就会引起一个运行错误．
已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁．


type RWMutex
    func (rw *RWMutex) Lock()
    func (rw *RWMutex) RLock()
    func (rw *RWMutex) RLocker() Locker
    func (rw *RWMutex) RUnlock()
    func (rw *RWMutex) Unlock()

RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景．

  func (rw *RWMutex) Lock()　　写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
  func (rw *RWMutex) Unlock()　写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误．
func (rw *RWMutex) RLock() 读锁，当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景
func (rw *RWMutex)RUnlock()　读锁解锁，RUnlock 撤销单次 RLock 调用，它对于其它同时存在的读取器则没有效果。若 rw 并没有为读取而锁定，调用 RUnlock 就会引发一个运行时错误(注：这种说法在go1.3版本中是不对的，运行过程中允许RUnLock早于RLock一个，也只能早于１个)。

*/

package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex

	// 逻辑中使用的某个变量
	count2 int
	// 与变量对应的使用互斥锁
	countGuard2 sync.RWMutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func SetCount2(c int) {
	// countGuard2.WLock()
	// count2 = c
	// countGuard2.WUnlock()
}

func main() {
	// 可以进行并发安全的设置
	SetCount(1)

	// SetCount2(2)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}
