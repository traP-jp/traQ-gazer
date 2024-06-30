<script setup lang="ts">
import { ref, watch } from 'vue'

const notify = ref<string>('true')
defineProps<{ title: string }>()

const model = defineModel<boolean>({ required: true })

watch(notify, () => {
  model.value = notify.value === 'true' ? true : false
})
</script>

<template>
  <div class="switch">
    {{ title }}
    <div>
      <input type="radio" id="true" :value="true" v-model="notify" />
      <label for="true">オン</label>
    </div>
    <div>
      <input type="radio" id="false" :value="false" v-model="notify" />
      <label for="false">オフ</label>
    </div>
  </div>
</template>

<style scoped lang="scss">
.switch {
  display: flex;
}
</style>
