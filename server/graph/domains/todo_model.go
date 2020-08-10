package domains

import (
  "context"
  "sort"
  "sync"
  "time"

  "github.com/google/uuid"
)

type TodoID string

type Todo struct {
  ID TodoID
  Text string
  Done bool
  DoneAt time.Time
  CreateAt time.Time
}

type TodoRepository interface {
  Get(ctx context.Context, id TodoID) (*Todo, error)
  GetAll(ctx context.Context) ([]*Todo, error)
  Create(ctx context.Context, todo *Todo) (*Todo, error)
  Update(ctx context.Context, todo *Todo) (*Todo, error)
  Delete(ctx context.Context, todo *Todo) (*Todo, error)
}

func NewTodoRepository () TodoRepository {
  return &todoRepository{
    database: map[TodoID]*Todo{},
  }
}

var _ TodoRepository = (*todoRepository)(nil)

type todoRepository struct {
  sync.RWMutex
  database map[TodoID]*Todo
}

func (repo *todoRepository) Get(ctx context.Context, id TodoID) (*Todo, error) {
  repo.RLock()
  defer repo.RUnlock()

  todo, ok := repo.database[id]
  if !ok {
    return nil, ErrNoSuchEntity
  }

  return todo, nil
}

func (repo *todoRepository) GetAll(ctx context.Context) ([]*Todo, error) {
  repo.RLock()
  defer repo.RUnlock()

  list := make([]*Todo, 0,len(repo.database))
  for _, todo := range repo.database {
    list = append(list, todo)
  }

  sort.Slice(list, func(i, j int) bool {
    a := list[i]
    b := list[j]

    return a.CreateAt.UnixNano() > b.CreateAt.UnixNano()
  })

  return list, nil
}

func (repo *todoRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
  if todo.ID != "" {
    return nil, ErrBadRequest
  }

  repo.Lock()
  defer repo.Unlock()

  todo.ID = TodoID(uuid.New().String())
  todo.CreateAt = time.Now()

  repo.database[todo.ID] = todo

  return todo, nil
}

func (repo *todoRepository) Update(ctx context.Context, todo *Todo) (*Todo, error) {
  if todo.ID == "" {
    return nil, ErrBadRequest
  }

  repo.Lock()
  defer repo.Unlock()

  old, ok := repo.database[todo.ID]
  if !ok {
    return nil, ErrNoSuchEntity
  }

  if todo.Done && old.DoneAt.IsZero() {
    todo.DoneAt = time.Now()
  } else if !todo.Done {
    todo.DoneAt = time.Time{}
  }

  repo.database[todo.ID] = todo

  return todo, nil
}

func (repo *todoRepository) Delete(ctx context.Context, todo *Todo) (*Todo, error) {
  if todo.ID == "" {
    return nil, ErrBadRequest
  }

  repo.Lock()
  defer repo.Unlock()

  delete(repo.database, todo.ID)

  return todo, nil

}
