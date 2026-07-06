<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import NotifySwitch from '../components/NotifySwitch.vue'
import WordList from '../components/WordList.vue'
import ArticleContainer from '../components/ArticleContainer.vue'

import apiClient from '../apis'
import type { WordRequest, WordsList } from '../apis/generated'
import PageContainer from '../components/PageContainer.vue'
import PrimaryButton from '../components/PrimaryButton.vue'

const maxWordLength = 50
const words = ref<WordsList>([])

const newWord = ref('')
const newBotNotify = ref(true)
const newSelfNotify = ref(false)
const isLoadingWords = ref(true)
const isRegistering = ref(false)
const listErrorMessage = ref('')
const formErrorMessage = ref('')

const normalizedWord = computed(() => newWord.value.trim())
const wordLength = computed(() => normalizedWord.value.length)
const wordLengthHelpText = computed(() => `${maxWordLength}文字以内`)
const validationMessage = computed(() => {
  if (wordLength.value > maxWordLength) {
    return `${maxWordLength}文字以内で入力してください`
  }
  return ''
})
const canRegister = computed(
  () => normalizedWord.value.length > 0 && validationMessage.value === '' && !isRegistering.value
)
const registerButtonText = computed(() => (isRegistering.value ? '登録中' : '登録'))

const fetchWords = async () => {
  isLoadingWords.value = true
  listErrorMessage.value = ''
  try {
    words.value = await apiClient.list.getListUserMe()
  } catch {
    listErrorMessage.value = '登録単語を取得できませんでした'
  } finally {
    isLoadingWords.value = false
  }
}

onMounted(() => {
  void fetchWords()
})

const registerNewWord = async () => {
  if (!canRegister.value) {
    return
  }

  const reqBody: WordRequest = {
    word: normalizedWord.value,
    includeBot: newBotNotify.value,
    includeMe: newSelfNotify.value
  }

  isRegistering.value = true
  formErrorMessage.value = ''
  try {
    await apiClient.words.postWords(reqBody)
    newWord.value = ''
    await fetchWords()
  } catch {
    formErrorMessage.value = '登録できませんでした'
  } finally {
    isRegistering.value = false
  }
}
</script>

<template>
  <PageContainer>
    <article-container
      title="新規単語の登録"
      description="以下のフォームで登録した単語がtraQ上に投稿された際、DM で通知を送信します"
    >
      <form :class="$style.form" novalidate @submit.prevent="registerNewWord">
        <div :class="$style.inputGroup">
          <label :class="$style.label" for="new-word">登録する単語</label>
          <input
            id="new-word"
            v-model="newWord"
            type="text"
            placeholder="例: ハッカソン"
            :maxlength="maxWordLength"
            :class="$style.inputForm"
            :aria-invalid="validationMessage !== ''"
            aria-describedby="new-word-meta"
          />
          <div id="new-word-meta" :class="$style.fieldMeta">
            <span :class="[validationMessage && $style.errorText]">
              {{ validationMessage || wordLengthHelpText }}
            </span>
            <span>{{ wordLength }}/{{ maxWordLength }}</span>
          </div>
        </div>

        <div
          :class="$style.settingsPanel"
          role="group"
          aria-labelledby="notification-settings-label"
        >
          <span id="notification-settings-label" :class="$style.label">通知設定</span>
          <div :class="$style.settings">
            <NotifySwitch v-model:notify="newBotNotify" title="Botの投稿" />
            <NotifySwitch v-model:notify="newSelfNotify" title="自分の投稿" />
          </div>
        </div>

        <div :class="$style.actions">
          <PrimaryButton :text="registerButtonText" type="submit" :disabled="!canRegister" />
        </div>
      </form>
      <p v-if="formErrorMessage" :class="[$style.formMessage, $style.error]" role="alert">
        {{ formErrorMessage }}
      </p>
    </article-container>

    <article-container
      title="登録単語の閲覧"
      description="登録した単語の通知設定を変更したり、単語の登録を解除したりできます"
      :class="$style.wordList"
    >
      <div :class="$style.wordList">
        <word-list
          :class="$style.wordList"
          :words="words"
          :is-loading="isLoadingWords"
          :error-message="listErrorMessage"
          @update="fetchWords"
        />
      </div>
    </article-container>
  </PageContainer>
</template>

<style module>
.form {
  display: grid;
  grid-template-columns: minmax(0, 1fr) auto auto;
  align-items: center;
  padding: 20px;
  row-gap: 20px;
  column-gap: 28px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  background: var(--surface-color);
  box-shadow: 0 8px 24px var(--shadow-color);
}

.inputGroup {
  display: grid;
  gap: 8px;
  min-width: 0;
}

.label {
  color: var(--heading-color);
  font-size: 0.86rem;
  font-weight: 800;
}

.inputForm {
  width: 100%;
  min-width: 0;
  padding: 16px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  font-size: inherit;
  color: var(--text-color);
  background: var(--background-color);
}

.inputForm:focus {
  border-color: var(--primary-color);
  outline: 3px solid var(--focus-color);
  outline-offset: 1px;
}

.fieldMeta {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  color: var(--muted-text-color);
  font-size: 0.82rem;
  font-weight: 600;
}

.errorText {
  color: var(--danger-color);
}

.settingsPanel {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.settings {
  display: flex;
  justify-content: center;
  gap: 12px;
}

.actions {
  display: flex;
  align-items: end;
}

.formMessage {
  width: fit-content;
  padding: 8px 12px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
}

.error {
  color: var(--danger-color);
  background: color-mix(in srgb, var(--danger-color), transparent 90%);
}

.wordList {
  overflow-x: auto;
}

@media screen and (max-width: 860px) {
  .form {
    grid-template-columns: 1fr;
    row-gap: 24px;
  }

  .settings {
    justify-content: start;
    flex-wrap: wrap;
  }

  .actions {
    justify-content: stretch;

    & button {
      width: 100%;
    }
  }
}
</style>
