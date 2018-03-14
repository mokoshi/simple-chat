<template>
  <section class="container">
    <div class="load-area">
      <button @click="loadMore">もっと読み込む</button>
    </div>
    <div class="messages-area">
      <div class="message" v-for="msg in $store.state.messages">
        <div>
          <span class="username">{{msg.id}} {{ msg.user.name }}</span>
          <span class="time">{{ msg.created_at }}</span>
        </div>
        <span class="message-body">{{ msg.text }}</span>
      </div>
    </div>
    <div class="post-area">
      <div v-if="$store.state.authUser">
        <textarea v-model="text" class="textarea" maxlength="255" @keydown.meta.13="post"></textarea>
        <button @click="post">投稿</button>
      </div>
      <div v-else>
        投稿するには<nuxt-link :to="{name: 'user-login'}">ログイン</nuxt-link>してください
      </div>
    </div>
  </section>
</template>

<script>
import axios from '~/plugins/axios'

export default {
  data () {
    return {
      text: null
    }
  },
  mounted () {
    setInterval(() => {
      this.$store.dispatch('fetchNewerMessages')
    }, 1000)
  },
  async fetch ({ store }) {
    await store.dispatch('fetchNewerMessages')
  },
  methods: {
    async post () {
      await axios.post(`/api/messages`, {text: this.text})
      await this.$store.dispatch('fetchNewerMessages')
    },
    async loadMore () {
      await this.$store.dispatch('fetchOlderMessages')
    }
  }
}
</script>

<style scoped>
.messages-area {
  margin: auto;
  width: 800px;
  text-align: left;
}
.message {
  margin: 10px;
}
.username {
  color: #2c2d30;
  font-weight: 900;
}
.time {
  margin-left: 5px;
  font-size: 12px;
  color: #717274;
}
.message-body {
  white-space: pre-wrap;
}
  
.post-area {
  margin: auto;
  width: 800px;
  text-align: left;
}

textarea {
  width: 300px;
  height: 50px;
}
button {
  margin: 10px;
}
</style>
