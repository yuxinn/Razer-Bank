<template>
  <div></div>
</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
import { mapMutations } from 'vuex'

export default {
  data() {
    return {
      token: this.$route.query.token
    }
  },
  mounted() {
    if (!this.token) {
      this.$router.push('/')
    } else {
      this.verifyToken()
    }

  },
  methods: {
    async verifyToken() {
      try {
        const url = 'https://bank.ntucbee.click/auth/verify?token=' + this.token
        let data = (await this.$axios.get(url)).data
        if (data.token == this.token) {
          var user = VueJwtDecode.decode(data.token)
          user.token = data.token
          this.login(user)
          this.$axios.defaults.headers.common['X-RBank-Token'] = data.token
          this.$router.push('/home')
        }
      } catch(err) {
        console.error(err)
      }
    },
    ...mapMutations({
      login: 'login'
    })
  }
}
</script>