<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
  <button :class="$style.secondary">{{ text }}</button>
</template>

<style lang="scss" module>
.secondary {
  color: $text-color-dark;
  background-color: $secondary-color;

  @include dark {
    color: $text-color;
    background-color: $secondary-color-dark;
  }
}
</style>
