<script setup lang="ts">
import { ref, watch } from 'vue'
import { Icon } from '@iconify/vue'

const notify = ref<string>('true')
defineProps<{ title: string }>()

const model = defineModel<boolean>({ required: true })

watch(notify, () => {
  model.value = notify.value === 'true' ? true : false
})
</script>

<template>
  <div class="switch">
    <div>{{ title }}</div>
    <Icon
      :icon="model ? 'mdi:notifications-active' : 'mdi:notifications-off'"
      width="40"
      height="40"
      class="notify"
      @click="model = !model"
    />
  </div>
</template>

<style scoped lang="scss">
.switch {
  display: block;
  padding: 4px;
}
.notify {
  cursor: pointer;
}
</style>
