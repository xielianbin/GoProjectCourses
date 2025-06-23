package main
import (
	"sync"
	"fmt"
	"sync/atomic"
)
type DBPool struct {
	maxOpenConns int
	maxIdleConns int
	connTimeout  int
	idleTimeout  int
}
var initialed uint32
var dbPoolIns *DBPool
var lock=&sync.Mutex{}
// GetDBPool 返回一个DBPool实例，如果已经初始化则返回现有实例
// 如果未初始化，则创建一个新的DBPool实例并进行初始化
// 注意：此函数是线程安全的，使用了互斥锁和原子操作来确保多协程环境下的安全性
func GetDBPool(m1,m2 int) *DBPool {
	if atomic.LoadUint32(&initialed) == 1 {
		fmt.Println("DBPool already initialized")
		return dbPoolIns
	}
	lock.Lock()
	defer lock.Unlock()
	if initialed==0{
		dbPoolIns = &DBPool{
			maxOpenConns: m1,
			maxIdleConns: m2,
			connTimeout:  30,
			idleTimeout:  60,
		}
		fmt.Println("DBPool initialized")
		// 使用原子操作设置initialed为1，表示DBPool已经初始化
		atomic.StoreUint32(&initialed, 1)
	}
	return dbPoolIns
}
var once sync.Once
func GetDBPool2(m1,m2 int) *DBPool {
	once.Do(func() {
		dbPoolIns = &DBPool{
			maxOpenConns: m1,
			maxIdleConns: m2,
			connTimeout:  30,
			idleTimeout:  60,
		}
		fmt.Println("DBPool initialized using sync.Once")
	})
	return dbPoolIns
}

func main() {
	for i:=0;i<10;i++{
		go func() {
			dbPool := GetDBPool2((i+2)*10, (i+2)*5)
			fmt.Println(dbPool)
			// fmt.Printf("DBPool: %p, MaxOpenConns: %d, MaxIdleConns: %d\n", dbPool, dbPool.maxOpenConns, dbPool.maxIdleConns)
		}()
	}
	fmt.Scanln() // 阻塞主线程，等待所有协程完成
	fmt.Println("Main function completed")
}