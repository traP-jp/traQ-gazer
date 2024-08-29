<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'

import apiClient from '../apis'
import { WordListItem, WordBotSetting, WordMeSetting, WordDelete } from '../apis/generated'

import { Icon } from '@iconify/vue'
import SecondaryButton from './SecondaryButton.vue'

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
const editDialogNumber = Math.random().toString()
const deleteDialogNumber = Math.random().toString()
const editDialog = ref<HTMLDialogElement>()
const deleteDialog = ref<HTMLDialogElement>()

onMounted(() => {
  editDialog.value = document.getElementById(editDialogNumber) as HTMLDialogElement
  deleteDialog.value = document.getElementById(deleteDialogNumber) as HTMLDialogElement
})

watch([includeBot, includeMe], () => {
  isEdit.value = true
})

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
  if (editDialog.value) {
    isEdit.value = false
    editDialog.value.showModal()
  }
}

const sendSetting = () => {
  if (isEdit.value) {
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
  <td>{{ item.word }} / {{ editDialog?.returnValue }}</td>
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
      :icon="'mdi:file-edit'"
      style="cursor: pointer"
      width="30"
      height="30"
      @click="editWord"
    />
    <Icon
      icon="mdi:delete"
      class="pointer"
      width="30"
      height="30"
      @click="deleteDialog?.showModal()"
    />
  </td>

  <dialog :id="editDialogNumber" @click.self="editDialog?.close()">
    <section class="dialog">
      <h2>通知設定</h2>
      <div class="settings">
        <NotifySwitch v-model:notify="includeBot" title="Botの投稿" />
        <NotifySwitch v-model:notify="includeMe" title="自分の投稿" />
      </div>
      <div class="downerButton">
        <form method="dialog">
          <SecondaryButton text="変更する" @click="sendSetting" />
        </form>
      </div>
    </section>
  </dialog>

  <dialog :id="deleteDialogNumber" @click.self="deleteDialog?.close()">
    <section class="dialog">
      <h2>単語の削除</h2>

      <div class="downerButton">
        <form method="dialog">
          <SecondaryButton text="削除する" @click="deleteWord" />
        </form>
      </div>
    </section>
  </dialog>
</template>

<style scoped lang="scss">
.icons {
  text-align: center;
}

.pointer {
  cursor: pointer;
}

dialog {
  margin: auto;
  border: none;
  border-radius: 16px;
  background-color: $secondary-background-color;

  @include dark {
    color: $text-color-dark;
    background-color: $secondary-background-color-dark;
  }

  &::backdrop {
    background-color: rgba($color: #000, $alpha: 0.4);
  }
}

.dialog {
  width: fit-content;
  min-width: 350px;
  max-width: 90vw;
  padding: 24px;
}

.settings {
  display: flex;
  justify-content: space-around;
  margin: 8px;
}

.downerButton {
  display: flex;
  justify-content: end;
  margin: 8px;
}
</style>
