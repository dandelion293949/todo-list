<template>
  <v-app>
    <TodoTable />
  </v-app>
</template>

<script>
import TodoTable from './components/TodoTable';

export default {
  name: 'App',

  components: {
    TodoTable,
  },
};
</script>
