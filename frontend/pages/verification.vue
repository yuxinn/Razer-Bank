<template>
  <div class="verification">

    <v-row>
      <v-flex class="text-center">
        <img
          style="width: 30%"
          src="~/assets/nric_sample.png"
          alt="bank-logo"
          class="mt-2"
        >
        <blockquote class="blockquote">
          Identity Verification
          <footer>
            <small>
              <em>Fast and easy</em>
            </small>
          </footer>
        </blockquote>
      </v-flex>
    </v-row>
    <v-sheet 
      dark
      class="mb-1 d-flex"
      height="auto"
    >
    <v-container class="my-4 mx-4">
    <v-row>
      <v-col :cols="5">
        <v-text-field label="First Name" v-model="firstName" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
      <v-col :cols="4">
        <v-text-field label="Last Name" v-model="lastName" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
      <v-col :cols="3">
        <v-text-field label="NRIC" v-model="nric" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col :cols="12">
        <v-text-field label="Address" v-model="address" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col :cols="4">
        <v-text-field label="Postal Code" v-model="postal" :rules="postal_rules" hide-details="auto"></v-text-field>
      </v-col>
      <v-col :cols="4">
        <v-select
          :items="countries"
          dark
          v-model="country"
          label="Country"
        ></v-select>
      </v-col>
      <v-col :cols="4">
        <v-select
          :items="languages"
          dark
          v-model="language"
          label="Language"
        ></v-select>
      </v-col>
    </v-row>

    <v-row class="ml-1 mt-3">
      <v-col :cols="3" style="border: 1px solid #3DDA25;">
        <img :src="front" height="150" alt="">
      </v-col>
      <v-col>
        <v-file-input 
          label="NRIC Front"
          prepend-icon="mdi-camera"
          outlined dense
          accept="image/png, image/jpeg, image/jpg"
          v-model="frontraw"
          @change="previewImage($event, 'front')"
        >
        </v-file-input>
      </v-col>
    </v-row>
    <v-row class="ml-1 mt-5">
      <v-col :cols="3" style="border: 1px solid #3DDA25;">
        <img :src="back" height="150" alt="">
      </v-col>
      <v-col>
        <v-file-input 
          label="NRIC Back"
          prepend-icon="mdi-camera"
          outlined dense
          v-model="backraw"
          accept="image/png, image/jpeg, image/jpg"
          @change="previewImage($event, 'back')"
        >
        </v-file-input>
      </v-col>
    </v-row>
    <v-row justify="center" class="mt-5">
      <v-btn
        class="logout-btn mt-4"
        dark
        @click="verify"
      >
        Verify Profile
      </v-btn>
    </v-row>
    </v-container>
    </v-sheet>

    <v-snackbar
      class="mt-5"
      v-model="success"
      color="success"
      :top="true"
      :timeout="5000"
    >
      Profile verification successful. Please relogin to start banking.
    </v-snackbar>

    <v-snackbar
      class="mt-5"
      v-model="error"
      color="error"
      :top="true"
      :timeout="5000"
    >
      {{errMsg}}
    </v-snackbar>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  middleware: 'auth',
  data() {
    return {
      // base64
      success: false,
      error: false,
      errMsg: '',

      frontraw: undefined,
      backraw: undefined,

      front: null,
      back: null,

      nric: '',
      firstName: this.$store.state.firstName || '',
      lastName: this.$store.state.lastName || '',
      address: '',
      postal: '',
      country: 'Singapore',
      language: 'English',

      countries: ["Singapore", "Malaysia"],
      languages: [ 
        { text: 'English', value: 'English' },
        { text: '中文', value: 'Mandarin' },
        { text: 'Melayu', value: 'Malay' },
        { text: 'தமிழ்', value: 'Tamil' },
        { text: '한국어', value: 'Korean' },
        { text: '日本語', value: 'Japanese' },
      ],

      rules: [
        value => !!value || 'Required.',
        value => (value && value.length >= 6) || 'Min 2 characters',
      ],
      postal_rules: [
        value => !!value || 'Required.',
        value => (value && value.length >= 6) || 'Min 6 characters',
      ],
    }
  },
  mounted() {
    this.success = false
  },
  methods: {
    previewImage: function(event, type) {
      const reader = new FileReader();
      var file = this.frontraw
      if (type==="back") {
        file = this.backraw
      }

      reader.addEventListener("load", () => {
        // convert image file to base64 string
        if (type==='front') { 
          this.front = reader.result 
        } else { 
          this.back = reader.result 
        }
        // console.log(reader.result)
      }, false);

      if (file) {
        reader.readAsDataURL(file);
      }
    },
    async verify() {
      const url = 'https://bank.ntucbee.click/bank/client/register'
      try {
        const data = {
          nric: this.nric,
          firstName: this.firstName,
          lastName: this.lastName,
          address: this.address,
          postal: this.postal,
          country: this.country.toUpperCase(),
          preferredLanguage: this.language.toUpperCase(),
          icFront: this.front.split(',')[1],
          icBack: this.back.split(',')[1]
        }
        let resp = await this.$axios({
          method: 'post',
          url: url,
          data: data
        })
        console.log(resp)
        this.success = true
        const delay = ms => new Promise(res => setTimeout(res, ms));
        await delay(3200);
        this.logout()
        this.$router.push('/')
      } catch(err) {
        this.error = true
        this.errMsg= err.response.data.error
      }
    },
    ...mapMutations({
      logout: 'logout'
    })
  }
}
</script>