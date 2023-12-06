// Package grtm
// @Description:
// @Author AN 2023-12-06 23:19:56
package grtm

func InitCoPool() {
	// 默认连接池
	// global.CoPool = defaultPool()
}

// func defaultPool() *ants.PoolWithFunc {
//	//var wg sync.WaitGroup
//	//wg.Add(1)
//	////p, _ := ants.NewPoolWithFunc(1, func(i interface{}) {
//	////	queryDbCount(i)
//	////	wg.Done()
//	////})
//	////p, _ := ants.NewPoolWithFunc(1, func(i interface{}) {
//	////	queryEsCount(i)
//	////	wg.Done()
//	////})
//	//return p
// }
