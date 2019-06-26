package datasource

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"lottery/conf"
	"sync"
	"time"
)

var rdsLock sync.Mutex
var cacheInstance *RedisConn

type RedisConn struct {
	pool *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string, args...interface{}) (reply interface{}, err error){
	conn := rds.pool.Get()
	defer conn.Close()

	t1 := time.Now().UnixNano()
	reply, err = conn.Do(commandName,args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println(e,err)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("[redis] [info] [%dus]cmd=%s, err=%s, arge=%v, reply=&s\n",(t2-t1)/1000,commandName,err,args,reply)
	}
	return reply,err

}

func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()

	if cacheInstance != nil {
		return cacheInstance
	}
	return NewCache()
}

func NewCache() *RedisConn {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RdsCache.Host, conf.RdsCache.Port))
			if err != nil {
				log.Fatal("rdshelper.newcache  dial error ",err)
				return nil,err
			}
			return c,nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:1000,
		MaxActive:1000,
		IdleTimeout:0,
		Wait:false,
		MaxConnLifetime:0,

	}
	instance := &RedisConn{
		pool:&pool,
	}
	cacheInstance = instance
	instance.ShowDebug(true)
	return instance
}


















