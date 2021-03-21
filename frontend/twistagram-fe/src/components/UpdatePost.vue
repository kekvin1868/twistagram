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
              <div class="caption">
                <td><h1>Caption:</h1></td><br>
                <td> <h1>  User Caption </h1> </td> <br>
              <hr>
              </div>
              <div class="button">
                <td><v-btn color="primary" @click="savePost"> Save </v-btn></td>
                <td><v-btn color="error" @click="savePost"> Cancel </v-btn></td>
              </div>
            </tr>
          </v-simple-table>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped>
.button {
  padding-top: 180px;
  margin-left: 90px;
}

.table {

  margin-top: 150px;
  margin-left: 390px;
}

.caption {
  margin-top: 5px;
  margin-left: 100px;
  padding-bottom: 100px;
}
</style>

<script>

export default {
    methods: {
        savePost(){
            this.$router.push({path: "/showpost"});
        }
    }
}
</script>