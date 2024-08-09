<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '../apis'
import { WordListItem, WordBotSetting, WordMeSetting, WordDelete } from '../apis/generated'

import { Icon } from '@iconify/vue'

const props = defineProps<{ item: WordListItem }>()
const emit = defineEmits<{
  update: []
}>()

const update = () => {
  emit('update')
}

const includeBot = ref(props.item.includeBot)
const includeMe = ref(props.item.includeMe)
const isEdit = ref(false)

const editBotNotify = () => {
  if (isEdit.value) {
    includeBot.value = !includeBot.value
  }
}

const editMeNotify = () => {
  if (isEdit.value) {
    includeMe.value = !includeMe.value
  }
}

const editWord = () => {
  isEdit.value = !isEdit.value
  if (!isEdit.value) {
    const editBotBody: WordBotSetting = {
      word: props.item.word,
      includeBot: includeBot.value
    }
    const editMeBody: WordMeSetting = {
      word: props.item.word,
      includeMe: includeMe.value
    }

    apiClient.bot.putWords(editBotBody)
    apiClient.me.putWordsMe(editMeBody)

    update()
  }
}

const deleteWord = () => {
  const req: WordDelete = { word: props.item.word }
  apiClient.words.deleteWords(req)
  update()
}
</script>

<template>
  <td>{{ item.word }}</td>
  <td class="icons">
    <Icon
      :icon="includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
      @click="editBotNotify"
    />
  </td>
  <td class="icons">
    <Icon
      :icon="includeMe ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
      @click="editMeNotify"
    />
  </td>
  <td class="icons">
    <Icon
      :icon="isEdit ? 'mdi:check-bold' : 'mdi:file-edit'"
      style="cursor: pointer"
      width="30"
      height="30"
      @click="editWord"
    />
    <Icon icon="mdi:delete" class="pointer" width="30" height="30" @click="deleteWord" />
  </td>
</template>

<style scoped lang="scss">
template {
  margin: 5%;
}

.icons {
  text-align: center;
}

.pointer {
  cursor: pointer;
}
</style>
