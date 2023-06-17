<script setup lang="ts">
import { ref } from 'vue'
import apiClient from '../apis'
import { WordsAllList } from '../apis/generated'
import PageLink from '../components/PageLink.vue'

const list = ref<WordsAllList>()
apiClient.words.getWords().then((value) => (list.value = value))
</script>

<template>
  <PageLink />
  <h1>Home</h1>
  <p>{{ list }}</p>
</template>
