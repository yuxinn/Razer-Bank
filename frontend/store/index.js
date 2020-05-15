export const state = () => ({
  token: '',
  verfied: false,
})

export const mutations = {
  logout (state) {
    state.token = ''
  },
  login (state, token) {
    state.token = token
  }
}