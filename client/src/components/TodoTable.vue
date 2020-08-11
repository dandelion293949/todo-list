<template>
  <v-container>
    <!--ツールバー-->
    <v-toolbar flat color="grey lighten-2">
      <v-toolbar-title>ToDoリスト</v-toolbar-title>
    </v-toolbar>

    <!-- データテーブル -->
    <v-data-table
      :headers="headers"
      :items="todos"
      :pagination.sync="pagination"
      no-data-text="Todoが登録されてません"
      class="elevatins-1"
    >

      <template slot="items" slot-scope="props">
        <td>{{ props.item.text }}</td>
        <td>{{ props.item.done }}</td>
        <td>{{ props.item.doneAt }}</td>
        <td>{{ props.item.createAt }}</td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
  import {GET_TODOS} from "../constants/todo-query";
  
  export default {
    name: "TodoTable",
    data: () => ({
      // Todo登録情報
      todos: [],
      // テーブルのヘッダ情報
      headers: [
        {text: 'Todo', value: 'text'},
        {text: '完了', value: 'done'},
        {text: '完了日', value: 'doneAt'},
        {text: '登録日', value: 'createAt'},
      ],
      // データテーブルの初期ソートと表示件数の設定
      pagination: {
        descending: true,
        rowsPerPage: 10,
      },
    }),
    apollo: {
      // すべてのTodoリストを取得
      todos: GET_TODOS
    },
  }
</script>
