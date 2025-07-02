package redis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"time"
)

// LockWithContext 分布式互斥锁 控制并发
func LockWithContext(ctx context.Context, sync *redsync.Redsync, name string, t time.Duration, option ...redsync.Option) (mutex *redsync.Mutex, err error) {
	if len(option) == 0 {
		option = append(option, redsync.WithExpiry(t))
	}

	mux := sync.NewMutex(name, option...)
	return mux, mux.LockContext(ctx)
}
