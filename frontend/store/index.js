import createPersistedState from 'vuex-persistedstate'

export const state = () => ({
  token: '',
  email: '',
  firstName: '',
  lastName: '',
  mambuKey: '',
  nric: '',
})

export const mutations = {
  logout(state) {
    state.token = ''
    state.email = ''
    state.firstName = ''
    state.lastName = ''
    state.mambuKey = ''
    state.nric = ''
    this.$axios.setToken(false)
  },
  login(state, user) {
    state.token = user.token
    state.email = user.email
    state.firstName = user.firstName
    state.lastName = user.lastName
    state.mambuKey = user.mambuKey
    state.nric = user.nric
  }
}

export const plugins = [
  createPersistedState()
]
