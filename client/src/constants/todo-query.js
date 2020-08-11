import gql from 'graphql-tag'

export const GET_TODOS = gql`
  query {
    todos {
      id
      text
      done
      doneAt
      createAt
    }
  }
`
