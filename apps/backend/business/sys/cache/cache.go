package cache

import (
	"context"
	"net/url"
	"time"

	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/redis/go-redis/v9"
)

func Open(cfg config.Cache) (*redis.Client, error) {
	q := url.Values{}
	q.Set("protocol", cfg.Protocol)

	u := url.URL{
		Host:     cfg.Host,
		Path:     cfg.DBName,
		RawQuery: q.Encode(),
		Scheme:   cfg.Scheme,
		User:     url.UserPassword(cfg.User, cfg.Password),
	}

	opts, err := redis.ParseURL(u.String())
	if err != nil {
		return &redis.Client{}, err
	}

	return redis.NewClient(opts), nil
}

func StatusCheck(ctx context.Context, client *redis.Client) error {
	// If the user doesn't give us a deadline set 2 second.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*2)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		if err := client.Ping(ctx).Err(); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 200 * time.Millisecond)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}
