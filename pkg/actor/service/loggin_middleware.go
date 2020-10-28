package service

import (
	"context"

	kitLog "github.com/go-kit/kit/log"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
)

// LoggingMiddleware describes a service (as opposed to endpoint) middleware.
type LoggingMiddleware func(Service) Service

type loggingMiddleware struct {
	logger kitLog.Logger
	next   Service
}

// NewLoggingMiddleware takes a logger as a dependency
// and returns a service LoggingMiddleware.
func NewLoggingMiddleware(logger kitLog.Logger) LoggingMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

func (m loggingMiddleware) log(method string, input interface{}, output interface{}, err error) {
	if err == nil {
		return
	}
	if input == nil && output == nil {
		_ = m.logger.Log(
			"method:", method,
			"err:", err,
		)
		return
	}
	if input == nil {
		_ = m.logger.Log(
			"method:", method,
			"output:", output,
			"err:", err,
		)
		return
	}
	if output == nil {
		_ = m.logger.Log(
			"method:", method,
			"input:", input,
			"err:", err,
		)
		return
	}
	_ = m.logger.Log(
		"method:", method,
		"input:", input,
		"output:", output,
		"err:", err,
	)
}

func (m loggingMiddleware) Create(ctx context.Context, newActor domain.Actor) (err error) {
	defer func() {
		m.log("Create", newActor, nil, err)
	}()
	return m.next.Create(ctx, newActor)
}

func (m loggingMiddleware) Destroy(ctx context.Context, id string) (err error) {
	defer func() {
		m.log("Destroy", id, nil, err)
	}()
	return m.next.Destroy(ctx, id)
}

func (m loggingMiddleware) List(ctx context.Context) (output []domain.Actor, err error) {
	defer func() {
		m.log("List", nil, output, err)
	}()
	return m.next.List(ctx)
}

func (m loggingMiddleware) Show(ctx context.Context, id string) (output domain.Actor, err error) {
	defer func() {
		m.log("Show", id, output, err)
	}()
	return m.next.Show(ctx, id)
}

func (m loggingMiddleware) Update(ctx context.Context, id string, updateActor domain.UpdateActorDTO) (err error) {
	defer func() {
		m.log(
			"Show",
			map[string]interface{}{
				"id":          id,
				"updateActor": updateActor,
			},
			nil,
			err,
		)
	}()
	return m.next.Update(ctx, id, updateActor)
}
