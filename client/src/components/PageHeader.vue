<script setup lang="ts"></script>

<template>
  <header>
    <div :class="$style.logo">
      <router-link to="/" :class="$style.logo">
        <img src="../assets/traQgazer_logo.svg" alt="logo" width="300" height="150" />
      </router-link>
    </div>
    <!-- <div :class="$style.links">
      <router-link to="/">新規単語登録</router-link> /
      <router-link to="/">登録単語の一覧</router-link>
    </div> -->
  </header>
</template>

<style lang="scss" module>
header {
  width: 100%;
  height: 150px;
  padding: 0px 4px;
  display: flex;
  align-items: center;
  border-bottom: 0.5px solid black;

  .logo {
    height: 100%;
    flex-grow: 2;
    a {
      width: 300px;
    }
    img {
      background-color: #ddd;
    }
  }

  // .links {
  //   color: $primary-color;
  //   flex-grow: 1;
  //   a {
  //     font-size: 1.5rem;
  //     font-weight: 700;
  //   }
  // }

  @include sp {
    display: block;
    height: fit-content;
  }

  @include dark {
    border-bottom: 0.5px solid white;
  }
}
</style>
