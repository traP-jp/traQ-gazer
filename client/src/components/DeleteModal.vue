<script setup lang="ts">
import { ref } from 'vue'
import { VueFinalModal } from 'vue-final-modal'

defineProps<{ isOpen: boolean }>()
const emits = defineEmits<{
  updateIsOpen: [boolean]
}>()

const isModalOpen = ref(false)
const emit = () => {
  isModalOpen.value = !isModalOpen.value
  emits('updateIsOpen', isModalOpen.value)
}

// const deleteWord = () => {
//   const req: WordDelete = { word: props.item.word }
//   apiClient.words.deleteWords(req)
//   update()
// }
</script>

<template>
  <vue-final-modal v-model="isModalOpen">
    <h3>単語の削除</h3>
    <p>単語を消去しますか？</p>
    <slot />
    <button @click="emit">Confirm</button>
  </vue-final-modal>
</template>
