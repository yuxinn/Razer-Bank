<template>
  <v-app>
    <v-app-bar
      v-if="loggedIn"
      :clipped-left="clipped"
      fixed
      app
    >
      <nuxt-link to="/home">
        <v-btn icon><v-icon class="home-icon">mdi-home</v-icon></v-btn>
      </nuxt-link>


      <v-spacer></v-spacer>
      <img
        style="height: 43px; width: auto; margin-right: 4%"
        src="~/assets/razer_bank.png"
        alt="bank-logo"
        class="mt-2"
      >
      <nuxt-link to="/">
        <v-btn
          class="logout-btn"
          dark
          @click="logout"
        > Log Out </v-btn>
      </nuxt-link>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
      :fixed="fixed"
      app
    >
      <span>&copy; Razer Bank {{ new Date().getFullYear() }}</span>
       <v-spacer></v-spacer><span class="text--secondary">For Bankers, By Farmers</span><v-icon color="primary">mdi-pine-tree</v-icon>
    </v-footer>
  </v-app>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  auth: 'guest',
  data () {
    return {
      clipped: false,
      drawer: false,
      fixed: false,
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'Vuetify.js'
    }
  },
  computed: {
    loggedIn () {
      return this.$store.state.token != ''
    }
  },
  methods: {
    ...mapMutations({
      logout: 'logout',
    })
  }
}
</script>

<style scoped>
.logout-btn {
  color: #3DDA25;
}

.home-icon {
  color: #1E1E1E;
}
.home-icon:hover {
  color: #3DDA25;
}

</style>
