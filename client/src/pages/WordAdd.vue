<script setup lang="ts">
import { ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'
import WordList from '../components/WordList.vue'
import SectionContainer from '../components/SectionContainer.vue'

import apiClient from '../apis'
import { WordRequest, WordsList } from '../apis/generated'
import PageContainer from '../components/PageContainer.vue'

const newBotNotify = ref(true)
const newSelfNotify = ref(false)

const isFailedOpen = ref(false)
const isClearedOpen = ref(false)

const openFailedDialog = () => {
  isFailedOpen.value = true
}
const openClearedDialog = () => {
  isClearedOpen.value = true
}

const words = ref<WordsList>([])
const newWord = ref('')

apiClient.list.getListUserMe().then((res) => (words.value = res))

const registerNewWord = () => {
  if (newWord.value.length === 0) {
    return
  } else if (newWord.value.length > 50) {
    openFailedDialog()
    return
  }

  const reqBody: WordRequest = {
    word: newWord.value,
    includeBot: newBotNotify.value,
    includeMe: newSelfNotify.value
  }

  // wordの登録リクエスト
  apiClient.words
    .postWords(reqBody)
    .catch((v) => console.log(v))
    .then(() => {
      // 登録後のリストで更新
      apiClient.list.getListUserMe().then((res) => (words.value = res))
    })

  newWord.value = ''
  openClearedDialog()
}

apiClient.list.getListUserMe().then((res) => (words.value = res))
</script>

<template>
  <PageContainer>
    <section-container
      title="新規単語の登録"
      description="以下のフォームで登録した単語がtraQ上に投稿された際、DMに通知を送信します。"
    >
      <input
        v-model="newWord"
        type="text"
        placeholder="登録したい単語をここに入力(50文字以内)"
        class="inputForm"
        @keypress.enter="registerNewWord"
      />
      <div class="flex justify-around my-4">
        <h3>通知設定</h3>
        <NotifySwitch :model-value="newBotNotify" title="Botの投稿" />
        <NotifySwitch :model-value="newSelfNotify" title="自分の投稿" />
      </div>
      <div class="registerButton mb-16 mt-4">
        <v-btn :disabled="newWord === ''" @click="registerNewWord">登録</v-btn>
      </div>
    </section-container>

    <section-container title="登録単語の閲覧" description="">
      <word-list :words="words" />
    </section-container>
  </PageContainer>
</template>

<style>
.inputForm {
  width: 90vw;
  min-width: 300px;
  max-width: 700px;
  padding: 1rem;
  margin: 8px;
}
</style>
