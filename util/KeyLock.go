package util

import (
	"context"
	"crypto/rand"
	"math/big"
	"pro/common"
	"time"
)

const (
	waitTimeOut    = 30 * 1000 //最长等待时间 单位毫秒  等待30s
	lockExpireTime = 10        //锁过期时间 单位秒  10s过期
)

var ctx = context.Background()

// keyLock
func Lock(key, val string) bool {
	println(val + "获取锁")
	var wait = waitTimeOut
	for wait > 0 {
		//获取锁
		success := common.SetNx(key, val, lockExpireTime)
		if success {
			println(val + "获取锁成功")
			return true
		}
		println(val + "休眠等待锁")
		//随机睡眠时间
		random, _ := rand.Int(rand.Reader, big.NewInt(1000))
		sleepTime := int(random.Int64())
		//扣除等待时间
		wait -= sleepTime
		//线程睡眠
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
	return false
}

func Unlock(key, val string) {
	println(val + "释放锁")
	exist := common.Get(key)
	if exist == val {
		println(val + "试释锁成功")
		common.Del(key)
	}
}
