<script setup lang="ts">
import { ref } from 'vue'
import { Switch } from '@headlessui/vue'

const newBotNotify = ref(true)
</script>

<template>
  <div class="bot">
    <label>
      <p>botの発言を通知する</p>
      <Switch
        v-model="newBotNotify"
        :class="newBotNotify ? 'bg-teal-900' : 'bg-teal-700'"
        class="relative inline-flex h-[38px] w-[74px] shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75"
      >
        <span class="sr-only">Use setting</span>
        <span
          aria-hidden="true"
          :class="newBotNotify ? 'translate-x-9' : 'translate-x-0'"
          class="pointer-events-none inline-block h-[34px] w-[34px] transform rounded-full bg-white shadow-lg ring-0 transition duration-200 ease-in-out"
        ></span>
      </Switch>
    </label>
  </div>
</template>
