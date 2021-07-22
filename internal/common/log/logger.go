package log

import "context"

type Logger interface {
	Info(ctx context.Context, msg interface{})
	Warn(ctx context.Context, msg interface{})
	Error(ctx context.Context, msg interface{})
	Fatal(ctx context.Context, msg interface{})
}
