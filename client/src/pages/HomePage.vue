<script setup lang="ts">
import { ref } from 'vue'
import apiClient from '../apis'
import { WordsAllList } from '../apis/generated'

const list = ref<WordsAllList>()
apiClient.words.getWords().then((value) => (list.value = value))
</script>

<template>
  <h1>Home</h1>
  <p>{{ list }}</p>
</template>
