<script setup lang="ts">
import { ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'

import apiClient from '../apis'
import {
  WordsList,
  WordListItem,
  WordBotSetting,
  WordMeSetting,
  WordDelete
} from '../apis/generated'

defineProps<{ item: WordListItem }>()

// const isFailedOpen = ref(false)
const isClearedOpen = ref(false)

// function openFailedDialog() {
//   isFailedOpen.value = true
// }
const closeClearedDialog = () => {
  isClearedOpen.value = false
}

const openClearedDialog = () => {
  isClearedOpen.value = true
}

const words = ref<WordsList>([])

apiClient.list.getListUserMe().then((res) => (words.value = res))

const isEditOpen = ref(false)
const isDeleteOpen = ref(false)

const editingWord = ref('')
const editingIncludeBot = ref(true)
const editingIncludeMe = ref(false)

apiClient.list.getListUserMe().then((res) => (words.value = res))

const closeEditDialog = () => {
  isEditOpen.value = false
}

const closeDeleteDialog = () => {
  isDeleteOpen.value = false
}

const deleteWord = () => {
  if (editingWord.value !== '') {
    const delBody: WordDelete = {
      word: editingWord.value
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
  if (editingWord.value !== '') {
    const editBotBody: WordBotSetting = {
      word: editingWord.value,
      includeBot: editingIncludeBot.value
    }
    const editMeBody: WordMeSetting = {
      word: editingWord.value,
      includeMe: editingIncludeMe.value
    }

    apiClient.bot.putWords(editBotBody).catch((v) => console.log(v))
    apiClient.me.putWordsMe(editMeBody).catch((v) => console.log(v))
    apiClient.list.getListUserMe().then((res) => (words.value = res))
  }
  openClearedDialog()
}
</script>

<template>
  <td class="">{{ item.word }}</td>
  <td class="">{{ item.includeBot ? 'ON' : 'OFF' }}</td>
  <td class="">{{ item.includeMe ? 'ON' : 'OFF' }}</td>
  <td class="">
    <button>aaa</button>
  </td>

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
                  <NotifySwitch :model-value="editingIncludeBot" title="botの通知を変更" />
                  <NotifySwitch :model-value="editingIncludeMe" title="自分の発言の通知を変更" />
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

<style scoped lang="scss">
template {
  margin: 5%;
}
</style>
