<template>
  <div class="createaccount">
    <p class="mt-4 display-1 font-weight-black">Create Accounts</p>
    <v-sheet 
      color="grey lighten-3"
      class="mb-1 d-flex"
      height="auto"
    >
      <v-container class="my-4 mx-4">
        <v-row>
          <v-col :cols="6">
            <v-row>
              <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Account Type</p></v-col>
              <v-col :cols="12" class="pt-0" style="margin-top:-15px">
                <v-radio-group v-model="acc">
                  <v-radio
                    v-for="type in accTypes"
                    :key="type.value"
                    :label="type.label"
                    :value="type.value"
                    @change="amount=0"
                  ></v-radio>
                </v-radio-group>
              </v-col>
            </v-row>
          </v-col>
          <v-col :cols="6">
            <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Interest Rate</p></v-col>
              <v-chip large class="ma-2" color="green"><span style="color:white;">{{ acc=='savings' ? '0.25' : '2.0' }}% per annum</span></v-chip>
              <v-chip v-if="acc=='deposit'" large class="ma-2" color="orange"><span style="color:white;">2 years duration</span></v-chip>
          </v-col>
        </v-row>

        <v-row v-if="acc=='deposit'">
          <v-col :cols="6">
            <v-row>
              <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Funds from Account</p></v-col>
              <v-col :cols="12" class="pt-0" style="margin-top:-15px">
                <v-radio-group v-model="fundAcc">
                  <v-radio
                    class="mb-4"
                    v-for="account in savings"
                    :key="account.id+account.availableBalance"
                    :label="`${account.name} #${account.id} \nBalance: $${parseInt(account.availableBalance).toFixed(2)}`"
                    :value="account.id"
                    :disabled="account.availableBalance<=0"
                  ></v-radio>
                </v-radio-group>
              </v-col>
            </v-row>
          </v-col>
          <v-col :cols="6">
            <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Amount</p></v-col>
            <v-col :cols="10" class="pt-0" style="">
              <v-text-field
                label="Outlined"
                v-model="amount"
                placeholder="$500.00"
                outlined
                style="width: 40%"
              ></v-text-field>
            </v-col>
          </v-col>
        </v-row>

        <v-row justify="center" class="mt-5">
          <v-btn
            :loading="loading"
            class="logout-btn mt-4"
            @click="create"
          >
            Create Account
          </v-btn>
        </v-row>
      </v-container>
    </v-sheet>

    <v-snackbar
      class="mt-5"
      v-model="success"
      :top="true"
      multi-line
    >
      {{accName}} #{{accId}} created<br>
      <template>Transaction ID #{{transactionId}}</template>

      <v-btn
        color="pink"
        text
        timeout="5500"
        @click="success = false"
      >
        Close
      </v-btn>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      success: false,
      loading: false,
      acc: 'savings',
      accTypes: [
        {
          label: "Digital Savings Account",
          value: "savings"
        },
        {
          label: "Fixed Deposit Account",
          value: "deposit"
        }
      ],
      fundAcc: '',
      amount: 0,
      accId: '',
      savings: {},
      transactionId: undefined,
    }
  },
  computed: {
    accName() {
      var match = this.acc
      var selected = this.accTypes.filter((el) => {
        return el.value === this.acc });
      return selected[0].label
    },
    fundAccArray() {
      var match = this.fundAcc
      var selected = this.savings.filter((el) => {
        return el.id === this.fundAcc });
      return selected[0]
    }
  },
  methods: {
    async create() {
      const url = 'https://bank.ntucbee.click/bank/client/accounts'
      const params = { 
        mambukey: this.$store.state.mambuKey,
        create: this.acc
      }
      const data = {
        amount: parseInt(this.amount)
      }
      try {        
        this.loading = true
        let resp = await this.$axios({
          method: 'post',
          url: url,
          params: params,
          data: data
        })
        this.accId = resp.data.savingsAccount.id
        if (this.acc === 'deposit') 
          await this.transferFunds()
        this.success=true
      } catch(err) {
        console.error(err)
      } finally {
        this.loading = false
      }
    },
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
    async transferFunds() {
      const url = `https://bank.ntucbee.click/bank/client/savings?bankaccountid=${this.fundAcc}`
      const data = {
        "amount": parseInt(this.amount),
        "target": this.accId
      }
      try {
        let resp = await this.$axios({
          method: 'PUT',
          url: url,
          data: data
        })
        this.transactionId = resp.data.transactionId
        this.fundAccArray.availableBalance = parseInt(this.fundAccArray.availableBalance) - parseInt(this.amount)
      } catch(err) {
        console.error(err)
      }
    }
  },
  mounted() {
    this.success = false
    this.retrieveAccounts()
  }
}
</script>