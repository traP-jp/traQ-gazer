<script setup lang="ts">
import { ref } from 'vue'
import { Icon } from '@iconify/vue'
import apiClient from '../apis'
import { WordListItem, WordBotSetting, WordMeSetting } from '../apis/generated'
import DeleteModal from './DeleteModal.vue'

const props = defineProps<{ item: WordListItem }>()
const emit = defineEmits<{
  update: []
}>()

const update = () => {
  emit('update')
}

const isDelete = ref(false)

const editBotNotify = () => {
  const editBotBody: WordBotSetting = {
    word: props.item.word,
    includeBot: !props.item.includeBot
  }
  apiClient.bot.putWords(editBotBody)

  update()
}

const editMeNotify = () => {
  const editMeBody: WordMeSetting = {
    word: props.item.word,
    includeMe: !props.item.includeMe
  }
  apiClient.me.putWordsMe(editMeBody)

  update()
}
</script>

<template>
  <td>{{ item.word }}</td>
  <td class="icons">
    <Icon
      :icon="item.includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
      class="pointer"
      @click="editBotNotify"
    />
  </td>
  <td class="icons">
    <Icon
      :icon="item.includeMe ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
      class="pointer"
      @click="editMeNotify"
    />
  </td>
  <td class="icons">
    <Icon icon="mdi:delete" class="pointer" width="30" height="30" @click="isDelete = true" />
    {{ isDelete }}
  </td>

  <delete-modal :isOpen="isDelete" @updateIsOpen="($event) => (isDelete = $event)" />
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
