<template>
  <v-app id="app">
    <v-app-bar app color="#222831" dark>
      <div class="d-flex align-center">
        <a href="">
          <v-img
            alt="Vuetify Logo"
            class="shrink ma-2"
            contain
            src="../assets/twistagram-logo.png"
            transition="scale-transition"
            width="150"
            @click.prevent="goHome"/>
        </a>
      </div>
      <v-spacer></v-spacer>
      <v-avatar size="50">
        <v-img :src="this.visitorAvatar"></v-img>
      </v-avatar>
      <p class="mt-3 ml-3 mr-13">
        <a href=""
          class="text-decoration-none"
          style="color:white;"
          @click.prevent="goToProfile()">{{this.visitorFullname}}</a>
      </p>
    </v-app-bar>

    <v-main>
      <v-card
        class="mx-auto mt-8"
        max-width="1000"
        outlined
        rounded
        elevation="5">
        
        <v-row class="ma-3" v-if="postPhoto != ''">
          <v-col>
            <v-img
              :src="postPhoto"
              aspect-ratio="1"
              max-height="800"/>
          </v-col>

          <v-col>
            <div id="contents">
              <v-card elevation="0" height="450px">
                <v-card-text class="ml-n5">
                  <v-list>
                    <v-list-item>
                      <v-list-item-avatar color="grey darken-3" size="30" class="mb-3 ml-n3"> 
                        <v-img
                          :src="userAvatar"></v-img>
                      </v-list-item-avatar>

                      <v-list-item-content>
                        <v-list-item-title><a class="mt-3 text-decoration-none" @click.prevent="goToAccount(userId)"><b>{{this.postFullname}}</b></a> {{this.postCaption}}</v-list-item-title>
                      </v-list-item-content>
                    </v-list-item>

                    <v-divider color="black"/>
                    <p class="mb-2"><b>Comments:</b></p> 
                    <template>
                      <v-list>
                        <v-list-tile-content
                        v-for="(comment, i) in postComment"
                        :key="i">
                          <!-- <div v-for="(comment, i) in postComment" :key="i"> -->
                            <p><a href="" class="mt-3 text-decoration-none" @click.prevent="goToAccount(comment.ID)"><b>{{comment.FullName}}</b></a> {{comment.Content}}</p>
                          <!-- </div> -->
                        </v-list-tile-content>
                      </v-list>
                    </template>
                  </v-list>
                </v-card-text>
                
                <v-footer class="ml-n2">
                  <div>
                    <v-row class="mx-n6 mt-1">
                      <v-col class="mr-n15" sm="2">
                        <v-btn icon class="mb-n7 mr-7">
                          <v-icon disabled >
                            mdi-cards-heart
                          </v-icon>
                        </v-btn>
                      </v-col>

                      <v-col class="ml-3"><p class="mt-4 ml-2">{{(this.postLike).length}} Likes</p></v-col>

                      <v-col align="right" v-if="this.visitorId == this.userId">
                        <v-btn tile color="success" @click="editPost">
                          <v-icon left>
                            mdi-pencil
                          </v-icon>
                          Edit
                        </v-btn>
                      </v-col>
                    </v-row>
                    
                    <v-row class="mx-n8 mt-n3 mb-n15">
                      <v-col>
                        <v-text-field
                          id="comment-pict"
                          v-model="newComment"
                          label="Post New Comment"
                          solo
                          style="width:400px"/>
                      </v-col>
                      <v-col class="mt-3">
                        <v-btn icon @click="addComment">
                          <v-icon>
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

        <v-row class="ma-3" v-if="postPhoto == ''">
          <v-col>
            <v-card-title class="ml-n3">
              <v-list-item-avatar color="grey darken-3">
                <img
                    lazy-src="../assets/default-profile.jpg"
                    :src="userAvatar"
                    alt="profilePicture"/>
              </v-list-item-avatar>
              <p class="pt-5">
                <a href=""
                  class="text-decoration-none"
                  style="color:#393E46"
                  @click="goToAccount(this.userId)"><b>{{this.postFullname}}</b></a>
              </p>
            </v-card-title>
          
            <v-card-text class="headline font-weight-normal" style="color:#393E46">
                {{this.postCaption}}
            </v-card-text>

            <v-divider class="my-5"/>

            <template>
              <v-list>
                <v-list-tile-content
                  v-for="(comment,i) in postComment"
                  :key ="i"
                  >
                  <!-- <div class="ml-5" v-for="(comment, i) in postComment" :key="i"> -->
                    <p><a href="" class="mt-3 text-decoration-none" @click.prevent="goToAccount(comment.ID)"><b>{{comment.FullName}}</b></a> {{comment.Content}}</p>
                  <!-- </div> -->
                </v-list-tile-content>
              </v-list>
            </template>

            <v-footer class="ml-n5" style="background-color:white">
              <v-row class="mx-4 mt-1">
                <v-col class="mr-n15" md="11">
                  <v-btn icon class="ml-6">
                    <v-icon disabled>
                      mdi-cards-heart
                    </v-icon>
                    <p class="mt-4 ml-2">{{(this.postLike).length}} Likes</p>
                  </v-btn>
                </v-col>
                
                <v-col class="mr-n8 mt-1" align="right" v-if="this.visitorId == this.userId">
                  <v-btn tile color="success" @click="editPost">
                    <v-icon left>
                      mdi-pencil
                    </v-icon>
                    Edit
                  </v-btn>
                </v-col>
              </v-row>

              <v-row class="mx-n8 mb-n10">
                <v-col class="ml-8" md="10">
                  <v-text-field
                    id="new-comment"
                    v-model="newComment"
                    label="Post New Comment"
                    solo
                    style="width:820px"/>
                </v-col>
                <v-spacer/>
                <v-col class="mt-3 mr-5">
                  <v-btn class="mr-n10" allign="right" icon @click="addComment">
                    <v-icon>
                      mdi-comment
                    </v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-footer>

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
</style>

<script>
import axios from "axios";

export default {
  mounted() {
    this.getVisitorId();
    this.getVisitorData();
    this.getPostId();
    this.getPostData();
  },
  data() {
    return {
      visitorId: "",
      visitorFullname: "",
      visitorAvatar: "",
      userId: "",
      postId: "",
      postLike: [],
      postComment: [],
      postCaption: "",
      postFullname: "",
      postPhoto: "",
      newComment: "",
      userAvatar: "",
      isLiked: false
    };
  },

  methods: {
    getVisitorId() {
      this.visitorId = this.$route.params.userId;
      // console.log(this.visitorId);
    },
    getVisitorData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.visitorId)
        .then((response) => {
          this.visitorFullname = response.data.data.fullname;
          this.visitorAvatar = response.data.data.profile;
        });
    },
    getPostId() {
      this.postId = this.$route.params.postId;
    },
    getPostData() {
      axios
        .get(`http://localhost:8081/getPost/` + this.postId)
        .then((response) => {
          this.postLike = response.data.data.like;
          this.postComment = response.data.data.comment;
          this.postCaption = response.data.data.caption;
          this.userId = response.data.data.user_id;
          this.postFullname = response.data.data.fullname;
          this.postPhoto = response.data.data.photo;

          axios
            .get(`http://localhost:8081/getUserData/` + this.userId)
            .then((response) => {
              this.userAvatar = response.data.data.profile;
            });
        });
    },
    goToAccount(userID) {
      this.$router.push({ path: "/" + userID + "/profile/" + this.visitorId });
    },
    goToProfile() {
      this.$router.push({
        path: "/" + this.visitorId + "/profile/" + this.visitorId,
      });
    },
    addComment() {
      var comment = document.getElementById("new-comment").value;

      var commentObj = {
        user_id: parseInt(this.visitorId),
        post_id: parseInt(this.postId),
        content: comment,
      };

      axios
        .post("http://localhost:8081/postComment", commentObj)
        .then((response) => {
          console.log(response);
          this.reloadPost();
        })
        .catch(function (error) {
          window.alert("Unable to Comment");
          console.log(error);
        });
    },
    reloadPost() {
      window.location.reload();
    },
    editPost() {
      this.$router.push({ path: "/updatePost/" + this.postId });
    },
    goHome() {
      this.$router.push({ path: "/home/" + this.visitorId });
    },
  },
};
</script>