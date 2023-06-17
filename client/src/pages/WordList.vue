<script setup lang="ts">
import { ref } from 'vue'

interface Word {
  word: string
  botNotify: boolean
  selfNotify: boolean
}
const words = ref<Word[]>([
  { word: 'ぷぐま', botNotify: true, selfNotify: true },
  { word: 'C#', botNotify: true, selfNotify: true },
  { word: 'C++', botNotify: false, selfNotify: true },
  { word: 'うわぁ', botNotify: true, selfNotify: true },
  { word: 'traP', botNotify: true, selfNotify: false },
  { word: 'Twitter', botNotify: true, selfNotify: true },
  { word: 'traQ', botNotify: false, selfNotify: false }
])
</script>

<template>
  <header>traQエゴサ支援ツール</header>
  <div class="expression">
    <h1>登録単語の閲覧ページ</h1>
    <p><a href="./words/add">新規単語登録</a> / <a href="./words">登録単語の一覧</a> / <a href="./words">他の人が登録している単語</a></p>
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
        <td v-if="item.botNotify === true">ON</td>
        <td v-if="item.botNotify === false">OFF</td>
        <td v-if="item.selfNotify === true">ON</td>
        <td v-if="item.selfNotify === false">OFF</td>
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