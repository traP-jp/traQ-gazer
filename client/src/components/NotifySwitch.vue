<script setup lang="ts">
import { Icon } from '@iconify/vue'

defineProps<{ title: string }>()

const model = defineModel<boolean>('notify', { required: true })
</script>

<template>
  <div :class="$style.switch">
    <div>{{ title }}</div>
    <div :class="$style.icon">
      <Icon
        :icon="model ? 'mdi:notifications-active' : 'mdi:notifications-off'"
        width="40"
        height="40"
        :class="$style.notify"
        @click="model = !model"
      />
    </div>
  </div>
</template>

<style lang="scss" module>
.switch {
  padding: 4px;
  align-content: center;

  * {
    margin: auto;
  }
}

.icon {
  display: flex;
  margin: auto;
}

.notify {
  cursor: pointer;
}
</style>
