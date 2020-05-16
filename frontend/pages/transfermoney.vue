<template>
  <div class="createaccount">
    <p class="mt-4 display-1 font-weight-black">Transfer Funds</p>
    acc: {{acc}}
    <v-sheet 
      color="grey lighten-3"
      class="mb-1 d-flex"
      height="auto"
    >

      <!-- <v-container class="my-4 mx-4"> -->
        <!-- <v-row> -->
          <!-- <v-col :cols="6">
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
          </v-col> -->
          <!-- <v-col :cols="6">
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
          </v-col> -->
        <!-- </v-row> -->
      <!-- </v-container> -->
    </v-sheet>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
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
  }
  // mounted() {
  //   this.retrieveAccounts()
  // },
  // methods: {
  //   async retrieveAccounts() {
  //     const url = 'https://bank.ntucbee.click/bank/client/accounts/all'
  //     console.log(this.savings)
  //     try {
  //       let resp = await this.$axios({
  //         method: 'get',
  //         url: url,
  //         params: { mambukey: this.$store.state.mambuKey }
  //       })
  //       this.savings = resp.data.savings
  //       console.log(this.savings)
  //     } catch(err) {
  //       console.error(err)
  //     }
  //   },
  // }
}