<script setup lang="ts">
import { ref } from 'vue'
import apiClient from '../apis'
import BotNotify from '../components/BotNotify.vue'
import SelfNotify from '../components/SelfNotify.vue'
import { WordsList, WordBotSetting, WordMeSetting, WordDelete } from '../apis/generated'
import { TransitionRoot, TransitionChild, Dialog, DialogPanel, DialogTitle } from '@headlessui/vue'
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue'
import { ChevronDownIcon } from '@heroicons/vue/20/solid'

const isEditOpen = ref(false)
const isDeleteOpen = ref(false)
const isClearedOpen = ref(false)

const words = ref<WordsList>([])
const edittingWord = ref('')
const edittingIncludeBot = ref(true)
const edittingIncludeMe = ref(false)

apiClient.list.getListUserMe().then((res) => (words.value = res))

const changeContents = (word: string, bot: boolean, self: boolean) => {
  edittingWord.value = word
  edittingIncludeBot.value = bot
  edittingIncludeMe.value = self
}

const openEditDialog = () => {
  isEditOpen.value = true
}
const closeEditDialog = () => {
  isEditOpen.value = false
}
const openDeleteDialog = () => {
  isDeleteOpen.value = true
}
const closeDeleteDialog = () => {
  isDeleteOpen.value = false
}
const openClearedDialog = () => {
  isClearedOpen.value = true
}
const closeClearedDialog = () => {
  isClearedOpen.value = false
}

const deleteWord = () => {
  if (edittingWord.value !== '') {
    const delBody: WordDelete = {
      word: edittingWord.value
    }

    apiClient.words
      .deleteWords(delBody)
      .catch((v) => console.log(v))
      .then(() => {
        // 登録後のリストで更新
        apiClient.list.getListUserMe().then((res) => (words.value = res))
      })
  }
  closeDeleteDialog()
}

const editWord = () => {
  if (edittingWord.value !== '') {
    const editBotBody: WordBotSetting = {
      word: edittingWord.value,
      includeBot: edittingIncludeBot.value
    }
    const editMeBody: WordMeSetting = {
      word: edittingWord.value,
      includeMe: edittingIncludeMe.value
    }

    apiClient.bot.putWords(editBotBody).catch((v) => console.log(v))
    apiClient.me.putWordsMe(editMeBody).catch((v) => console.log(v))
    apiClient.list.getListUserMe().then((res) => (words.value = res))
  }
  openClearedDialog()
}

const updateNewBotNotify = (newValue: boolean) => {
  edittingIncludeBot.value = newValue
}
const updateNewSelfNotify = (newValue: boolean) => {
  edittingIncludeMe.value = newValue
}
</script>

<template>
  <div class="m-2 mb-8">
    <h1>登録単語の閲覧</h1>
    <PageLink />
  </div>
  <div class="table">
    <table class="wordList">
      <tr>
        <th class="">単語</th>
        <th class="">bot通知</th>
        <th class="">自分の発言の通知</th>
        <th></th>
      </tr>
      <tr v-for="item in words" :key="item.word">
        <td class="">{{ item.word }}</td>
        <td class="">{{ item.includeBot ? 'ON' : 'OFF' }}</td>
        <td class="">{{ item.includeMe ? 'ON' : 'OFF' }}</td>
        <td class="">
          <div class="editDialog">
            <Popover v-slot="{ open }" class="relative">
              <PopoverButton
                :class="open ? '' : 'text-opacity-90'"
                class="group inline-flex items-center rounded-md bg-gray-300 px-3 py-2 text-base font-medium text-white hover:text-opacity-100 focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75"
                @click="changeContents(item.word, item.includeBot, item.includeMe)"
              >
                <span>︙</span>
                <ChevronDownIcon
                  :class="open ? '' : 'text-opacity-70'"
                  class="ml-2 h-5 w-5 text-base transition duration-150 ease-in-out group-hover:text-opacity-80"
                  aria-hidden="true"
                />
              </PopoverButton>
              <transition>
                <PopoverPanel
                  class="absolute left-1/2 z-10 mt-3 w-screen max-w-sm -translate-x-1/2 transform px-4 sm:px-0 lg:max-w-3xl"
                >
                  <div
                    class="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5"
                  >
                    <div class="relative grid gap-8 bg-white p-7 lg:grid-cols-2">
                      <button
                        class="-m-3 flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
                        @click="openEditDialog"
                      >
                        通知設定の変更
                      </button>
                      <button
                        class="-m-3 flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
                        @click="openDeleteDialog"
                      >
                        削除
                      </button>
                    </div>
                  </div>
                </PopoverPanel>
              </transition>
            </Popover>
          </div>
        </td>
      </tr>
    </table>
  </div>

  <div>
    <TransitionRoot appear :show="isDeleteOpen" as="template">
      <Dialog as="div" @close="closeDeleteDialog" class="relative z-10">
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
                  この単語を削除しますか？
                </DialogTitle>

                <div class="mt-4">
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="deleteWord"
                  >
                    削除する
                  </button>
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeDeleteDialog"
                  >
                    キャンセル
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
    <TransitionRoot appear :show="isEditOpen" as="template">
      <Dialog as="div" @close="closeEditDialog" class="relative z-10">
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
                  単語の通知設定の変更
                </DialogTitle>

                <div class="flex justify-around my-4">
                  <BotNotify @updateBotNotify="(newValue: boolean) => updateNewBotNotify(newValue)" />
                  <SelfNotify
                    @updateSelfNotify="(newValue :boolean) => updateNewSelfNotify(newValue)"
                  />
                </div>
                <div class="mt-4">
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="editWord"
                  >
                    更新する
                  </button>
                  <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeEditDialog"
                  >
                    キャンセル
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
                  通知設定の変更ができました
                </DialogTitle>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">引き続き単語の通知が設定に合わせて届きます</p>
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
.inputForm {
  font-size: 20px;
  width: 60%;
  min-width: 10em;
  max-width: 100%;
  padding: 1.2em;
}
</style>
