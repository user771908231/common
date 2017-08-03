package data

import (
	"casino_common/common/cfg"
	"casino_common/common/log"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"time"
)

var (
	Redis_svr   string
	Redis_name  string
	Redis_pwd   string
	RedisClient *redis.Pool
)

func initPool() {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Redis_svr, redis.DialPassword(Redis_pwd))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", Redis_name)
			return c, nil
		},
	}
}

func InitRedis(redisAddr, redisName string, redisPwd string) {
	config := cfg.Get()

	Redis_svr = redisAddr
	Redis_name = redisName
	Redis_pwd = redisPwd
	if len(config["redis_svr"]) != 0 {
		Redis_svr = config["redis_svr"]
	}

	log.T("redis_svr:%v", Redis_svr)

	//初始化连接池
	initPool()
}

type Data struct {
	conn redis.Conn
}

func (t *Data) Open() error {
	var err error
	//t.conn, err = redis.DialTimeout("tcp", Redis_svr, 3*time.Second, 10*time.Second, 10*time.Second) //timeout=10s
	t.conn = RedisClient.Get()
	if err != nil {
		log.Error("ERR: redis Open(table:%v) ret err:%v t.conn:%v", Redis_name, err, t.conn)
		return err
	}
	return err
}

func (t *Data) Close() (err error) {
	if t.conn != nil {
		//log.T("[TRACE] t.conn.Close().")
		err = t.conn.Close()
		t.conn = nil
		return err
	} else {
		log.Fatal("FATAL ERR: try redis.Close() BUT t.conn=nil")
	}
	return err
}

func (t *Data) Select(table string) (err error) {
	//	_, err = t.conn.Do("SELECT", table)
	//	if err != nil {
	//		log.Error("[ERROR] redis.Select(%v) ret err:%v", table, err)
	//		return err
	//	}
	//	log.T("db.select(%v) ok.", table)
	//
	//	return err
	return nil
}

//================= String ==================

func (t *Data) GetKeys(keyword string) (keys []string, err error) {
	if t.conn != nil {
		if keyword == "" {
			keyword = "*"
		}

		keys, err = redis.Strings(t.conn.Do("KEYS", keyword))
		if err != nil {
			if err == redis.ErrNil {
				err = nil
			}else {
				log.E("[Redis] GetKeys fail %s err:%v", keyword, err)
			}
		}
		return keys, err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}

	return keys, err
}

func (t *Data) Get(key string) (value string, err error) {
	if t.conn != nil {
		value, err := redis.String(t.conn.Do("GET", key))
		//log.T("redis.GET(%v) ret err:%v value:%v", key, err, value)
		if err != nil {
			if err == redis.ErrNil {
				err = nil
			}else {
				log.E("[Redis] Get fail %s err:%v", key, err)
			}
		}


		return value, err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}

	return "", err
}

//
func (t *Data) MGet(args []interface{}) (values []interface{}, err error) {
	if t.conn != nil {
		values, err := redis.Values(t.conn.Do("MGET", args...))
		if err != nil {
			if err == redis.ErrNil {
				err = nil
			}else {
				log.E("[Redis] MGet fail %v err:%v", args, err)
			}
		}

		//log.T("redis.GET(%v) ret err:%v value:%v", args, err, values)
		return values, err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}

	return values, err
}

//return []byte
func (t *Data) Gets(key string) (value []byte, err error) {
	if t.conn != nil {
		value, err := redis.Bytes(t.conn.Do("GET", key))
		if err != nil {
			if err == redis.ErrNil {
				err = nil
			}else {
				log.E("[Redis] Get fail %s err:%v", key, err)
			}
		}

		return value, err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}

	return nil, err
}

func (t *Data) GetInt64(key string) (value int64, err error) {
	if t.conn != nil {
		value, err := redis.Int64(t.conn.Do("GET", key))
		if err != nil {
			if err == redis.ErrNil {
				err = nil
			}else {
				log.E("[Redis] GetInt64 fail %s err:%v", key, err)
			}
		}

		return value, err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}

	return 0, err
}

func (t *Data) Set(key string, value []byte) error {
	if t.conn != nil {
		//log.T("[TRACE] try redis.Set(%v) value:%v", key, value)
		_, err := redis.String(t.conn.Do("SET", key, value))
		//log.T("[TRACE] after redis.Set(%v) ret err:%v", key, err)
		if err != nil {
			log.E("[Redis] Set fail %s err:%v", key, err)
		}
		return err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}
	return nil
}

func (t *Data) SetInt(key string, value int64) error {
	if t.conn != nil {
		_, err := redis.String(t.conn.Do("SET", key, value))
		if err != nil {
			log.E("[Redis] SetInt fail %s err:%v", key, err)
		}
		return err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}
	return nil
}

func (t *Data) SetUInt(key string, value uint32) error {
	if t.conn != nil {
		_, err := redis.String(t.conn.Do("SET", key, value))
		if err != nil {
			log.E("[Redis] SetUint fail %s err:%v", key, err)
		}
		return err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}
	return nil
}

func (t *Data) SetObj(key string, pb proto.Message) error {
	d, err := proto.Marshal(pb)
	if err != nil {
		log.E("[Redis] SetObj fail %s err:%v", key, err)
		return err
	}
	return t.Set(key, d)
}

func (t *Data) GetObj(key string, pb proto.Message) error {
	d, err := t.Gets(key)
	if err != nil {
		log.E("[Redis] SetObj fail %s err:%v", key, err)
		return err
	}
	return proto.Unmarshal(d, pb)
}

func (t *Data) GetObjv2(key string, pb proto.Message) proto.Message {
	d, err := t.Gets(key)
	if err != nil {
		log.E("[Redis] SetObjV2 fail %s err:%v", key, err)
		return nil
	}
	if d == nil {
		return nil
	} else {
		proto.Unmarshal(d, pb)
		return pb
	}
}

func (t *Data) Del(key string) error {
	if t.conn != nil {
		_, err := t.conn.Do("DEL", key)
		if err != nil {
			log.E("[Redis] Del fail %s err:%v", key, err)
		}
		return err
	} else {
		log.Fatal("invalid redis conn:%v", t.conn)
	}
	return nil
}

//================= List ==================
func (t *Data) ListGetAll(key string) (values []interface{}, err error) {

	values, err = redis.Values(t.conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] ListGetAll fail %s err:%v", key, err)
		}
	}


	return values, err
}

func (t *Data) LIndex(key string, index int) (value []byte, err error) {

	value, err = redis.Bytes(t.conn.Do("LINDEX", key, index))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] LIndex fail %s err:%v", key, err)
		}
	}

	return value, err
}

func (t *Data) LRange(key string, start int, stop int) (values []interface{}, err error) {

	values, err = redis.Values(t.conn.Do("LRANGE", key, start, stop))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] LRange fail %s err:%v", key, err)
		}
	}

	return values, err
}

func (t *Data) LLen(key string) (count int, err error) {

	count, err = redis.Int(t.conn.Do("LLEN", key))
	if err != nil {
		if err == redis.ErrNil {
			count = 0
			err = nil
		}else {
			log.E("[Redis] LLen fail %s err:%v", key, err)
		}
	}

	return count, err
}

func (t *Data) LPush(key string, value []byte) (err error) {
	_, err = t.conn.Do("LPUSH", key, value)
	if err != nil {
		log.E("[Redis] LPush fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) LPushString(key string, value string) (err error) {
	_, err = t.conn.Do("LPUSH", key, value)
	if err != nil {
		log.E("[Redis] LPushString fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) LPushInt(key string, value int32) (err error) {
	_, err = t.conn.Do("LPUSH", key, value)
	if err != nil {
		log.E("[Redis] LPushInt fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) LRem(key string, removeValue []byte) (err error) {
	_, err = redis.Values(t.conn.Do("LREM", 0, removeValue))
	if err != nil {
		log.E("[Redis] LRem fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) LREM(key string, removeValue string) (err error) {
	_, err = redis.Values(t.conn.Do("LREM", 0, removeValue))
	if err != nil {
		log.E("[Redis] LREM fail %s err:%v", key, err)
	}
	return err
}

//================= HASH ==================
func (t *Data) HGetAll(key string) (values []interface{}, err error) {
	values, err = redis.Values(t.conn.Do("HGETALL", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] HGetAll fail %s err:%v", key, err)
		}
	}

	return values, err
}

func (t *Data) HGet(key string, field string) (value []byte, err error) {
	value, err = redis.Bytes(t.conn.Do("HGET", key, field))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] HGet fail %s err:%v", key, err)
		}
	}
	return
}

func (t *Data) HLen(key string) (len int, err error) {
	len, err = redis.Int(t.conn.Do("HLEN", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}else {
			log.E("[Redis] HLen fail %s err:%v", key, err)
		}
	}
	return len, err
}

func (t *Data) HSet(key string, field string, value []byte) (err error) {
	_, err = t.conn.Do("HSET", key, field, value)
	if err != nil {
		log.E("[Redis] HSet fail %s err:%v", key, err)
	}
	return
}

func (t *Data) HMSet(key string, fields ...[]byte) (err error) {
	_, err = t.conn.Do("HMSET", key, fields)
	if err != nil {
		log.E("[Redis] HMSet fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) HDel(key string, field string) (num int, err error) {
	num, err = redis.Int(t.conn.Do("HDEL", key, field))
	if err != nil {
		log.E("[Redis] HDel fail %s err:%v", key, err)
	}
	return num, err
}

//func (t *Data) HDel(key string, field ...interface {}) (num int, err error) {
//TODO: HMDel still has problems
func (t *Data) HMDel(key string, field []string) (num int, err error) {
	num, err = redis.Int(t.conn.Do("HDEL", key, field))
	if err != nil {
		log.E("[Redis] HMDel fail %s err:%v", key, err)
	}
	return num, err
}

//================= ZSET ==================
func (t *Data) ZAdd(key string, member string, score int64) (err error) {
	_, err = t.conn.Do("ZADD", key, score, member)
	if err != nil {
		log.E("[Redis] ZAdd fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) ZRem(key string, member string) (err error) {
	_, err = t.conn.Do("ZREM", key, member)
	if err != nil {
		log.E("[Redis] ZRem fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) ZRangeByScore(key string, min int, max int, offset int, count int) (values []interface{}, err error) {
	//log.T("[TRACE] ZRangeByScore key:%v %v %v limit %v %v", key, min, max, offset, count)
	values, err = redis.Values(t.conn.Do("ZRANGEBYSCORE", key, min, max, "limit", offset, count))
	if err != nil {
		log.E("[Redis] ZRangeByScore fail %s err:%v", key, err)
	}
	return values, err
}

func (t *Data) ZCount(key string, min int, max int) (count int, err error) {
	count, err = redis.Int(t.conn.Do("ZCOUNT", key, min, max))
	if err != nil {
		log.E("[Redis] ZCount fail %s err:%v", key, err)
	}
	return count, err
}

func (t *Data) ZRemRangeByRank(key string, start, stop int) (err error) {
	_, err = t.conn.Do("ZREMRANGEBYRANK", key, start, stop)
	if err != nil {
		log.E("[Redis] ZRemRangeByRank fail %s err:%v", key, err)
	}
	return err
}

func (t *Data) ZREVRANK(key string, mem string) (interface{}, error) {
	values, err := t.conn.Do("ZREVRANK", key, mem)
	if err != nil {
		log.E("[Redis] ZREVRANK fail %s err:%v", key, err)
	}
	return values, err
}

//=================原子加减

func (t *Data) INCRBY(key string, i int64) (interface{}, error) {
	v, err := t.conn.Do("INCRBY", key, i)
	if err != nil {
		log.E("[Redis] INCRBY fail %s err:%v", key, err)
	}
	return v, err
}

func (t *Data) DECRBY(key string, i int64) (interface{}, error) {
	v, err := t.conn.Do("DECRBY", key, i)
	if err != nil {
		log.E("[Redis] DECRBY fail %s err:%v", key, err)
	}
	return v, err
}

func (t *Data) SETNX(k string, i int64) (interface{}, error) {
	v, err := t.conn.Do("DECRBY", k, i)
	if err != nil {
		log.E("[Redis] SETNX fail %s err:%v", k, err)
	}
	return v, err
}

//==================操作集合
func (t *Data) SADD(k string, kv string) (interface{}, error) {
	v, err := t.conn.Do("SADD", k, kv)
	if err != nil {
		log.E("[Redis] SADD fail %s err:%v", k, err)
	}
	return v, err
}

func (t *Data) SREM(k string, kv string) (interface{}, error) {
	v, err := t.conn.Do("SREM", k, kv)
	if err != nil {
		log.E("[Redis] SREM fail %s err:%v", k, err)
	}
	return v, err
}

func (t *Data) SCARD(k string) (int64, error) {
	v, err := t.conn.Do("SCARD", k)
	if err != nil {
		log.E("[Redis] SCARD fail %s err:%v", k, err)
	}
	if err == nil {
		return v.(int64), err
	}
	return 0, nil
}

func (d *Data) SMEMBERS(k string) ([]interface{}, error) {
	v, err := d.conn.Do("SMEMBERS", k)
	if err != nil {
		log.E("[Redis] SMEMBERS fail %s err:%v", k, err)
	}
	if err == nil {
		return v.([]interface{}), err
	}
	return nil, nil
}
