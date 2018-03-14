<template>
  <section class="container">
    <form @submit.prevent="login">
      <p class="error" v-if="formError">{{ formError }}</p>
      <p><label>Username: <input type="text" v-model="formName" name="name" /></label></p>
      <p><label>Password: <input type="password" v-model="formPassword" name="password" /></label></p>
      <button type="submit">Login</button>
    </form>
    <br>
    <nuxt-link :to="{name: 'user-signup'}">Signup はこちら</nuxt-link>
  </section>
</template>

<script>
export default {
  middleware: 'authenticated',
  data () {
    return {
      formError: null,
      formName: null,
      formPassword: null
    }
  },
  methods: {
    async login () {
      try {
        await this.$store.dispatch('login', {
          name: this.formName,
          password: this.formPassword
        })
        this.formError = null

        this.$router.push('/messages')
      } catch (e) {
        this.formError = e.message
      }
    }
  }
}
</script>

<style scoped>
</style>
