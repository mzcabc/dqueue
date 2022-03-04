package dqueue

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// Options keeps the settings to setup dqueue.
type Options struct {
	// queue name
	Name string

	// The settings to setup redis connection.
	RedisOpt *redis.Options

	// Daemon thread num, default is 1
	DaemonWorkerNum int
	// Daemon routine interval, default is 100ms, 0 for keep polling
	DaemonWorkerInterval time.Duration

	// Pull message worker num, default is 1
	ConsumeWorkerNum int

	EnableCancel bool
}
