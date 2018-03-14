<template>
  <section class="container">
    <p>ようこそ</p>
    <nuxt-link :to="{name: 'messages'}">Enter</nuxt-link>
  </section>
</template>

<script>
export default {
}
</script>

<style scoped>
.title
{
  margin: 30px 0;
}
.users
{
  list-style: none;
  margin: 0;
  padding: 0;
}
.user
{
  margin: 10px 0;
}
</style>
