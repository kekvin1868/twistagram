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
      <v-card
        class="mx-auto mt-8"
        max-width="1000"
        outlined
        rounded
        elevation="5">
        
        <v-row class="ma-3">
          <v-col>
            <v-img
              src="../assets/kenji.jpg"
              aspect-ratio="1"
              max-height="800"/>
          </v-col>

          <v-col>
            <div id="contents">
              <v-card elevation="0" height="450px">
                <v-card-text class="ml-n5">
                  <v-list>
                    <p><a href="" class="mt-3 text-decoration-none" @onClick="goToAccount"><b>{{this.postFullname}}</b></a> {{this.postCaption}}</p>
                    <template v-for="i in postComment">
                      <v-list-tile :key="i">
                        <v-list-tile-content>
                          <div>
                            <p><a href="" class="mt-3 text-decoration-none"><b>username</b></a> test</p>
                          </div>
                        </v-list-tile-content>
                      </v-list-tile>
                    </template>
                  </v-list>
                </v-card-text>
                
                <v-footer class="ml-n2">
                  <div>
                    <v-row class="mx-n6 mt-1">
                      <v-btn icon>
                        <v-icon disabled>
                          mdi-cards-heart
                        </v-icon>
                      </v-btn>
                      <p class="">{{(this.postLike).length}} Likes</p>
                    </v-row>
                    <v-row class="mx-n8 mt-n3 mb-n15">
                      <v-col>
                        <v-text-field
                          label="Post New Comment"
                          solo/>
                      </v-col>
                      <v-col class="mt-3">
                        <v-btn icon>
                          <v-icon disabled>
                            mdi-comment
                          </v-icon>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </div>
                </v-footer>
              </v-card>
            </div>
          </v-col>
        </v-row>
      </v-card>               
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

html {
  overflow: hidden !important;
}

.v-card {
  display: flex !important;
  flex-direction: column;
}

.v-card__text {
  flex-grow: 1;
  overflow: auto;
}

.v-text-field {
  font-size: 12px;
  width: 410px;
}
</style>

<script>
import axios from 'axios'

export default {
  mounted() {
    this.getId();
    this.getPostData();
  },
  data() {
    return {
      userId: "",
      postId: "",
      postLike: [],
      postComment: [],
      postCaption: "",
      postFullname: "",
    };
  },

  methods: {
    getId(){
      this.postId = this.$route.params.postId;
    },
    getPostData(){
      axios.get(`http://localhost:8081/getPost/`+this.postId)
        .then(response=>{
            this.postLike = response.data.data.like;
            this.postComment = response.data.data.comment;
            this.postCaption = response.data.data.caption;
            this.userId = response.data.data.user_id;
            this.postFullname = response.data.data.fullname;
        });
    },
    goToAccount(){
      window.alert(this.user_id)
      this.$router.push({path:"/"+this.userId+"/profile"})
    }
  },
};
</script>