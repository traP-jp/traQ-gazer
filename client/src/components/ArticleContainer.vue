<script setup lang="ts">
defineProps<{ title: string; description?: string }>()
</script>

<template>
  <article>
    <div class="heading">
      <h1>{{ title }}</h1>
      <span v-if="description">{{ description }}</span>
    </div>
    <slot />
  </article>
</template>

<style scoped lang="scss">
.heading {
  display: flex;
  margin: 8px;
  align-items: center;

  h1 {
    flex-shrink: 0;
  }

  * {
    margin: 8px 16px;
  }
  @include sp {
    display: block;
  }
}
</style>
