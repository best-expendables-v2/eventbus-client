package consumer_middleware

import (
	"context"

	eventbusclient "github.com/best-expendables-v2/eventbus-client"
	"github.com/best-expendables-v2/trace"
)

func StoreTraceIdIntoContext(next ConsumeFunc) ConsumeFunc {
	return func(ctx context.Context, msg *eventbusclient.Message) {
		ctx = trace.ContextWithRequestID(ctx, msg.Header.TraceId)
		next(ctx, msg)
	}
}

func StoreUserIdIntoContext(next ConsumeFunc) ConsumeFunc {
	return func(ctx context.Context, msg *eventbusclient.Message) {
		ctx = trace.ContextWithUserID(ctx, msg.Header.UserId)
		next(ctx, msg)
	}
}
