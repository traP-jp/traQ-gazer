<script setup lang="ts">
import { ref } from 'vue'
import apiClient from '../apis'
import { WordsAllList } from '../apis/generated'

const list = ref<WordsAllList>()
apiClient.words.getWords().then((value) => (list.value = value))
</script>

<template>
  <p>
    <a href="./words/add">新規単語登録</a> / <a href="../../words">登録単語の一覧</a> /
    <a href="../../">他の人が登録している単語</a>
  </p>
  <h1>Home</h1>
  <p>{{ list }}</p>
</template>
