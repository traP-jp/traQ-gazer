<script setup lang="ts">
import type { WordsList } from '../apis/generated'
import WordListItem from '../components/WordListItem.vue'

withDefaults(
  defineProps<{
    words: WordsList
    isLoading?: boolean
    errorMessage?: string
  }>(),
  {
    isLoading: false,
    errorMessage: ''
  }
)
const emit = defineEmits<{
  update: []
}>()

const update = () => {
  emit('update')
}
</script>

<template>
  <div v-if="errorMessage" :class="[$style.state, $style.error]" role="alert">
    {{ errorMessage }}
  </div>
  <div v-else-if="isLoading" :class="$style.state" role="status">読み込み中...</div>
  <div v-else-if="words.length === 0" :class="$style.state">登録単語はまだありません</div>
  <ul v-else :class="$style.wordList" aria-label="登録単語一覧">
    <word-list-item v-for="item in words" :key="item.word" :item="item" @update="update" />
  </ul>
</template>

<style module>
.wordList {
  width: 100%;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
  list-style: none;
}

.state {
  padding: 28px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--muted-text-color);
  background: var(--surface-color);
  text-align: center;
  font-weight: 700;
}

.error {
  color: var(--danger-color);
}

</style>
