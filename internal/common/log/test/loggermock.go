package test

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Info(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *LoggerMock) Warn(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *LoggerMock) Error(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *LoggerMock) Fatal(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}
