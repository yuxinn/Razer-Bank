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
    <v-row class="my-3">
      <v-flex class="text-center">
        <v-alert type="error" v-if="!verified">
          Verify your profile to start banking
        </v-alert>
      </v-flex>
    </v-row>

    <!-- Menu -->
    <v-row class="mt-2">
      <v-flex class="text-center">
        <v-sheet
          class="d-flex"
          color="grey lighten-3"
          height="80"
          :elevation="8"
        >
          <v-col v-for="item in items" :key="item.title">
            <v-icon large class="btn is-clickable">{{item.icon}}</v-icon>
            <p style="color: black;" class="body-2 font-weight-light mt-1">{{item.title}}</p>
          </v-col>
        </v-sheet>
      </v-flex>
    </v-row>

    <!-- Accounts -->
    <v-row class="mt-4">
      <p class="mt-4 headline font-weight-black">My Accounts</p>
    </v-row>
  </div>
</template>
<script>
import { mapMutations } from 'vuex'

export default {
  data () {
    return {
      items: [
        {
          icon: 'mdi-bank',
          title: 'Accounts',
          to: '/home'
        },
        {
          icon: 'mdi-cash-multiple',
          title: 'Transfer',
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
        },
        {
          icon: 'mdi-cash-multiple',
          title: 'Loans',
          to: '/home'
        }
      ],
    }
  },
  computed: {
    verified() {
      return this.$store.state.verified
    }
  }
}
</script>

<style scoped>
.btn {
  color: grey;
}
.btn:hover{
  color: #3DDA25;
}
</style>