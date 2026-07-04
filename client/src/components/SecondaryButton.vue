<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
  <button :class="$style.secondary">{{ text }}</button>
</template>

<style module>
.secondary {
  color: var(--button-text-color);
  background-color: var(--secondary-color);
}
</style>
