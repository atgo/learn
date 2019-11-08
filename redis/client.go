package learn_redis

import (
    "context"
    "crypto/tls"
    "net/url"
    "time"

    "github.com/go-redis/redis"
)

type RedisConfig struct {
    Url             string        `yaml:"url"`
    QueueName       string        `yaml:"queueName"`
    PubSubChannel   string        `yaml:"pubsubChannel"`
    MaxRetries      int           `yaml:"maxRetries"`
    MinRetryBackoff time.Duration `yaml:"minRetryBackoff"`
    MaxRetryBackoff time.Duration `yaml:"maxRetryBackoff"`
}

func NewClient(ctx context.Context, cnf RedisConfig) (*redis.Client, error) {
    u, err := url.Parse(cnf.Url)
    if err != nil {
        return nil, err
    }

    pass, _ := u.User.Password()
    options := &redis.Options{
        Addr:     u.Host,
        Password: pass,
        DB:       0,
    }

    if nil != cnf.Retry {
        options.MaxRetries = cnf.Retry.MaxRetries
        options.MinRetryBackoff = cnf.Retry.MinRetryBackoff
        options.MaxRetryBackoff = cnf.Retry.MaxRetryBackoff
    }

    if u.Query().Get("ssl") == "true" {
        options.TLSConfig = &tls.Config{
            InsecureSkipVerify: true,
        }
    }

    c := redis.NewClient(options)
    c.WithContext(ctx)

    return c, nil
}
