package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type OTPRepositoryMock struct {
	Mock mock.Mock
}

func (o *OTPRepositoryMock) Save(ctx context.Context, key string, value string, ttl time.Duration) error {
	ret := o.Mock.Called(ctx, key, value, ttl)

	return ret.Error(0)
}

func (o *OTPRepositoryMock) Get(ctx context.Context, key string) (string, error) {
	ret := o.Mock.Called(ctx, key)

	return ret.Get(0).(string), ret.Error(1)
}
