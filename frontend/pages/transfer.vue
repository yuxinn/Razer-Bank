<template>
  <div class="createaccount">
    <p class="mt-4 display-1 font-weight-black">Transfer Funds</p>
    <v-sheet 
      color="grey lighten-3"
      class="mb-1 d-flex"
      height="auto"
    >
      <v-container class="my-4 mx-4">
      </v-container>
    </v-sheet>
  </div>
</template>

<script>
export default {
  data() {
    return {
      savings: {}
    }
  },
  mounted() {
    this.retrieveAccounts()
  },
  methods: {
    async retrieveAccounts() {
      const url = 'https://bank.ntucbee.click/bank/client/accounts/all'
      try {
        let resp = await this.$axios({
          method: 'get',
          url: url,
          params: { mambukey: this.$store.state.mambuKey }
        })
        this.savings = resp.data.savings
      } catch(err) {
        console.error(err)
      }
    },
  }
}