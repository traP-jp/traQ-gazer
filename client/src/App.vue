<script setup lang="ts">
import PageLink from './components/PageLink.vue'
import TitleLogo from './components/TitleLogo.vue'
</script>

<template>
  <TitleLogo />
  <PageLink />
  <div class="my-4">
    <RouterView />
  </div>
</template>

<style scoped></style>
