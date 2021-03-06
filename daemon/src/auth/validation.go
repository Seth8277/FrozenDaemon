/*
本包提供了一些关于用户的ValidationKey的一些函数
用户在直连daemon时，使用这个key进行鉴权
鉴权Key由Panel进行指定和分发
*/
package auth

import "time"

// 全局Map
// 该Map指定了某key -> Key对
var keys = make(map[string]*KeyPair, 0)

const (
	KEY_VERIFY_FAILED = -1
)

func Timer() {
	for {
		for k, v := range keys {
			if v.Time <= time.Now().Unix() {
				delete(keys, k)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

// 传入Key,返回服务器id或一些状态值
func VerifyKey(key string) int {
	if pair, ok := keys[key]; ok {
		return pair.ID
	}
	return KEY_VERIFY_FAILED
}

func KeyRegister(key string, id int) {
	keys[key] = &KeyPair{
		ID:   id,
		Time: time.Now().Unix() + 600,
	}
}
