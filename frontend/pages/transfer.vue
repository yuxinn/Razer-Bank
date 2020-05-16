<template>
  <div class="transfer">
    <p class="mt-4 display-1 font-weight-black">Transfer Funds</p>
    <v-sheet 
      color="grey lighten-3"
      class="mb-1 d-flex"
      height="auto"
    >
    </v-sheet>
      <v-container class="my-4 mx-4">
        <v-row>
          <v-col :cols="6">
            <v-row>
              <v-col :cols="12" class="pb-0"><p class="title font-weight-black">From</p></v-col>
              <v-col :cols="12" class="pt-0" style="margin-top:-15px">
                <v-radio-group v-model="fundAcc">
                  <v-radio
                    class="mb-4"
                    v-for="account in savingsWithMoney"
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
            <v-row>
              <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Target Account</p></v-col>
              <v-col :cols="12" class="pt-0" style="margin-top:-15px">
                <v-radio-group 
                  v-if="targetType=='own'"
                  v-model="targetAcc" 
                >
                  <v-radio
                    class="mb-4"
                    v-for="account in savings"
                    :key="account.id+account.availableBalance"
                    :label="`${account.name} #${account.id} \nBalance: $${parseInt(account.availableBalance).toFixed(2)}`"
                    :value="account.id"
                    :disabled="account.id==fundAcc"
                  ></v-radio>
                </v-radio-group>
                <v-text-field
                  v-else
                  class="mt-4"
                  label="Account #"
                  v-model="targetAcc"
                  placeholder="FF12452"
                  outlined
                  style="width: 60%"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
        <v-row>
          <v-col :cols="6">
            <v-row>
              <v-col :cols="12" class="pb-0"><p class="title font-weight-black">Target</p></v-col>
              <v-col :cols="12" class="pt-0" style="margin-top:-15px">
                <v-radio-group v-model="targetType" row @change="targetAcc=''">
                  <v-radio label="My Account" value="own"></v-radio>
                  <v-radio label="Others" value="other"></v-radio>
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
                style="width: 60%"
              ></v-text-field>
            </v-col>
          </v-col>
        </v-row>

        <v-row justify="center" class="mt-5">
          <v-btn
            :loading="loading"
            class="logout-btn mt-4"
            @click="transfer"
          >
            Transfer
          </v-btn>
        </v-row>
      </v-container>

    <v-snackbar
      class="mt-5"
      v-model="success"
      :top="true"
      multi-line
    >
      Transfer ${{parseInt(amount).toFixed(2)}} Successful<br>
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
      savings: [],
      fundAcc: '',
      amount: '',
      targetAcc: '',
      targetType: 'own',
      transactionId: ''
    }
  },
  mounted() {
    this.success = false
    this.retrieveAccounts()
  },
  computed: {
    savingsWithMoney() {
      var selected = this.savings.filter((el) => {
        return el.availableBalance > 0 });
      return selected
    },
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
    async transfer() {
      const url = `https://bank.ntucbee.click/bank/client/savings?bankaccountid=${this.fundAcc}`
      const data = {
        "amount": parseInt(this.amount),
        "target": this.targetAcc
      }
      try {
        this.loading = true
        let resp = await this.$axios({
          method: 'PUT',
          url: url,
          data: data
        })
        this.transactionId = resp.data.transactionId
        await this.retrieveAccounts()
        this.success = true
      } catch(err) {
        console.error(err)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>