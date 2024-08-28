<script setup lang="ts">
import { ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'
import WordList from '../components/WordList.vue'
import SectionContainer from '../components/ArticleContainer.vue'

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
    >
      <section class="form">
        <input
          v-model="newWord"
          type="text"
          placeholder="登録したい単語をここに入力(50文字以内)"
          class="inputForm"
          @keypress.enter="registerNewWord"
        />
        <section>
          <h3>通知設定</h3>
          <div class="settings">
            <NotifySwitch v-model:notify="newBotNotify" title="Botの投稿" />
            <NotifySwitch v-model:notify="newSelfNotify" title="自分の投稿" />
          </div>
        </section>
        <section class="registerButton">
          <primary-button text="登録" :disabled="newWord === ''" @click="registerNewWord" />
        </section>
      </section>
    </section-container>

    <section-container title="登録単語の閲覧">
      <word-list :words="words" @update="update()" />
    </section-container>
  </PageContainer>
</template>

<style scoped lang="scss">
.inputForm {
  width: 90vw;
  min-width: 300px;
  max-width: 500px;
  padding: 0.5rem;
  margin: 20px;
  font-size: inherit;
}

.form {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  padding: 8px;

  @include sp {
    display: block;
  }
}

.settings {
  display: flex;
  padding: 4px;
}
</style>
