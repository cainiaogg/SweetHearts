package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/cache"
)

type MyRedisClient struct {
	RedisClientList  []cache.Cache
	masterClientList *cache.Cache
}

func NewMyRedisClient() *MyRedisClient {
	myRedisClient := MyRedisClient{}
	for _, addr := range REDIS_LIST {
		redisClient, err := cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s", "key":"beecacheRedis"}`, addr))
		if err != nil {
			fmt.Println("Error ", err)
		}
		myRedisClient.RedisClientList = append(myRedisClient.RedisClientList, redisClient)
	}
	return &myRedisClient
}

func (this *MyRedisClient) Get(key string) interface{} {
	var content []uint8
	for _, addr := range this.RedisClientList {
		tmp_content := addr.Get(key)
		if tmp_content == nil {
			continue
		}
		content = addr.Get(key).([]uint8)
		if content != nil {
			break
		}
	}
	if content == nil {
		return nil
	}
	var result string = ""
	for _, val := range content {
		result = result + string(rune(val))
	}
	if result[0] == 91 && result[len(result)-1] == 93 {
		result = result[1 : len(result)-1]
		result = strings.Trim(result, " ")
		result_list := strings.Split(result, " ")
		fmt.Printf("Info get key:%v from redis success\n", result_list)
		return result_list
	} else {
		fmt.Printf("Info get key:%v from redis success\n", result)
		return result
	}
}

func (this *MyRedisClient) Put(key string, val interface{}, timeout time.Duration) error {
	M_REDIS.Lock()
	err := this.PutNext(key, val, timeout)
	M_REDIS.Unlock()
	return err
}

func (this *MyRedisClient) PutNext(key string, val interface{}, timeout time.Duration) error {
	flag := 0
	if this.masterClientList != nil {
		err := (*this.masterClientList).Put(key, val, timeout)
		if err == nil {
			flag = 1
		}
		if err != nil {
			fmt.Println("info ", err)
			for _, addr := range this.RedisClientList {
				err := addr.Put(key, val, timeout)
				if err == nil {
					this.masterClientList = &addr
					flag = 1
					break
				}
				fmt.Println("info ", err)
			}
		}
	} else {
		for _, addr := range this.RedisClientList {
			err := addr.Put(key, val, timeout)
			if err == nil {
				this.masterClientList = &addr
				flag = 1
				break
			}
			fmt.Println("info", err)
		}
	}

	if flag == 0 {
		var err error
		err = fmt.Errorf("error put key into redis")
		return err
	}
	fmt.Printf("info put %s %s into redis success\n", key, val)
	return nil
}
