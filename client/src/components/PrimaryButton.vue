<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
  <button :class="$style.primary">{{ text }}</button>
</template>

<style lang="scss" module>
.primary {
  color: $text-color-dark;
  background-color: $primary-color;

  @include dark {
    color: $text-color;
    background-color: $primary-color-dark;
  }
}
</style>
