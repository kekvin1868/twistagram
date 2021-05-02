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
            @click.prevent="goHome"
          />
        </a>
      </div>
      <v-spacer></v-spacer>
      <v-avatar size="50">
        <v-img :src="visitorAvatar"></v-img>
      </v-avatar>
      <p class="mt-3 ml-3 mr-13">
        <a
          href=""
          class="text-decoration-none"
          style="color: white"
          @click.prevent="goToProfile()"
          >{{ visitorFullname }}</a
        >
      </p>
    </v-app-bar>

    <v-main>
      <v-card
        class="mx-auto mt-8"
        max-width="1000"
        outlined
        rounded
        elevation="5"
      >
        <v-row class="ma-3" v-if="postPhoto != ''">
          <v-col>
            <v-img :src="postPhoto" aspect-ratio="1" max-height="800" />
          </v-col>

          <v-col>
            <div id="contents">
              <v-card elevation="0" height="450px">
                <v-card-text class="ml-n5">
                  <v-list>
                    <v-list-item>
                      <v-list-item-avatar
                        color="grey darken-3"
                        size="30"
                        class="mb-3 ml-n3"
                      >
                        <v-img :src="userAvatar"></v-img>
                      </v-list-item-avatar>

                      <v-list-item-content>
                        <v-list-item-title
                          ><a
                            class="mt-3 text-decoration-none"
                            @click.prevent="goToAccount(userId)"
                            ><b>{{ postFullname }}</b></a
                          >
                          {{ this.postCaption }}</v-list-item-title
                        >
                      </v-list-item-content>
                    </v-list-item>

                    <v-divider color="black" />
                    <p class="my-2"><b>Comments:</b></p>
                    <template>
                      <v-list>
                        <v-list-item-content
                          v-for="(comment, i) in postComment"
                          :key="i"
                        >
                          <v-row class="ml-1">
                            <p>
                              <a
                                href=""
                                class="mt-3 text-decoration-none"
                                @click.prevent="goToAccount(comment.ID)"
                                ><b>{{ comment.FullName }}</b></a
                              >
                              {{ comment.Content }}
                            </p>
                          </v-row>
                        </v-list-item-content>
                      </v-list>
                    </template>
                  </v-list>
                </v-card-text>

                <v-footer class="ml-n2" style="background-color: white">
                  <div>
                    <v-row class="mx-n6 mt-1">
                      <v-col class="mr-n15" sm="2">
                        <v-btn
                          icon
                          class="mb-n7 mr-7"
                          color="pink"
                          v-if="isLiked"
                          @click="likePost()"
                        >
                          <v-icon> mdi-cards-heart </v-icon>
                        </v-btn>
                        <v-btn
                          icon
                          class="mb-n7 mr-7"
                          v-if="!isLiked"
                          @click="likePost()">
                          <v-icon disabled> mdi-cards-heart </v-icon>
                        </v-btn>
                      </v-col>

                      <v-col class="ml-3 mr-n5">
                        <p class="mt-4 ml-2">
                          {{ postLike.length }} Likes
                        </p>
                      </v-col>

                      <v-col>
                        <v-btn
                          icon
                          class="mb-n6 ml-n15"
                          large>
                          <v-icon disabled> mdi-share </v-icon>
                        </v-btn>
                      </v-col>

                      <v-col align="right" v-if="visitorId == userId">
                        <v-btn tile color="success" @click="editPost">
                          <v-icon left> mdi-pencil </v-icon>
                          Edit
                        </v-btn>
                      </v-col>

                      <v-col align="right" class="mt-1 mr-n5" v-else>
                        <v-btn icon large v-if="!isBookmarked" @click="bookmarkPost">
                          <v-icon left> mdi-bookmark-outline </v-icon>
                        </v-btn>
                        <v-btn icon large v-if="isBookmarked" @click="bookmarkPost">
                          <v-icon left> mdi-bookmark </v-icon>
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
                          style="width: 400px"/>
                      </v-col>
                      <v-col class="mt-3">
                        <v-btn icon @click="addComment()">
                          <v-icon> mdi-comment </v-icon>
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
                  alt="profilePicture"
                />
              </v-list-item-avatar>
              <p class="pt-5">
                <a
                  href=""
                  class="text-decoration-none"
                  style="color: #393e46"
                  @click="goToAccount(userId)"
                  ><b>{{ postFullname }}</b></a>
              </p>
            </v-card-title>

            <v-card-text
              class="headline font-weight-normal"
              style="color: #393e46">
              {{ postCaption }}
            </v-card-text>

            <v-divider class="my-5" color="black" />

            <template>
              <v-list>
                <v-list-item-content
                  v-for="(comment, i) in postComment"
                  :key="i">
                  <v-row class="ml-5">
                    <p>
                      <a
                        href=""
                        class="mt-3 text-decoration-none"
                        @click.prevent="goToAccount(comment.ID)"
                        ><b>{{ comment.FullName }}</b></a>
                      {{ comment.Content }}
                    </p>
                  </v-row>
                </v-list-item-content>
              </v-list>
            </template>

            <v-footer class="ml-n5" style="background-color: white">
              <v-row class="mx-4 mt-1">
                <v-col class="mr-n15" md="11">
                  <v-btn
                    icon
                    class="ml-6"
                    color="pink"
                    v-if="isLiked"
                    @click="likePost()">
                    <v-icon> mdi-cards-heart </v-icon>
                    <p class="mt-4 ml-2">{{ postLike.length }} Likes</p>
                  </v-btn>

                  <v-btn icon class="ml-6" v-if="!isLiked" @click="likePost()">
                    <v-icon disabled> mdi-cards-heart </v-icon>
                    <p class="mt-4 ml-2">
                      <b>{{ postLike.length }} Likes</b>
                    </p>
                  </v-btn>
                  <v-btn
                    icon
                    class="ml-15"
                    large>
                    <v-icon disabled> mdi-share </v-icon>
                  </v-btn>
                </v-col>
                
                <v-col
                  class="mr-n8 mt-1"
                  align="right"
                  v-if="visitorId == userId">
                  <v-btn tile color="success" @click="editPost">
                    <v-icon left> mdi-pencil </v-icon>
                    Edit
                  </v-btn>
                </v-col>
                <v-col 
                  class="mr-n8 mt-3"
                  align="right"
                  v-if="visitorId != userId">
                  <v-btn icon large v-if="!isBookmarked" @click="bookmarkPost">
                    <v-icon left> mdi-bookmark-outline </v-icon>
                  </v-btn>
                  <v-btn icon large v-if="isBookmarked" @click="bookmarkPost">
                    <v-icon left> mdi-bookmark </v-icon>
                  </v-btn>
                </v-col>
              </v-row>

              <v-row class="ml-n5 mb-n10">
                <v-col class="ml-8" md="11">
                  <v-text-field
                    id="new-comment"
                    v-model="newComment"
                    label="Post New Comment"
                    solo
                    style="width: 840px"/>
                </v-col>
                <v-spacer />
                <v-col class="mt-3">
                  <v-btn
                    class="ml-n5"
                    allign="right"
                    icon
                    large
                    @click="addComment()">
                    <v-icon> mdi-comment </v-icon>
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
    this.getBookmarkStatus();
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
      isLiked: false,
      likeId: null,
      idBookmark: "",
      isBookmarked: false,
      userBookmark: []
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
          for (let index = 0; index < this.postLike.length; index++) {
            if (this.postLike[index].user_id == this.visitorId) {
              this.isLiked = true;
              this.likeId = this.postLike[index].id;
              break;
            }
          }

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
    getBookmarkStatus() {
      axios
        .get(`http://localhost:8081/getBookmark/` + this.visitorId)
        .then((response) => {
          this.userBookmark = response.data.data;
          var bookmark = null;

          bookmark = this.userBookmark.find(
              (book) => book.post_id == this.postId
          );

          if(bookmark != null) {
              this.isBookmarked = true;
              this.idBookmark = bookmark.ID;
          } else {
            this.isBookmarked = false;
          }
          console.log(this.idBookmark);
        });
    },
    addComment() {
      var comment = this.newComment;

      if (this.newComment != "") {
        var commentObj = {
          user_id: parseInt(this.visitorId),
          post_id: parseInt(this.postId),
          content: comment,
        };

        axios
          .post("http://localhost:8081/postComment", commentObj)
          .then((response) => {
            console.log(response);
            this.newComment = "";
            this.getPostData();
          })
          .catch(function (error) {
            window.alert("Unable to Comment");
            console.log(error);
          });
      } else {
        window.alert("Comment cannot be empty");
      }
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
    async likePost() {
      if (!this.isLiked) {
        axios
          .post("http://localhost:8081/postLike", {
            user_id: parseInt(this.visitorId),
            post_id: parseInt(this.postId),
          })
          .then((response) => {
            this.getPostData();
            this.likeId = response.data.data.id;
          });
      } else {
        await axios
          .delete("http://localhost:8081/deleteLike/" + this.likeId)
          .then(() => {
            this.getPostData();
            this.isLiked = false;
          });
      }
    },
    async bookmarkPost() {
      if (!this.isBookmarked) {
        axios
          .post(`http://localhost:8081/postBookmark`, {
              post_id: parseInt(this.postId),
              user_id: parseInt(this.userId)
          })
          .then((response) => {
              this.idBookmark = response.data.data.id;
              this.getBookmarkStatus();
          })
          .catch(function (error) {
              window.alert("Bookmark Post Failed");
              console.log(error);
          });
      } else {
        await axios
            .delete(`http://localhost:8081/deleteBookmark/` + this.idBookmark)
            .then((response) => {
                console.log(response);
                this.isBookmarked = false;
                this.getBookmarkStatus();
            })
            .catch(function (error) {
                window.alert("Bookmark unPost Failed");
                console.log(error);
            });
      }
    },
  },
};
</script>