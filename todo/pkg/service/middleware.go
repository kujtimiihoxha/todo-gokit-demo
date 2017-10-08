package service

import (
	context "context"

	log "github.com/go-kit/kit/log"
	io "github.com/kujtimiihoxha/todo-gokit-demo/todo/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(TodoService) TodoService

type loggingMiddleware struct {
	logger log.Logger
	next   TodoService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TodoService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TodoService) TodoService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (t []io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "Add", "todo", todo, "t", t, "error", error)
	}()
	return l.next.Add(ctx, todo)
}
func (l loggingMiddleware) SetComplete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "SetComplete", "id", id, "error", error)
	}()
	return l.next.SetComplete(ctx, id)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "error", error)
	}()
	return l.next.Delete(ctx, id)
}

func (l loggingMiddleware) RemoveComplete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "RemoveComplete", "id", id, "error", error)
	}()
	return l.next.RemoveComplete(ctx, id)
}
