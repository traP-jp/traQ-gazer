<script setup lang="ts">
import { ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'
import WordList from '../components/WordList.vue'
import ArticleContainer from '../components/ArticleContainer.vue'

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
    <article-container
      title="新規単語の登録"
      description="以下のフォームで登録した単語がtraQ上に投稿された際、DM で通知を送信します"
    >
      <section :class="$style.form">
        <input
          size="1"
          v-model="newWord"
          type="text"
          placeholder="登録したい単語をここに入力(50文字以内)"
          :class="$style.inputForm"
          @keypress.enter="registerNewWord"
        />
        <section
          style="
            display: flex;
            align-items: center;
            justify-content: space-around;
            flex-shrink: 0;
            gap: 16px;
          "
        >
          <section>
            <h2>通知設定</h2>
            <div :class="$style.settings">
              <NotifySwitch v-model:notify="newBotNotify" title="Botの投稿" />
              <NotifySwitch v-model:notify="newSelfNotify" title="自分の投稿" />
            </div>
          </section>
          <div :class="$style.registerButton">
            <primary-button text="登録" :disabled="newWord" @click="registerNewWord" />
          </div>
        </section>
      </section>
    </article-container>

    <article-container
      title="登録単語の閲覧"
      description="登録した単語の通知設定を変更したり、単語の登録を解除したりできます"
      :class="$style.wordList"
    >
      <div :class="$style.wordList">
        <word-list :class="$style.wordList" :words="words" @update="update()" />
      </div>
    </article-container>
  </PageContainer>
</template>

<style lang="scss" module>
.form {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  padding: 8px;
  gap: 16px;
}

.inputForm {
  width: 55%;
  flex-grow: 2;
  min-width: 350px;
  max-width: 80%;
  padding: 1.25rem;
  border-radius: 8px;
  margin: 0px 4px;
  font-size: inherit;
}

.settings {
  display: flex;
  justify-content: center;
  padding: 8px;
  gap: 8px;
}

.wordList {
  margin: auto;
  overflow-x: scroll;
}
</style>
