<script setup lang="ts">
import { Icon } from '@iconify/vue'

defineProps<{ title: string }>()

const model = defineModel<boolean>({ required: true })
</script>

<template>
  <div class="switch">
    <div>{{ title }}</div>
    <div class="icon">
      <Icon
        :icon="model ? 'mdi:notifications-active' : 'mdi:notifications-off'"
        width="40"
        height="40"
        class="notify"
        @click="model = !model"
      />
    </div>
  </div>
</template>

<style scoped lang="scss">
.switch {
  padding: 4px;
  justify-content: center;
}

.icon {
  display: flex;
  margin: auto;
}

.notify {
  cursor: pointer;
}
</style>
