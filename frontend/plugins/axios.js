export default function ({ $axios, app, store }) {
  $axios.onRequest(config => {
    if (store.state.token) {
      config.headers.common['X-RBank-Token'] = store.state.token
    }
  })
}