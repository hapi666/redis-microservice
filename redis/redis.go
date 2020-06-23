package redis

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type DB interface {
	PushQueue(url []string) error
	PopQueue() (string, error)
	RangeQueue() ([]string, error)
	SisMember(URLmd5 string) (bool, error)
	SadD(crawledURL []string) error
}

type Pool struct {
	*redis.Pool
}

func NewRedisPool(server, password string) (*Pool, error) {
	if server == "" {
		server = ":6379"
	}
	p := redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
	}
	return &Pool{&p}, nil
}

func (p *Pool) PushQueue(urls []string) error {
	conn := p.Get()
	if conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()
	_, err := conn.Do("rpush", "queue", urls)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) PopQueue() (string, error) {
	conn := p.Get()
	if conn.Err() != nil {
		return "", conn.Err()
	}
	defer conn.Close()
	result, err := redis.String(conn.Do("LPOP", "queue"))
	if err == redis.ErrNil {
		return "", errors.New("URL队列已空")
	} else if err != nil {
		return "", err
	}
	return result, nil

}

func (p *Pool) RangeQueue() ([]string, error) {
	conn := p.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()
	result, err := conn.Do("lrange", "queue", "0", "-1")
	if err == redis.ErrNil {
		return nil, redis.ErrNil
	} else if err != nil {
		return nil, err
	}
	if urls, ok := result.([]string); ok {
		return urls, nil
	}
	return nil, errors.New("类型断言失败")
}

func (p *Pool) SisMember(URLmd5 []byte) (bool, error) {
	conn := p.Get()
	if conn.Err() != nil {
		return false, conn.Err()
	}
	defer conn.Close()
	isExist, err := redis.Bool(conn.Do("sismember", "urlmd5", URLmd5))
	if err != redis.ErrNil {
		return false, err
	}
	if !isExist {
		return false, nil
	}
	return true, err
}

func (p *Pool) SadD(crawledURL []byte) error {
	conn := p.Get()
	if conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()
	_, err := conn.Do("sadd", "urlmd5", crawledURL)
	if err != redis.ErrNil {
		return err
	}
	return redis.ErrNil
}
