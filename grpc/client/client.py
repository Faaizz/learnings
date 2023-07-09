from faker import Faker

import grpc
from google.protobuf.empty_pb2 import Empty

import todo_pb2
import todo_pb2_grpc


fake = Faker()
stub = todo_pb2_grpc.TodoStub(grpc.insecure_channel('localhost:8080'))

def create_single():
  t_item = todo_pb2.TodoItem(
    id=0,
    text=fake.name(),
  )

  t_item = stub.createTodo(t_item)

def generate_todo_items():
  for _ in range(0, 10):
    yield todo_pb2.TodoItem(
      id=0,
      text=fake.name(),
    )

def create_multiple():
  todos = generate_todo_items()
  t_items = stub.createTodos(todos)

def read():
  t_items = stub.readTodos(Empty())
  for t_item in t_items:
    print(t_item)

if __name__ == '__main__':
  create_single()
  create_multiple()
  read()
