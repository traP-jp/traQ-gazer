<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
  <button class="secondary">{{ text }}</button>
</template>

<style scoped lang="scss">
.secondary {
  color: $text-color-dark;
  background-color: $secondary-color;

  @include dark {
    color: $text-color;
    background-color: $secondary-color-dark;
  }
}
</style>
