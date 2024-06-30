<script setup lang="ts"></script>

<template>
  <header>
    <div class="logo">
      <router-link to="/" class="logo">
        <img src="../assets/traQgazer_logo.svg" alt="logo" width="300" height="150" />
      </router-link>
    </div>
    <div class="links">
      <router-link to="/">新規単語登録</router-link> /
      <router-link to="/">登録単語の一覧</router-link>
    </div>
  </header>
</template>

<style scoped lang="scss">
header {
  width: 100%;
  height: 150px;
  padding: 4px;
  display: flex;
  align-items: center;
  border-bottom: 0.5px solid black;

  .logo {
    flex-grow: 2;
    a{
      width: 300px;
    }
  }

  .links {
    color: $primary-color;
    flex-grow: 1;
    a {
      font-size: 1.5rem;
      font-weight: 700;
    }
  }

  @include sp {
    display: block;
    height: fit-content;
  }
}
</style>
