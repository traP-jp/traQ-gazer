<script setup lang="ts">
withDefaults(defineProps<{ text: string; type?: 'button' | 'submit' | 'reset' }>(), {
  type: 'button'
})
</script>

<template>
  <button :class="$style.secondary" :type="type">{{ text }}</button>
</template>

<style module>
.secondary {
  color: var(--button-text-color);
  background-color: var(--secondary-color);

  &:not(:disabled):hover {
    background-color: var(--secondary-hover-color);
  }
}
</style>
