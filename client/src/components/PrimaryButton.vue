<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
  <button :class="$style.primary">{{ text }}</button>
</template>

<style module>
.primary {
  color: var(--button-text-color);
  background-color: var(--primary-color);
}
</style>
