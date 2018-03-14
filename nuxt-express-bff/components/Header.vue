<template>
  <header>
    <div v-if="$store.state.authUser">
      {{ $store.state.authUser.name }} としてログインしています
      <button @click="logout">ログアウト</button>
    </div>
    <div v-else>
      ゲスト
    </div>
  </header>
</template>

<script>
  export default {
    methods: {
      async logout () {
        await this.$store.dispatch('logout')
      }
    }
  }
</script>
