package model

import (
	"context"
	"fmt"
	"io"
	"math/rand"
)

var todos []*TodoItem

func init() {
	todos = make([]*TodoItem, 0, 20)
}

type TodoServerImpl struct {
	UnimplementedTodoServer
}

func (*TodoServerImpl) CreateTodo(ctx context.Context, t *TodoItem) (*TodoItem, error) {
	id := rand.Int31()
	t.Id = id
	todos = append(todos, t)
	return t, nil
}

func (*TodoServerImpl) CreateTodos(cts Todo_CreateTodosServer) error {
	for {
		t, err := cts.Recv()
		if err == io.EOF {
			return cts.SendAndClose(&TodoItems{Items: todos})
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		id := rand.Int31()
		t.Id = id
		todos = append(todos, t)
	}
}

func (*TodoServerImpl) ReadTodos(_ *Void, rts Todo_ReadTodosServer) error {
	for _, t := range todos {
		if err := rts.Send(t); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
