<template>
  <div id="app" class="login-bg">
    <v-content>
      <v-container
        class="fill-width"
        fluid
      >
        <v-row style="margin-top: 22%;"
          align="center"
          justify="center"
        >
          <v-col
            cols="12"
            sm="8"
            md="4"
          >
            <v-card class="elevation-12 mt-5">
              <v-toolbar
                color="primary"
                flat
              >
                <v-toolbar-title><span class="font-weight-black">Razer Bank</span>
                
                <v-icon class="mb-1" large>mdi-pine-tree</v-icon></v-toolbar-title>
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
  </div>
</template>

<script>
export default {
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

.login-bg {
  background: url('../assets/background.jpg');
  position: absolute;
  background-size: cover;
  height: 100%;
  width: 100%;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  z-index: 0;
}

.login-container {
  text-align: center;
  margin: auto;
  width: 300px;
  height: 286px;
  padding: 10px;
  background: rgba(0, 0, 0, 0.5);
  top: 20%;
  position: relative;
  -webkit-border-radius: 4px;
  -moz-border-radius: 4px;
  border-radius: 4px;
  color: white;
}
</style>
