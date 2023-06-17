<script setup lang="ts">
import { ref } from 'vue'
import BotNotify from '../components/BotNotify.vue'
import SelfNotify from '../components/SelfNotify.vue'
import { TransitionRoot, TransitionChild, Dialog, DialogPanel, DialogTitle } from '@headlessui/vue'

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

interface Word {
  word: string
  botNotify: boolean
  selfNotify: boolean
}

const words = ref<Word[]>([])
const newWord = ref('')

const addWord = () => {
  if (newWord.value.length > 50) {
    openFailedDialog()
  } else if (newWord.value.length <= 50 && newWord.value.length > 0) {
    words.value.push({
      word: newWord.value,
      botNotify: newBotNotify.value,
      selfNotify: newSelfNotify.value
    })
    newWord.value = ''
    openClearedDialog()
  }
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
  <header>
    <p>traQエゴサ支援ツール</p>
  </header>
  <div class="expression">
    <h1>新規単語の登録</h1>
    <p>
      <router-link to="./add">新規単語登録</router-link> /
      <router-link to="../../words">登録単語の一覧</router-link> /
      <router-link to="../../words">他の人が登録している単語</router-link>
    </p>
    <br />
  </div>
  <div>
    <p>以下のフォームで登録した単語がtraQ上に投稿された際、DMに通知を送信します。</p>
  </div>
  <div>
    <label>
      <input
        v-model="newWord"
        type="text"
        placeholder="登録したい単語をここに入力(50文字以内)"
        class="inputForm"
      />
    </label>
  </div>
  <div>
    <BotNotify @updete-bot-notify="(newValue) => updateNewBotNotify(newValue)" />
    <SelfNotify @updete-self-notify="(newValue) => updateNewSelfNotify(newValue)" />
  </div>
  <div class="registerButton">
    <button @click="addWord">登録</button>
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
.expression {
  display: flex;
}
p {
  font-size: 20px;
}
.bot {
  text-align: left;
  display: flex;
}
.myself {
  text-align: left;
  display: flex;
}
.inputForm {
  font-size: 20px;
  width: 60%;
  min-width: 10em;
  max-width: 100%;
  padding: 1.2em;
}
.regiserButton {
  display: flex;
}
</style>
