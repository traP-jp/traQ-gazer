<script setup lang="ts">
import { ref, watch } from 'vue'
import { Switch } from '@headlessui/vue'

const emit = defineEmits<{
  (e: 'updeteSelfNotify', value: boolean): boolean
}>()

const newSelfNotify = ref(false)

watch(newSelfNotify, () => {
  emit('updeteSelfNotify', newSelfNotify.value)
})
</script>

<template>
  <div class="myself">
    <label>
      <p>自分の発言を通知する</p>
      <Switch
        v-model="newSelfNotify"
        :class="newSelfNotify ? 'bg-blue-600' : 'bg-gray-200'"
        class="relative inline-flex h-6 w-11 items-center rounded-full"
      >
        <span class="sr-only">自分の発言を通知する</span>
        <span
          :class="newSelfNotify ? 'translate-x-6' : 'translate-x-1'"
          class="inline-block h-4 w-4 transform rounded-full bg-white transition"
        />
      </Switch>
    </label>
  </div>
</template>
