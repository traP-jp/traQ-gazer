<script setup lang="ts">
import { ref, watch } from 'vue'
import { Switch } from '@headlessui/vue'

const emit = defineEmits<{
  (e: 'updateSelfNotify', value: boolean): boolean
}>()

const newSelfNotify = ref(false)

const props = defineProps<{
  selfNotify: boolean
}>()

watch(props, () => {
  newSelfNotify.value = props.selfNotify
})

watch(newSelfNotify, () => {
  emit('updateSelfNotify', newSelfNotify.value)
})
</script>

<template>
  <div class="myself">
    <label class="flex flex-col items-center">
      <p>自分の発言を通知する</p>
      <Switch
        v-model="newSelfNotify"
        :class="newSelfNotify ? 'bg-emerald-400' : 'bg-gray-200'"
        class="relative inline-flex h-[38px] w-[74px] shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 p-0"
      >
        <span class="sr-only">自分の発言を通知する</span>
        <span
          :class="newSelfNotify ? 'translate-x-9' : 'translate-x-0'"
          class="pointer-events-none inline-block h-[34px] w-[34px] transform rounded-full bg-white shadow-lg ring-0 transition duration-200 ease-in-out"
        />
      </Switch>
    </label>
  </div>
</template>
