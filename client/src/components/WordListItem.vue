<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '../apis'
import {
  WordsList,
  WordListItem,
  WordBotSetting,
  WordMeSetting,
  WordDelete
} from '../apis/generated'

import { Icon } from '@iconify/vue'

defineProps<{ item: WordListItem }>()

const words = ref<WordsList>([])

const editingWord = ref('')
const editingIncludeBot = ref(true)
const editingIncludeMe = ref(false)

apiClient.list.getListUserMe().then((res) => (words.value = res))

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
}

const deleteWord = (word: string) => {
  const req: WordDelete = { word: word }

  apiClient.words.deleteWords(req)
}
</script>

<template>
  <td>{{ item.word }}</td>
  <td class="icons">
    <Icon
      :icon="item.includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
    />
  </td>
  <td class="icons">
    <Icon
      :icon="item.includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
    />
  </td>
  <td class="icons">
    <Icon icon="mdi:file-edit" style="cursor: pointer" width="30" height="30" />
    <Icon
      icon="mdi:delete"
      style="cursor: pointer"
      width="30"
      height="30"
      @click="deleteWord(item.word)"
    />
  </td>
</template>

<style scoped lang="scss">
template {
  margin: 5%;
}

.icons {
  text-align: center;
}
</style>
