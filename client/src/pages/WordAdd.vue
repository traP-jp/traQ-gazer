<script setup lang="ts">
import { ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'
import WordList from '../components/WordList.vue'
import SectionContainer from '../components/SectionContainer.vue'

import apiClient from '../apis'
import { WordRequest, WordsList } from '../apis/generated'
import PageContainer from '../components/PageContainer.vue'
import PrimaryButton from '../components/PrimaryButton.vue'

const words = ref<WordsList>([])

const newWord = ref('')
const newBotNotify = ref(true)
const newSelfNotify = ref(false)

apiClient.list.getListUserMe().then((res) => (words.value = res))

const registerNewWord = () => {
  if (newWord.value.length === 0) {
    return
  } else if (newWord.value.length > 50) {
    return
  }

  const reqBody: WordRequest = {
    word: newWord.value,
    includeBot: newBotNotify.value,
    includeMe: newSelfNotify.value
  }

  // wordの登録リクエスト
  apiClient.words.postWords(reqBody).then(() => update())

  newWord.value = ''
}

const update = () => {
  apiClient.list.getListUserMe().then((res) => (words.value = res))
}
</script>

<template>
  <PageContainer>
    <section-container
      title="新規単語の登録"
      description="以下のフォームで登録した単語がtraQ上に投稿された際、DMに通知を送信します。"
      class="aaa"
    >
      <div class="form">
        <input
          v-model="newWord"
          type="text"
          placeholder="登録したい単語をここに入力(50文字以内)"
          class="inputForm"
          @keypress.enter="registerNewWord"
        />
        <div>
          <h3>通知設定</h3>
          <div class="settings">
            <NotifySwitch v-model:notify="newBotNotify" title="Botの投稿" />
            <NotifySwitch v-model:notify="newSelfNotify" title="自分の投稿" />
          </div>
        </div>
        <div class="registerButton mb-16 mt-4">
          <primary-button text="登録" :disabled="newWord === ''" @click="registerNewWord" />
        </div>
      </div>
    </section-container>

    <section-container title="登録単語の閲覧" description="">
      <word-list :words="words" @update="update()" />
    </section-container>
  </PageContainer>
</template>

<style scoped lang="scss">
.aaa {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.inputForm {
  width: 90vw;
  min-width: 300px;
  max-width: 500px;
  padding: 1rem;
  margin: 20px;
  font-size: 1.25rem;
}

.form {
  display: flex;

  @include sp {
    display: block;
  }
}

.settings {
  display: flex;
  padding: 4px;
}
</style>
