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
                v-model="text"
                @keyup.native="showAlert"
                class="rounded-xl"
                label="Search by Full name"
                prepend-icon="mdi-magnify"
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-card class="mt-15" display="justify-center" width="500">
            <v-container fluid>
              <v-row v-for="x in showProfile" :key="x.image">
                <v-col sm="1">
                  <p>1.</p>
                </v-col>
                <v-col>
                  <v-divider height="auto" vertical> </v-divider>
                </v-col>
                <v-avatar size="60">
                  <img :src="x.image" />
                </v-avatar>
                <v-col sm="6">
                  {{ x.name }}
                </v-col>
                <v-col xs="4">
                  <v-btn color="primary"> Follow </v-btn>
                </v-col>
              </v-row>
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
export default {
  data: () => ({
    text: "",
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
  methods: {
    showAlert: function(e) {
        if(e.keyCode === 13 && this.text === ""){
            alert("Field Fullname masih kosong");
        }else if (e.keyCode===13){
            alert(this.text); // ini nanti aka trigger, 
        }
    },
  },
};
</script>