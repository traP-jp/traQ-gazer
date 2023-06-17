<script setup lang="ts">
import { ref } from 'vue'
import apiClient from '../apis'
import { WordsAllList } from '../apis/generated'

const list = ref<WordsAllList>()
apiClient.words.getWords().then((value) => (list.value = value))
</script>

<template>
  <p>
    <router-link to="./words/add">新規単語登録</router-link> /
    <router-link to="../../words">登録単語の一覧</router-link> /
    <router-link to="../../">他の人が登録している単語</router-link>
  </p>
  <h1>Home</h1>
  <p>{{ list }}</p>
</template>
