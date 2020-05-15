export default function ({ store, redirect }) {
  if (!store.state.token) {
    console.log(store.state)
    console.log('unauthorized')
    return redirect('/')
  }
}
