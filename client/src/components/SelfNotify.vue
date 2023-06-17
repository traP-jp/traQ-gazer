<script lang="ts">
import {ref} from 'vue'
import { Switch } from '@headlessui/vue'

const newSelfNotify = ref(false)

export default {
    components: {newSelfNotify}
}
</script>

<template>
    <div class="myself">
        <label>
        <p>自分の発言を通知する</p>
        <Switch
            v-model="newSelfNotify"
            :class="newSelfNotify ? 'bg-blue-600' : 'bg-gray-200'"
            class="relative inline-flex h-6 w-11 items-center rounded-full"
        >
            <span class="sr-only">Enable notifications</span>
            <span
            :class="newSelfNotify ? 'translate-x-6' : 'translate-x-1'"
            class="inline-block h-4 w-4 transform rounded-full bg-white transition"
            />
        </Switch>
        </label>
    </div>
</template>