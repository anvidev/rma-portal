package queue

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	// events
	TicketCreated = "ticket/created"
	TicketClosed  = "ticket/closed"
)

type Queue struct {
	logger   *slog.Logger
	eventCh  chan job
	handlers map[string]func(ctx context.Context, msg any) error
	workers  int
}

type job struct {
	event   string
	payload any

	maxRetries uint
	currentTry uint
}

func (q *Queue) Enqueue(event string, payload any) {
	if _, ok := q.handlers[event]; ok {
		q.eventCh <- job{event: event, payload: payload, maxRetries: 5, currentTry: 1}
	} else {
		q.logger.Warn("[QUEUE] job event does not exist", "event", event)
	}
}

func (q *Queue) Start(ctx context.Context) {
	for i := range q.workers {
		q.logger.Info("[QUEUE] worker " + fmt.Sprintf("%d", i+1) + " starting")
		go func(queue *Queue) {
			for {
				select {
				case <-ctx.Done():
					queue.logger.Info("[QUEUE] done signal received. shutting down queue")
					return
				case qj := <-q.eventCh:
					queue.logger.Info("[QUEUE] job received", "event", qj.event)

					if handler, ok := queue.handlers[qj.event]; ok {
						if err := handler(ctx, qj.payload); err != nil {
							if qj.currentTry < qj.maxRetries {
								q.eventCh <- job{
									event:      qj.event,
									payload:    qj.payload,
									maxRetries: qj.maxRetries,
									currentTry: qj.currentTry + 1,
								}
							}
							queue.logger.Error(
								"[QUEUE] job failed",
								"event",
								qj.event,
								"current_try",
								qj.currentTry,
								"max_retries",
								qj.maxRetries,
								// "payload",
								// qj.payload,
								"error",
								err,
							)
						}
					}
				}
			}
		}(q)
	}
}

func New(opts ...option) *Queue {
	q := Queue{
		logger:   slog.Default(),
		eventCh:  make(chan job, 10),
		handlers: make(map[string]func(ctx context.Context, msg any) error),
		workers:  1,
	}

	for _, opt := range opts {
		opt(&q)
	}

	q.logger.Info("[QUEUE] queue initialized")

	return &q
}

type option func(*Queue)

func WithLogger(logger *slog.Logger) option {
	return func(q *Queue) {
		q.logger = logger
	}
}

func WithWorkers(n int) option {
	return func(q *Queue) {
		q.workers = n
	}
}

func WithEvent[T any](event string, h func(ctx context.Context, payload T) error) option {
	return func(q *Queue) {
		fn := func(ctx context.Context, payload any) error {
			p, ok := payload.(T)
			if !ok {
				return fmt.Errorf("[QUEUE] invalid job payload type for job event registration %s: %+v", event, payload)
			}
			return h(ctx, p)
		}

		q.handlers[event] = fn
	}
}
