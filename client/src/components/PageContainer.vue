<template>
  <main :class="$style.container">
    <slot />
  </main>
</template>

<style module>
.container {
  width: min(1120px, 100%);
  min-height: calc(100vh - 72px);
  padding: 32px 24px 56px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 36px;
}

@media screen and (max-width: 640px) {
  .container {
    min-height: calc(100vh - 64px);
    padding: 20px 16px 40px;
  }
}
</style>
