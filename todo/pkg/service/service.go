package service

import (
	"context"

	"github.com/kujtimiihoxha/todo-gokit-demo/todo/pkg/io"
	"github.com/kujtimiihoxha/todo-gokit-demo/todo/pkg/db"
	"gopkg.in/mgo.v2/bson"
)

// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, error error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
}

type basicTodoService struct{}

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	error = c.Find(nil).All(&t)
	return t, error
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	todo.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	error = c.Insert(&todo)
	return todo, error
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"complete": true}})
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"complete": false}})
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService() TodoService {
	return &basicTodoService{}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
