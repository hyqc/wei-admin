package redis

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"runtime"
)

type Client struct {
	Lock *redsync.Redsync
	Cli  *redis.Client
}

type Config struct {
	Host         string `json:"host"`          //地址
	Port         int    `json:"port"`          //端口
	DB           int    `json:"db"`            //库名
	User         string `json:"user"`          //用户名
	Password     string `json:"password"`      //密码
	PoolSize     int    `json:"pool_size"`     //连接池最大socket连接数
	MinIdleConns int    `json:"minIdle_conns"` //闲置连接数量
}

func New(conf *Config) (*Client, error) {

	conn := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Username: conf.User,     //用户名
		Password: conf.Password, //密码
		DB:       conf.DB,       // redis数据库index

		//连接池容量及闲置连接数量
		PoolSize:     4 * runtime.NumCPU(), // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 1,
		MaxIdleConns: runtime.NumCPU(), //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
	})
	ping := conn.Ping(context.Background())
	lock := redsync.New(goredis.NewPool(conn))

	r := &Client{
		Lock: lock,
		Cli:  conn,
	}

	return r, ping.Err()
}
