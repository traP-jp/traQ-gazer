<script setup lang="ts">
import { WordsList } from '../apis/generated'
import WordListItem from '../components/WordListItem.vue'
defineProps<{ words: WordsList }>()
</script>

<template>
  <table class="wordList">
    <tr>
      <th>単語</th>
      <th>bot</th>
      <th>自分</th>
      <th></th>
    </tr>
    <tr v-for="item in words" :key="item.word">
      <word-list-item :item="item" />
    </tr>
  </table>
</template>

<style scoped lang="scss">
table {
  overflow-x: scroll;
  overflow-y: scroll;
}
.wordList {
  width: 80%;
  white-space: nowrap;
  border-collapse: collapse;
}
th {
  border-bottom: 1px solid black;
  @include dark {
    border-bottom: 1px solid white;
  }
}
</style>
