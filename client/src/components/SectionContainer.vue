<script setup lang="ts">
defineProps<{ title: string; description: string }>()
</script>

<template>
  <section>
    <div class="header">
      <h1>{{ title }}</h1>
      <span>{{ description }}</span>
    </div>
    <slot />
  </section>
</template>

<style scoped lang="scss">
.header {
  display: flex;
  margin: 8px;
  align-items: center;
  * {
    margin: 8px 16px;
  }
  @include sp {
    display: block;
  }
}
</style>
