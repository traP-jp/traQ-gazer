<script setup lang="ts">
import { onMounted, ref, useCssModule, watch } from 'vue'
import apiClient from '../apis'
import type { WordListItem, WordBotSetting, WordMeSetting, WordDelete } from '../apis/generated'
import { Icon } from '@iconify/vue'
import NotifySwitch from './NotifySwitch.vue'
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
const editDialog = ref<HTMLDialogElement>()
const deleteDialog = ref<HTMLDialogElement>()
const $style = useCssModule()

onMounted(() => {
  editDialog.value = document.getElementById($style.editDialog) as HTMLDialogElement
  deleteDialog.value = document.getElementById($style.deleteDialog) as HTMLDialogElement
})

watch([includeBot, includeMe], () => {
  isEdit.value = true
})

const openEdit = () => {
  if (editDialog.value) {
    isEdit.value = false
    editDialog.value.showModal()
  }
}

const openDelete = () => {
  if (deleteDialog.value) {
    deleteDialog.value.showModal()
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
  <td>{{ item.word }}</td>
  <td :class="$style.icons">
    <Icon
      :icon="includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
    />
  </td>
  <td :class="$style.icons">
    <Icon
      :icon="includeMe ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="30"
      height="30"
    />
  </td>
  <td :class="$style.icons">
    <Icon :class="$style.pointer" icon="mdi:file-edit" width="30" height="30" @click="openEdit" />
    <Icon :class="$style.pointer" icon="mdi:delete" width="30" height="30" @click="openDelete" />
  </td>

  <dialog :id="$style.editDialog" @click.self="editDialog?.close()">
    <section :class="$style.dialog">
      <h2>通知設定</h2>
      <div :class="$style.settings">
        <NotifySwitch v-model:notify="includeBot" title="Botの投稿" />
        <NotifySwitch v-model:notify="includeMe" title="自分の投稿" />
      </div>
      <div :class="$style.downerButton">
        <form method="dialog">
          <SecondaryButton text="変更する" @click="sendSetting" />
        </form>
      </div>
    </section>
  </dialog>

  <dialog :id="$style.deleteDialog" @click.self="deleteDialog?.close()">
    <section :class="$style.dialog">
      <h2>単語の削除</h2>

      <div :class="$style.downerButton">
        <form method="dialog">
          <SecondaryButton text="削除する" @click="deleteWord" />
        </form>
      </div>
    </section>
  </dialog>
</template>

<style lang="scss" module>
.icons {
  width: 0px;
  text-align: center;
}

.pointer {
  cursor: pointer;
}

#editDialog,
#deleteDialog {
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
  justify-content: center;
  margin: 16px;
}

.downerButton {
  display: flex;
  justify-content: end;
  margin: 8px;
}
</style>
