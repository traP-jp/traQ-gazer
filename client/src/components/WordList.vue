<script setup lang="ts">
import { WordsList } from '../apis/generated'
import WordListItem from '../components/WordListItem.vue'

defineProps<{ words: WordsList }>()
const emit = defineEmits<{
  update: []
}>()

const update = () => {
  emit('update')
}
</script>

<template>
  <table :class="$style.wordList">
    <thead>
      <tr>
        <th>単語</th>
        <th>bot</th>
        <th>自分</th>
        <th>編集・削除</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in words" :key="item.word">
        <word-list-item :item="item" @update="update" />
      </tr>
    </tbody>
  </table>
</template>

<style lang="scss" module>
.wordList {
  width: 80%;
  white-space: nowrap;
  border-collapse: collapse;
  border-spacing: 8px;
}
tbody > tr:nth-child(2n) {
  background: $secondary-background-color;
  @include dark {
    background: $secondary-background-color-dark;
  }
}
th {
  padding: 0px 8px;
  border-bottom: 1px solid $text-color;
  @include dark {
    border-bottom: 1px solid $text-color-dark;
  }
}
</style>
