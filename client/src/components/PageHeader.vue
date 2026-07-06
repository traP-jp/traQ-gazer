<script setup lang="ts"></script>

<template>
  <header :class="$style.header">
    <div :class="$style.inner">
      <router-link to="/" :class="$style.logoLink" aria-label="traQ-gazer ホーム">
        <img src="../assets/traQgazer_logo.svg" alt="traQ-gazer" width="180" height="52" />
      </router-link>
      <span :class="$style.subtitle">単語通知ダッシュボード</span>
    </div>
  </header>
</template>

<style module>
.header {
  width: 100%;
  border-bottom: 1px solid var(--border-color);
  background: var(--surface-color);
}

.inner {
  width: min(1120px, 100%);
  min-height: 72px;
  padding: 12px 24px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
}

.logoLink {
  padding: 8px;
  display: inline-flex;
  align-items: center;
  flex-shrink: 0;
  border-radius: 8px;
  background: #fff;

  & img {
    display: block;
    width: 180px;
    height: auto;
  }
}

.subtitle {
  color: var(--muted-text-color);
  font-size: 0.9rem;
  font-weight: 600;
}

@media screen and (max-width: 640px) {
  .inner {
    min-height: 64px;
    padding: 12px 16px;
    gap: 12px;
  }

  .logoLink img {
    width: 150px;
  }

  .subtitle {
    display: none;
  }
}
</style>
