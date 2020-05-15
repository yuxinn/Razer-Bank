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
          Profile Verification
          <footer>
            <small>
              <em>Fast and easy</em>
            </small>
          </footer>
        </blockquote>
      </v-flex>
    </v-row>

    <v-row>
      <v-col :cols="7">
        <v-text-field label="First Name" v-model="firstName" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
      <v-col :cols="5">
        <v-text-field label="Last Name" v-model="lastName" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col :cols="12">
        <v-text-field label="Address" v-model="address" :rules="rules" hide-details="auto"></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col :cols="4">
        <v-text-field label="Postal Code" v-model="postal" :rules="rules" hide-details="auto"></v-text-field>
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
          @change="previewImage(true)"
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
          accept="image/png, image/jpeg, image/jpg"
          @change="previewImage(false)"
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      // base64
      front: null,
      back: null,

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
        value => (value && value.length >= 3) || 'Min 3 characters',
      ],
    }
  },
  methods: {
    previewImage: function(front) {
      const preview = document.querySelector('img');
      const file = document.querySelector('input[type=file]').files[0];
      const reader = new FileReader();

      reader.addEventListener("load", () => {
        // convert image file to base64 string
        if (front) { this.front = reader.result; }
        else { this.back = reader.result }
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
          firstName: this.firstName,
          lastName: this.lastName,
          address: this.address,
          postal: this.postal,
          country: this.country.toUpperCase(),
          language: this.language.toUpperCase(),
          icFront: this.front,
          icBack: this.back
        }
        console.log(data)
        let resp = (await this.$axios.post(url), data).data
        console.log(resp)
      } catch(err) {
        console.error(err)
      }
    }
  }
}
</script>