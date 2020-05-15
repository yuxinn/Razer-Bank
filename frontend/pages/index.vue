<template>
  <div id="app">
    <v-app id="inspire">
      <v-app id="inspire">
        <v-content>
          <v-container
            class="fill-width"
            fluid
          >
            <v-row
              align="center"
              justify="center"
            >
              <v-col
                cols="12"
                sm="8"
                md="4"
              >
                <v-card class="elevation-12">
                  <v-toolbar
                    color="primary"
                    light
                    flat
                  >
                    <v-toolbar-title>Razer Bank</v-toolbar-title>
                    <v-spacer></v-spacer>
                  </v-toolbar>
                  <v-card-text>
                    <v-form>
                      <img class="center-img mt-2 mb-4" src="~/assets/razer_bank.png" style="width: 65%;" />
                      <div class="login-box is-clickable" @click="redirect"></div>
                    </v-form>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-content>
      </v-app>
    </v-app>
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo
  },
  data: () => ({
    source: '',
  }),
  methods: {
    redirect() {
      const newUrl = "https://bank.ntucbee.click/auth/razer/login"
      document.location.href = newUrl;
    }
  }
}
</script>

<style scoped>
.login-btn-back {
  display: none;
}

.login-btn-back:hover {
  display: inline;
}

.login-btn:hover {
  display: none;
}

.login-box {
  width: auto;
  height: 70px;
  background: url('../assets/razer_login_btn.png') no-repeat center;
}

.login-box:hover {
  background: url('../assets/razer_login_btn_dark.png') no-repeat center;
}
</style>
