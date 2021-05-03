<template>
  <v-app id="app">
    <v-app-bar app color="#222831" dark>
      <div class="d-flex align-center">
        <v-img
          alt="twistagram Logo"
          class="shrink mr-2"
          contain
          src="../assets/twistagram-logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer> </v-spacer>
      <v-avatar size="50" class="mr-3">
        <v-img :src="userData.profile"></v-img>
      </v-avatar>
      <h1>{{ userData.fullname }}</h1>
    </v-app-bar>

    <v-main>
      <v-layout justify-center>
        <v-flex offset-sm3 md6>
          <v-row no-gutters>
            <v-col class="mt-10" sm="6">
              <v-text-field
                solo
                v-model="searchText"
                @keyup.native="showAlert"
                class="rounded-xl"
                label="Search by Full name"
                prepend-icon="mdi-magnify"
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-card class="mt-15" display="justify-center" width="500">
            <v-container fluid>
              <div v-for="x in this.userObj" :key="x.image">
                <v-row>
                  <v-col sm="2">
                    <v-avatar size="50">
                      <v-img size="50" :src="x.profile"></v-img>
                    </v-avatar>
                  </v-col>
                  <v-col sm="7">
                    <p class="mt-3">{{ x.fullname }}</p>
                  </v-col>
                  <v-col class="mt-2">
                    <v-btn color="primary"> Follow </v-btn>
                  </v-col>
                </v-row>
                <v-col>
                  <v-divider height="auto"> </v-divider>
                </v-col>
              </div>
            </v-container>
          </v-card>
        </v-flex>
      </v-layout>
    </v-main>
  </v-app>
</template>

<style scoped>
#margin {
  margin-top: 100px;
  padding-left: 100px;
}
</style>

<script>
import axios from "axios";
export default {
  data: () => ({
    userId: "",
    userData: [],
    userObj: [],
    searchText: "",

  }),

  mounted() {
    this.getUserId();
    this.getUserData();
  },
  methods: {
    getUserId() {
      this.userId = this.$route.params.userID;
    },

    getUserData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.userId)
        .then((response) => {
          this.userData = response.data.data;
        });
    },

    showAlert: function (e) {
      if (e.keyCode === 13 && this.searchText === "") {
        alert("Field Fullname masih kosong");
      } else if (e.keyCode === 13) {
        axios
          .get(`http://localhost:8081/searchUser/` + this.searchText)
          .then((response) => {
            var searchUserId = response.data.data;

            if (searchUserId != []) {
              for (let index = 0; index < searchUserId.length; index++) {
                if (this.userId != searchUserId[index].ID ){

                  var userSearchObj = {
                    fullname: searchUserId[index].FullName,
                    id: searchUserId[index].ID,
                    profile: searchUserId[index].Profile,
                  }
                  this.userObj.push(userSearchObj);
                }
              }
            }
          });
      }
    },
  },
};
</script>