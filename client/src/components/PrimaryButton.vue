<script setup lang="ts">
withDefaults(defineProps<{ text: string; type?: 'button' | 'submit' | 'reset' }>(), {
  type: 'button'
})
</script>

<template>
  <button :class="$style.primary" :type="type">{{ text }}</button>
</template>

<style module>
.primary {
  color: var(--button-text-color);
  background-color: var(--primary-color);

  &:not(:disabled):hover {
    background-color: var(--primary-hover-color);
  }
}
</style>
