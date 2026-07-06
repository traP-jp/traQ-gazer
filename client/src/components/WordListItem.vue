<script setup lang="ts">
import { ref, watch } from 'vue'
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
const isSaving = ref(false)
const isDeleting = ref(false)
const dialogErrorMessage = ref('')

const editDialog = ref<HTMLDialogElement>()
const deleteDialog = ref<HTMLDialogElement>()

watch(
  () => props.item,
  (item) => {
    includeBot.value = item.includeBot
    includeMe.value = item.includeMe
  }
)

const openEdit = () => {
  if (editDialog.value) {
    includeBot.value = props.item.includeBot
    includeMe.value = props.item.includeMe
    dialogErrorMessage.value = ''
    editDialog.value.showModal()
  }
}

const openDelete = () => {
  if (deleteDialog.value) {
    dialogErrorMessage.value = ''
    deleteDialog.value.showModal()
  }
}

const closeEdit = () => {
  editDialog.value?.close()
}

const closeDelete = () => {
  deleteDialog.value?.close()
}

const sendSetting = async () => {
  const requests: Promise<unknown>[] = []
  if (includeBot.value !== props.item.includeBot) {
    const editBotBody: WordBotSetting = {
      word: props.item.word,
      includeBot: includeBot.value
    }
    requests.push(apiClient.bot.putWords(editBotBody))
  }
  if (includeMe.value !== props.item.includeMe) {
    const editMeBody: WordMeSetting = {
      word: props.item.word,
      includeMe: includeMe.value
    }
    requests.push(apiClient.me.putWordsMe(editMeBody))
  }
  if (requests.length === 0) {
    closeEdit()
    return
  }

  isSaving.value = true
  dialogErrorMessage.value = ''
  try {
    await Promise.all(requests)
    closeEdit()
    update()
  } catch {
    dialogErrorMessage.value = '通知設定を変更できませんでした'
  } finally {
    isSaving.value = false
  }
}

const deleteWord = async () => {
  const req: WordDelete = { word: props.item.word }
  isDeleting.value = true
  dialogErrorMessage.value = ''
  try {
    await apiClient.words.deleteWords(req)
    closeDelete()
    update()
  } catch {
    dialogErrorMessage.value = '削除できませんでした'
  } finally {
    isDeleting.value = false
  }
}
</script>

<template>
  <li :class="$style.item">
    <div :class="$style.wordCell">
      <span :class="$style.word">{{ item.word }}</span>
    </div>
    <div :class="$style.statuses" aria-label="通知状態">
      <div :class="$style.statusItem">
        <span :class="$style.statusLabel">Bot</span>
        <span :class="[$style.statusBadge, item.includeBot ? $style.enabled : $style.disabled]">
          <Icon :icon="item.includeBot ? 'mdi:notifications-active' : 'mdi:notifications-off'" />
          {{ item.includeBot ? 'ON' : 'OFF' }}
        </span>
      </div>
      <div :class="$style.statusItem">
        <span :class="$style.statusLabel">自分</span>
        <span :class="[$style.statusBadge, item.includeMe ? $style.enabled : $style.disabled]">
          <Icon :icon="item.includeMe ? 'mdi:notifications-active' : 'mdi:notifications-off'" />
          {{ item.includeMe ? 'ON' : 'OFF' }}
        </span>
      </div>
    </div>
    <div :class="$style.actions">
      <button
        type="button"
        :class="$style.iconButton"
        :aria-label="`${item.word}の通知設定を編集`"
        title="通知設定を編集"
        @click="openEdit"
      >
        <Icon icon="mdi:file-edit" width="22" height="22" />
      </button>
      <button
        type="button"
        :class="[$style.iconButton, $style.deleteButton]"
        :aria-label="`${item.word}を削除`"
        title="削除"
        @click="openDelete"
      >
        <Icon icon="mdi:delete" width="22" height="22" />
      </button>
    </div>
  </li>

  <Teleport to="body">
    <dialog ref="editDialog" :class="$style.dialog" @click.self="closeEdit">
      <section :class="$style.dialogContent">
        <header :class="$style.dialogHeader">
          <h3>通知設定</h3>
          <p>{{ item.word }}</p>
        </header>
        <div :class="$style.settings">
          <NotifySwitch v-model:notify="includeBot" title="Botの投稿" />
          <NotifySwitch v-model:notify="includeMe" title="自分の投稿" />
        </div>
        <p v-if="dialogErrorMessage" :class="$style.dialogError" role="alert">
          {{ dialogErrorMessage }}
        </p>
        <div :class="$style.downerButton">
          <button type="button" :class="$style.ghostButton" @click="closeEdit">閉じる</button>
          <SecondaryButton
            :text="isSaving ? '変更中' : '変更する'"
            :disabled="isSaving"
            @click="sendSetting"
          />
        </div>
      </section>
    </dialog>

    <dialog ref="deleteDialog" :class="$style.dialog" @click.self="closeDelete">
      <section :class="$style.dialogContent">
        <header :class="$style.dialogHeader">
          <h3>単語の削除</h3>
          <p>{{ item.word }}</p>
        </header>
        <p v-if="dialogErrorMessage" :class="$style.dialogError" role="alert">
          {{ dialogErrorMessage }}
        </p>
        <div :class="$style.downerButton">
          <button type="button" :class="$style.ghostButton" @click="closeDelete">閉じる</button>
          <button
            type="button"
            :class="$style.dangerButton"
            :disabled="isDeleting"
            @click="deleteWord"
          >
            {{ isDeleting ? '削除中' : '削除する' }}
          </button>
        </div>
      </section>
    </dialog>
  </Teleport>
</template>

<style module>
.item {
  padding: 16px;
  display: grid;
  grid-template-columns: minmax(0, 1fr) auto auto;
  align-items: center;
  gap: 16px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  background: var(--surface-color);
  box-shadow: 0 8px 24px var(--shadow-color);
}

.wordCell {
  min-width: 0;
}

.word {
  color: var(--heading-color);
  font-weight: 800;
  overflow-wrap: anywhere;
}

.statuses {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
}

.statusItem {
  display: flex;
  align-items: center;
  gap: 8px;
}

.statusLabel {
  color: var(--muted-text-color);
  font-size: 0.78rem;
  font-weight: 800;
}

.statusBadge {
  width: fit-content;
  min-width: 70px;
  padding: 4px 8px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-radius: 999px;
  font-size: 0.78rem;
  font-weight: 800;
}

.enabled {
  color: var(--success-color);
  background: color-mix(in srgb, var(--success-color), transparent 88%);
}

.disabled {
  color: var(--muted-text-color);
  background: var(--secondary-background-color);
}

.actions {
  display: flex;
  justify-self: end;
  gap: 8px;
}

.iconButton {
  width: 44px;
  height: 44px;
  padding: 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--secondary-color);
  background: var(--surface-color);
  border-color: var(--border-color);
}

.iconButton:hover {
  color: var(--button-text-color);
  background: var(--secondary-color);
}

.deleteButton {
  color: var(--danger-color);
}

.deleteButton:hover {
  color: var(--button-text-color);
  background: var(--danger-color);
}

.dialog {
  margin: auto;
  border: none;
  border-radius: 8px;
  color: var(--text-color);
  background-color: var(--surface-color);
  box-shadow: 0 18px 60px rgb(0 0 0 / 28%);

  &::backdrop {
    background-color: rgb(0 0 0 / 40%);
  }
}

.dialogContent {
  width: fit-content;
  min-width: 320px;
  max-width: 90vw;
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.dialogHeader {
  display: grid;
  gap: 4px;

  & h3 {
    color: var(--heading-color);
    font-size: 1.12rem;
  }

  & p {
    color: var(--muted-text-color);
    overflow-wrap: anywhere;
  }
}

.settings {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 8px;
}

.dialogError {
  color: var(--danger-color);
  font-weight: 700;
}

.downerButton {
  display: flex;
  justify-content: end;
  gap: 12px;
}

.ghostButton {
  color: var(--text-color);
  background: var(--secondary-background-color);
}

.dangerButton {
  color: var(--button-text-color);
  background: var(--danger-color);
}

.dangerButton:not(:disabled):hover {
  background: var(--danger-hover-color);
}

@media screen and (max-width: 720px) {
  .item {
    grid-template-columns: 1fr;
    align-items: start;
  }

  .actions {
    justify-self: start;
  }

  .dialogContent {
    min-width: min(320px, calc(100vw - 32px));
    padding: 20px;
  }
}
</style>
