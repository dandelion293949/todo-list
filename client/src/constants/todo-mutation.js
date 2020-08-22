import gql from 'graphql-tag'

// Todoの新規登録
export const CREATE = gql`
  mutation createTodo($createTodoInput: CreateTodoInput!) {
    createTodo(input: $createTodoInput) {
      id
      text
      done
      doneAt
      createAt
    }
  }
`

// Todoの更新
export const UPDATE = gql`
  mutation updateTodo($updateTodoInput: UpdateTodoInput!) {
    updateTodo(input: $updateTodoInput) {
      id
      text
      done
      doneAt
      createAt
    }
  }
`
