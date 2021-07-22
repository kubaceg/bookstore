package test

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *MockLogger) Warn(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *MockLogger) Error(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}

func (m *MockLogger) Fatal(ctx context.Context, msg interface{}) {
	m.Called(ctx, msg)
}
