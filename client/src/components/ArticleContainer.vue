<script setup lang="ts">
defineProps<{ title: string; description?: string }>()
</script>

<template>
  <article :class="$style.article">
    <div :class="$style.heading">
      <h2>{{ title }}</h2>
      <p v-if="description">{{ description }}</p>
    </div>
    <slot />
  </article>
</template>

<style module>
.article {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.heading {
  display: grid;
  gap: 8px;

  & h2 {
    color: var(--heading-color);
    font-size: 1.25rem;
    line-height: 1.3;
  }

  & p {
    max-width: 680px;
    color: var(--muted-text-color);
    font-size: 0.9rem;
  }
}
</style>
