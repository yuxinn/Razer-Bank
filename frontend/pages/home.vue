<template>
  <div class="home">
    <v-row>
      <v-flex class="text-center">
      <img
        style="width: 10%"
        src="~/assets/razer_bank.png"
        alt="bank-logo"
        class="mt-2"
      >
      <blockquote class="blockquote">
        &#8220;Start banking with Razer.&#8221;
        <footer>
          <small>
            <em>&mdash;Razer Bank</em>
          </small>
        </footer>
      </blockquote>
    </v-flex>
    </v-row>

    <!-- Verification -->
    <v-row class="my-3" v-if="!verified">
      <v-flex class="text-center">
        <nuxt-link to="/verification" style="color: inherit; text-decoration: inherit;">
          <v-alert class="is-clickable" type="error" >
            Verify your profile to start banking!
          </v-alert>
        </nuxt-link>
      </v-flex>
    </v-row>
    
    <!-- Menu -->
    <v-row class="mt-2">
          <v-container class="mx-3">
      <v-flex class="text-center">
        <v-sheet
          dark
          class="d-flex"

          height="80"
          :elevation="8"
        >
            <v-col v-for="item in items" :key="item.title">
              <nuxt-link :to="item.to" style="text-decoration: inherit;">
                <v-icon large class="btn is-clickable">{{item.icon}}</v-icon>
              </nuxt-link>
              <p style="color: whitesmoke;" class="body-2 font-weight-light mt-1">{{item.title}}</p>
            </v-col>
        </v-sheet>
      </v-flex>
          </v-container>
    </v-row>

    <!-- Accounts -->

    <Accounts></Accounts>

    <!-- Welcome msg -->
    <v-snackbar
      class="mt-5"
      v-model="welcome"
      color="primary"
      :top="true"
      :timeout="5000"
    >
      Welcome to Razer Bank, {{name}}!
    </v-snackbar>
  </div>
</template>
<script>
import { mapMutations } from 'vuex'
import Accounts from '~/components/Accounts.vue'

export default {
  middleware: 'auth',
  components: {
    Accounts
  },
  data() {
    return {
      welcome: false,
      items: [
        {
          icon: 'mdi-bank',
          title: 'Accounts',
          to: '/create'
        },
        {
          icon: 'mdi-cash-multiple',
          title: 'Transfer',
          to: '/transfer'
        },
        {
          icon: 'mdi-cash-multiple',
          title: 'Loans',
          to: '/home'
        },
        {
          icon: 'mdi-card-bulleted',
          title: 'Cards',
          to: '/home'
        },
        {
          icon: 'mdi-chart-areaspline',
          title: 'Invest',
          to: '/home'
        }
      ],
    }
  },
  computed: {
    name() {
      return this.$store.state.firstName + ' ' + this.$store.state.lastName
    },
    loggedIn() {
      return this.$store.state.token != ''
    },
    verified() {
      return this.$store.state.nric != ''
    }
  },
  created() {
    this.welcome = true
  }
}
</script>

<style scoped>
.btn {
  color: whitesmoke;
}
.btn:hover{
  color: #3DDA25;
}
</style>