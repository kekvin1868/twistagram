<template>
  <v-app id="app">
    <v-app-bar app color="#222831" dark>
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../assets/twistagram-logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer> </v-spacer>
      <v-avatar size="50">
        <v-img src="../assets/kenji.jpg"></v-img>
      </v-avatar>
      <p>JohnDoe</p>
    </v-app-bar>

    <v-main>
      <v-container>
        <div class="table">
          <v-simple-table>
            <tr>
              <td>
                <v-img
                  src="../assets/kenji.jpg"
                  max-width="600"
                  max-height="400"
                  class="rounded"
                >
                </v-img>
              </td>
              <div class="dot">
                <v-menu transition="fade-transition">
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn icon x-large color="black" v-on="on" v-bind="attrs">
                      <v-icon x-large color="black"> mdi-dots-vertical </v-icon>
                    </v-btn>
                  </template>
                  <v-list>
                    <v-list-item
                      v-for="item in massages"
                      :key="item.massage"
                      @click="goToUpdatePost"
                    >
                      <v-list-item-title color="black" v-text="item.massage">
                      </v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </div>
              <div class="h1">
                <h1>John Doe + Caption</h1>
              </div>
              <div class="icon">
                <v-btn icon x-large color="pink">
                  <v-icon>mdi-heart</v-icon>
                </v-btn>
                <v-btn icon x-large color="blue">
                  <v-icon>mdi-message</v-icon>
                </v-btn>
                <hr />
                <hr />
              </div>
              <div class="comment">
                <p>User + user comment</p>
              </div>
            </tr>
          </v-simple-table>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped>
.dot {
  padding-top: 10px;
  margin-bottom: 125px;
}

.h1 {
  margin-top: -100px;
  margin-left: -250px;
  font-family: "Times New Roman", Times, serif;
}

.icon {
  display: block;
  padding-top: -10px;
  margin-left: -260px;
}

.table {
  margin-top: 250px;
  margin-left: 390px;
}

.comment {
  padding-top: 10px;
  margin-left: -220px;
}
</style>

<script>
export default {
  data() {
    return {
      massages: [{ massage: "Update" }, { massage: "Delete" }],
    };
  },

  methods: {
    goToUpdatePost() {
      this.$router.push({ path: "/updatepost" });
    },
  },
};
</script>