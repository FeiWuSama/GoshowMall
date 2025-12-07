package redis

import (
	"context"
	"fmt"
	"github.com/cnchef/gconv"
	"time"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/constants"
)

type AccessToken struct {
	adaptor *adaptor.Adaptor
}

func NewAccessToken(adaptor *adaptor.Adaptor) *AccessToken {
	return &AccessToken{
		adaptor: adaptor,
	}
}

func (a *AccessToken) SetAccessToken(ctx context.Context, cacheKey string, accessToken string, tokenExpire int64) (string, error) {
	result, err := a.adaptor.Redis.Set(ctx, cacheKey, accessToken, time.Duration(tokenExpire-constants.TokenExpire)*time.Second).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (a *AccessToken) GetAccessToken(ctx context.Context, cacheKey string) (string, int64, error) {
	result, err := a.adaptor.Redis.Get(ctx, cacheKey).Result()
	if err != nil {
		return "", 0, err
	}
	s := a.adaptor.Redis.ExpireTime(ctx, cacheKey).String()
	return result, int64(gconv.ToInt(s)), nil
}

type Locker struct {
	adaptor *adaptor.Adaptor
}

func NewLocker(adaptor *adaptor.Adaptor) *Locker {
	return &Locker{
		adaptor: adaptor,
	}
}

func (l Locker) GetLocker(ctx context.Context, lockKey string) (bool, error) {
	result, err := l.adaptor.Redis.SetNX(ctx, lockKey, "1", constants.LockerExpire*time.Minute).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (l Locker) AwaitLock(ctx context.Context, lockKey string) error {
	// 定义一个通知频道（当锁被释放时发布消息）
	notifyChannel := lockKey + ":released"

	// 订阅解锁事件
	pubsub := l.adaptor.Redis.Subscribe(ctx, notifyChannel)
	defer pubsub.Close()

	// 创建一个带超时的 context（防止永久等待）
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 接收来自解锁频道的消息
	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			// 接收到解锁消息
			if msg != nil && msg.Payload == "unlocked" {
				return nil
			}
		case <-timeoutCtx.Done():
			return fmt.Errorf("lock wait timeout for key: %s", lockKey)
		}
	}
}
