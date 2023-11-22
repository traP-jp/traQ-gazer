<script setup lang="ts">
import { ref } from 'vue'
import BotNotify from '../components/BotNotify.vue'
import SelfNotify from '../components/SelfNotify.vue'
import { TransitionRoot, TransitionChild, Dialog, DialogPanel, DialogTitle } from '@headlessui/vue'
import apiClient from '../apis'
import { WordRequest, WordsList } from '../apis/generated'

const newBotNotify = ref(true)
const newSelfNotify = ref(false)

const isFailedOpen = ref(false)
const isClearedOpen = ref(false)

function closeFailedDialog() {
  isFailedOpen.value = false
}
function openFailedDialog() {
  isFailedOpen.value = true
}
function closeClearedDialog() {
  isClearedOpen.value = false
}
function openClearedDialog() {
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

const updateNewBotNotify = (newValue: boolean) => {
  newBotNotify.value = newValue
}
const updateNewSelfNotify = (newValue: boolean) => {
  newSelfNotify.value = newValue
}
</script>

<template>
  <title>新規単語追加ページ</title>
  <div class="m-2 mb-8">
    <h1>新規単語の登録</h1>
  </div>
  <div>
    <p>以下のフォームで登録した単語がtraQ上に投稿された際、DMに通知を送信します。</p>
  </div>
  <div class="mb-4">
    <label>
      <input
        v-model="newWord"
        type="text"
        placeholder="登録したい単語をここに入力(50文字以内)"
        class="inputForm"
        @keypress.enter="registerNewWord"
      />
    </label>
  </div>
  <div class="flex justify-around my-4">
    <BotNotify @update-bot-notify="(newValue: boolean) => updateNewBotNotify(newValue)" />
    <SelfNotify @update-self-notify="(newValue: boolean) => updateNewSelfNotify(newValue)" />
  </div>
  <div class="registerButton mb-16 mt-4">
    <v-btn :disabled="newWord === ''" @click="registerNewWord">登録</v-btn>
  </div>

  <div class="table">
    <table class="wordList">
      <tr>
        <th class="">単語</th>
        <th class="">bot通知</th>
        <th class="">自分の発言の通知</th>
        <!-- <th>他の登録者</th> -->
      </tr>
      <tr v-for="item in words" :key="item.word">
        <td class="">{{ item.word }}</td>
        <td class="">{{ item.includeBot ? 'ON' : 'OFF' }}</td>
        <td class="">{{ item.includeMe ? 'ON' : 'OFF' }}</td>
      </tr>
    </table>
  </div>

  <div>
    <TransitionRoot appear :show="isFailedOpen" as="template">
      <Dialog as="div" @close="closeFailedDialog" class="relative z-10">
        <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-black bg-opacity-25" />
        </TransitionChild>

        <div class="fixed inset-0 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4 text-center">
            <TransitionChild
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <DialogPanel
                class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
              >
                <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                  単語の登録に失敗しました
                </DialogTitle>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">登録する単語は50文字以内に収めてください。</p>
                </div>

                <div class="mt-4">
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeFailedDialog"
                  >
                    閉じる
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
  <div>
    <TransitionRoot appear :show="isClearedOpen" as="template">
      <Dialog as="div" @close="closeClearedDialog" class="relative z-10">
        <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-black bg-opacity-25" />
        </TransitionChild>

        <div class="fixed inset-0 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4 text-center">
            <TransitionChild
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <DialogPanel
                class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
              >
                <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                  単語の登録に成功しました
                </DialogTitle>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    traQ内で登録されると通知が来るようになりました！！！
                  </p>
                </div>

                <div class="mt-4">
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeClearedDialog"
                  >
                    閉じる
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>

<style>
p {
  font-size: 20px;
}
.inputForm {
  font-size: 20px;
  width: 60%;
  min-width: 10em;
  max-width: 100%;
  padding: 1.2em;
}
</style>
