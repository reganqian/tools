package redis

import (
	// "errors"
	// "strconv"
	// "time"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/labstack/gommon/log"
	// log "github.com/sirupsen/logrus"
)

var (
	RDSTest RedisDataStore
)

func init() {
	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
		log.Error(err1.Error())
	}

	// 2. 通过对象获取数据
	redisAddr := iniconf.String("redis::redis_addr")
	redisPwd := iniconf.String("redis::redis_pwd")
	redisDb := iniconf.String("redis::redis_db")

	// redis test 库初始化
	RDSTest = RedisDataStore{
		RedisHost: redisAddr,
		RedisPwd:  redisPwd,
		RedisDB:   redisDb,
		Timeout:   20,
		RedisPool: nil,
	}
	RDSTest.RedisPool = RDSTest.NewPool()
}

func RedisGet(t string) (v interface{}, err error) {
	v, err = RDSTest.Get(t)
	if err != nil {
		fmt.Println(v)
	}
	return v, err
}

func RedisSetEx(k, v string, timeLong int64) {
	err := RDSTest.SetEx(k, v, timeLong)
	if err != nil {
		fmt.Println("redis保存失败", err)
	} else {
		fmt.Println("redis保存成功", v)
	}

}

/**
* 默认保存时间为12小时 60 * 60 * 12
* //30天
**/
func RedisSetExDefault(k, v string) {
	err := RDSTest.SetEx(k, v, 60*60*24*30)
	if err != nil {
		fmt.Println("redis保存失败", err)
	} else {
		fmt.Println("redis保存成功", v)
	}

}

func RedisSet(k, v string) {
	err := RDSTest.Set(k, v)
	if err != nil {
		fmt.Println(err)
	}

}

func RedisDel(k string) bool {
	err := RDSTest.Del(k)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}
