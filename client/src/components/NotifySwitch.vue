<script setup lang="ts">
import { Icon } from '@iconify/vue'

defineProps<{ title: string }>()

const model = defineModel<boolean>('notify', { required: true })
</script>

<template>
  <button
    type="button"
    :class="[$style.switch, model && $style.active]"
    :aria-pressed="model"
    :aria-label="`${title}の通知を${model ? 'オフ' : 'オン'}にする`"
    @click="model = !model"
  >
    <span :class="$style.title">{{ title }}</span>
    <span :class="$style.status">
      <Icon
        :icon="model ? 'mdi:notifications-active' : 'mdi:notifications-off'"
        width="22"
        height="22"
        :class="$style.notify"
      />
      <span>{{ model ? 'ON' : 'OFF' }}</span>
    </span>
  </button>
</template>

<style module>
.switch {
  min-width: 148px;
  padding: 12px;
  display: grid;
  gap: 8px;
  color: var(--text-color);
  background: var(--surface-color);
  border-color: var(--border-color);
  text-align: left;
}

.switch:hover {
  border-color: var(--primary-color);
}

.active {
  border-color: color-mix(in srgb, var(--primary-color), transparent 55%);
  background: color-mix(in srgb, var(--primary-color), var(--surface-color) 96%);
}

.title {
  color: var(--muted-text-color);
  font-size: 0.78rem;
  font-weight: 700;
}

.status {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  color: var(--heading-color);
  font-size: 0.95rem;
  font-weight: 800;
}

.notify {
  color: var(--primary-color);
}
</style>
