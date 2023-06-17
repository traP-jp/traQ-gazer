<script setup lang="ts">
import { ref } from 'vue'
import PageLink from '../components/PageLink.vue'
import apiClient from '../apis'
import { WordsAllList } from '../apis/generated'

const words = ref<WordsAllList>([])

apiClient.words.getWords().then((res) => (words.value = res))
</script>

<template>
  <header>traQエゴサ支援ツール</header>
  <div class="expression">
    <h1>登録単語の閲覧ページ</h1>
    <PageLink />
  </div>
  <div class="table">
    <table class="wordList">
      <tr>
        <th>単語</th>
        <th>bot通知</th>
        <th>自分の発言の通知</th>
        <th>他の登録者</th>
        <th></th>
      </tr>
      <tr v-for="item in words" :key="item.word">
        <td>{{ item.word }}</td>
        <td>{{ item.includeBot ? 'ON' : 'OFF' }}</td>
        <td>{{ item.includeMe ? 'ON' : 'OFF' }}</td>
      </tr>
    </table>
  </div>
</template>

<style>
.expression {
  text-align: left;
  display: flex;
}
.table {
  overflow-x: scroll;
  overflow-y: scroll;
}
.wordList {
  width: 1000px;
  white-space: nowrap;
  border-collapse: collapse;
}
th {
  border-bottom: 1px solid white;
}
p {
  font-size: 20px;
}
</style>
