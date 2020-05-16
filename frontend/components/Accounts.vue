<template>
  <div>
    <p class="mt-5 headline font-weight-black">My Accounts</p>
    <v-sheet 
      light
      class="d-flex"
      color="grey lighten-3"
      height="auto"
      :elevation="8"
    >
      <v-container class="mx-3">
        <template v-if="!isEmptyObject(savings)">
          <p class="mt-4 title font-weight-black">Savings</p>
          <v-sheet 
            dark
            class="mb-1 d-flex"
            height="auto"
            v-for="acc in savings" :key="acc.id"
          >
            <v-row>
              <v-col :lg="8" :sm="6">
                <span class="align-middle ml-5 mt-3">{{acc.name}}</span><br>
                <span class="align-middle ml-5 mt-3">{{acc.id}}</span>              </v-col>
              <v-col :lg="4" :sm="6">
                  <v-chip class="ma-2" color="green">Balance: $ {{parseInt(acc.availableBalance).toFixed(2)}}</v-chip>
              </v-col>
            </v-row>
          </v-sheet>
        </template>

        <template v-if="!isEmptyObject(fixed)">
          <p class="mt-4 title font-weight-black">Fixed Deposit</p>
          <v-sheet 
            dark
            class="mb-1 d-flex"
            height="auto"
            v-for="acc in fixed" :key="acc.id"
          >
            <v-row>
              <v-col :lg="8" :sm="6">
                <span class="align-middle ml-5 mt-3">{{acc.name}}</span><br>
                <span class="align-middle ml-5 mt-3">{{acc.id}}</span>
              </v-col>
              <v-col :lg="4" :sm="6">
                <v-chip class="ma-2" color="green">Balance: $ {{parseInt(acc.balance).toFixed(2)}}</v-chip>
                <v-chip class="ma-2" color="orange">{{parseInt(acc.interestRate).toFixed(1)}}%</v-chip>
              </v-col>
            </v-row>
          </v-sheet>
        </template>

        <template v-if="!isEmptyObject(loans)">
          <p class="mt-4 title font-weight-black">Loans</p>
          <v-expansion-panels dark multiple class="mb-3">
            <v-expansion-panel
              v-for="acc in loans"
              :key="acc.id"
            >
              <v-expansion-panel-header>
                <p>{{acc.loanName}} #{{acc.id}} 
                  <v-icon class="ml-5" color="primary" left>{{getLoanIcon(acc.accountState)}}</v-icon>
                  {{acc.accountState}}
                </p>
              </v-expansion-panel-header>
              <v-expansion-panel-content> 
                <v-chip class="ma-2" color="red">Fees Balance ${{acc.feesBalance}}</v-chip>
                <v-chip class="ma-2" color="orange">Fees Due ${{acc.feesDue}}</v-chip>
                <v-chip class="ma-2" color="green">Fees Paid ${{acc.feesPaid}}</v-chip>
                <br>
                <v-chip class="ma-2" color="red">Principal Balance ${{acc.principalBalance}}</v-chip>
                <v-chip class="ma-2" color="orange">Principal Due ${{acc.principalDue}}</v-chip>
                <v-chip class="ma-2" color="green">Principal Paid ${{acc.principalPaid}}</v-chip>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
        </template>
      </v-container>



    </v-sheet>

  </div>
</template>

<script>
export default {
  data() {
    return {
      accounts: {},
      savings: {},
      fixed: {},
      loans: {},
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
        this.accounts = resp.data
        this.savings = resp.data.savings
        this.fixed = resp.data.deposit
        this.loans = resp.data.loans
      } catch(err) {
        console.error(err)
      }
    },
    isEmptyObject(obj) {
      return Object.keys(obj).length === 0 && obj.constructor === Object
    },
    getLoanIcon(state) {
      if (state == "APPROVED") return 'mdi-check-bold'
      if (state == "ACTIVE") return 'mdi-credit-card-check'
      return ''
    }
  }
}
</script>