import axios from '../plugins/axios'

export const state = () => ({
  authUser: null,
  messages: []
})

export const getters = {
  minMessageId (state) {
    return state.messages.length === 0 ? null : state.messages[0].id
  },
  maxMessageId (state) {
    return state.messages.length === 0 ? null : state.messages[state.messages.length - 1].id
  }
}

export const mutations = {
  SET_USER (state, user) {
    state.authUser = user
  },

  APPEND_MESSAGES (state, messages) {
    state.messages = [...state.messages, ...messages].sort((a, b) => a.id > b.id ? 1 : -1)
  }
}

export const actions = {
  async nuxtServerInit ({ dispatch, commit }, { req }) {
    if (req.session && req.session.auth) {
      commit('SET_USER', req.session.auth.user)
    }
    await dispatch('fetchOlderMessages')
  },

  async fetchNewerMessages ({ state, commit, getters }) {
    const params = {limit: 5, start_id: getters.maxMessageId + 1}
    let { data } = await axios.get(`/api/messages`, {params})
    commit('APPEND_MESSAGES', data)
  },

  async fetchOlderMessages ({ state, commit, getters }) {
    const params = {limit: 5, end_id: getters.minMessageId - 1}
    let { data } = await axios.get(`/api/messages`, {params})
    commit('APPEND_MESSAGES', data)
  },

  async login ({ commit }, { name, password }) {
    let { data } = await axios.post(`/api/login`, {name, password})
    commit('SET_USER', data)
  },

  async signup ({ commit }, { name, password }) {
    let { data } = await axios.post(`/api/signup`, {name, password})
    commit('SET_USER', data)
  },

  async logout ({ commit }) {
    await axios.post('/api/logout')
    commit('SET_USER', null)
  },

  async post ({ commit }, { text }) {
    let { data } = await axios.post(`/api/messages`, {text})
    console.log(data)
  }
}
