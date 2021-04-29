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
      <v-avatar size="60">
        <v-img src="../assets/kenji.jpg"></v-img>
      </v-avatar>
      <h1>Felix</h1>
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
                  <v-col sm="1">
                    <v-divider height="auto" vertical> </v-divider>
                  </v-col>
                  <v-col sm="2">
                    {{ x.fullname }}
                  </v-col>
                </v-row>
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
    searchUserId: [],
    userObj: [],
    searchText: "",
    showProfile: [
      {
        image: "/images/kenji.jpg",
        name: "Felix",
      },
      {
        image: "/images/default-profile.jpg",
        name: "Joh Doe",
      },
    ],

    model: 1,
  }),

  mounted() {
    this.getUserId;
  },
  methods: {
    getUserId() {
      this.userId = this.$route.params.userId;
    },

    showAlert: function (e) {
      if (e.keyCode === 13 && this.searchText === "") {
        alert("Field Fullname masih kosong");
      } else if (e.keyCode === 13) {
        axios
          .get(`http://localhost:8081/searchUser/` + this.searchText)
          .then((response) => {
            this.searchUserId = response.data.data;

            if (this.searchUserId != []) {
              this.userObj = [];
              for (let index = 0; index < this.searchUserId.length; index++) {
                axios
                  .get(
                    `http://localhost:8081/getUserData/` +
                      this.searchUserId[index].ID
                  )
                  .then((response) => {
                    this.userObj.push(response.data.data);
                  });
              }
            }
          });
      }
    },
  },
};
</script>