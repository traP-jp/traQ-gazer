<script setup lang="ts">
defineProps<{ title: string; description?: string }>()
</script>

<template>
  <article :class="$style.article">
    <div :class="$style.heading">
      <h1>{{ title }}</h1>
      <span v-if="description">{{ description }}</span>
    </div>
    <slot />
  </article>
</template>

<style lang="scss" module>
.article {
  margin: 8px;
}

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
