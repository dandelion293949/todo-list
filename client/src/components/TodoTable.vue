<template>
  <v-container>
    <!-- 入力フォーム -->
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-container>
          <h2 v-if="isCreate">Todoを登録する</h2>
          <h2 v-if="!isCreate">Todoを更新する</h2>
          <v-form ref="form" v-model="valid" lazy-validation>
            <!-- Todo -->
            <v-text-field
              v-model="todo.text"
              :rules="textRules"
              :counter="50"
              label="Todo"
              required
            ></v-text-field>
            <v-checkbox
              v-if="!isCreate"
              v-model="todo.done"
              label="Done"
              required
            ></v-checkbox>
            <!-- 登録ボタン -->
            <v-btn
              v-if="isCreate"
              :disabled="!valid"
              @click="createTodo"
            >
              登録
            </v-btn>
            <!-- 更新ボタン -->
            <v-btn
              v-if="!isCreate"
              :disabled="!valid"
              @click="updateTodo"
            >
              更新
            </v-btn>
            <v-btn @click="clear">クリア</v-btn>
          </v-form>
        </v-container>
      </v-card>
    </v-dialog>

    <!--ツールバー-->
    <v-toolbar flat color="grey lighten-2">
      <v-toolbar-title>ToDoリスト</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn color="primary" dark class="mb-2" @click="showDialogNew">新規登録</v-btn>
    </v-toolbar>

    <!-- データテーブル -->
    <v-data-table
      :headers="headers"
      :items="todos"
      :pagination.sync="pagination"
      :loading="progress"
      no-data-text="Todoが登録されてません"
      class="elevation-1"
    >
      <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
      <template v-slot:item.actions="{ item }">
        <v-icon
          small
          class="mr-2"
          @click="showDialogUpdate(item.id, item.text, item.done)"
        >
          edit
        </v-icon>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
  import {GET_TODOS} from "../constants/todo-query";
  import {CREATE, UPDATE} from "../constants/todo-mutation";
  
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
        {text: 'Actions', value: 'actions', sortable: false},
      ],
      // データテーブルの初期ソートと表示件数の設定
      pagination: {
        descending: true,
        rowsPerPage: 10,
      },
      todo: {
        id: '',
        text: '',
        done: false,
      },
      // バリデーション
      valid: true,
      textRules: [
        v => !!v || 'Todoの説明は必須項目です',
        v => (v && v.length <= 50) || '説明は50文字以内で入力してください'
      ],
      // ローディングの表示フラグ
      progress: false,
      // ダイアログの表示フラグ
      dialog: false,
      // 新規・更新のフラグ
      isCreate: true,
    }),
    apollo: {
      // すべてのTodoリストを取得
      todos: GET_TODOS
    },
    methods: {
      // 新規登録
      createTodo: function() {
        if (this.$refs.form.validate()) {
          this.progress = true
          this.$apollo.mutate({
            mutation: CREATE,
            variables: {
              createTodoInput: {
                text: this.todo.text
              }
            }
          }).then(() => {
            this.$apollo.queries.todos.fetchMore({
              updateQuery: (previousResult, {fetchMoreResult}) => {
                return {
                  todos: fetchMoreResult.todos
                }
              }
            })
            /* this.dialog = false */
            /* this.progress = false */
          }).catch((error) => {
            console.error(error)
          }).finally(() => {
            this.dialog = false
            this.progress = false
          })
        }
      },
      // 更新
      updateTodo: function () {
        this.progress = true
        this.$apollo.mutate({
          mutation: UPDATE,
          variables: {
            updateTodoInput: {
              id: this.todo.id,
              text: this.todo.text,
              done: this.todo.done,
            }
          }
        }).then(() => {
          this.$apollo.queries.todos.fetchMore({
            updateQuery: (previousResult, {fetchMoreResult}) => {
              return {
                todos: fetchMoreResult.todos
              }
            }
          })
          /* this.dialog = false */
          /* this.progress = false */
        }).catch((error) => {
          console.error(error)
        }).finally(() => {
          this.dialog = false
          this.progress = false
        })
      },
      // 更新ダイアログの表示
      showDialogUpdate: function (id, text, done) {
        this.todo.id = id
        this.todo.text = text
        this.todo.done = done
        this.isCreate = false
        this.dialog = true
      },
      // 入力フォームのクリア
      clear: function () {
        this.$refs.form.reset()
      },
      // 新規登録ダイアログの表示
      showDialogNew: function () {
        /* this.clear() */
        this.isCreate = true
        this.dialog = true
      },
    }
  }
</script>
