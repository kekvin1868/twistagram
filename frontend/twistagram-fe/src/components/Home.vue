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
          width="150"
        />
      </div>

      <v-spacer></v-spacer>

      <v-btn @click="goToSearch()" icon>
        <v-icon> mdi-magnify </v-icon>
      </v-btn>
      <v-avatar size="50">
        <v-img :src="userData.profile"></v-img>
      </v-avatar>
      <p class="mt-3 ml-3 mr-13">{{ userData.fullname }}</p>
    </v-app-bar>

    <v-main>
      <v-app id="app">
        <v-main class="grey lighten-3">
          <v-container>
            <!-- CARD FOR PROFILE -->
            <v-row id="margin">
              <v-col cols="2">
                <v-card class="rounded-xl" color="#222831">
                  <v-card-title>
                    <v-layout align-center justify-center>
                      <v-avatar size="95">
                        <v-img :src="userData.profile"> </v-img>
                      </v-avatar>
                    </v-layout>
                  </v-card-title>
                  <v-card-text>
                    <v-layout align-center justify-center>
                      <h3>
                        {{ userData.fullname }}
                      </h3>
                    </v-layout>
                    <v-divider class="my-3" color="white"> </v-divider>
                    <h3 class="text-center">
                      Follower {{ this.followerObj.Count }}
                      <v-divider class="my-3" color="white"> </v-divider>
                      Following {{ this.followingObj.Count }}
                      <v-divider class="my-3" color="white"> </v-divider>
                      <v-btn text color="white" @click="goToProfile">
                        View Profile
                      </v-btn>
                    </h3>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- CARD FOR UPLOAD CONTENT & POST-->
              <v-col sm="6">
                <v-card
                  class="mx-auto rounded-xl"
                  max-width="800"
                  color="#222831"
                >
                  <v-card-text>
                    <v-row no-gutters>
                      <v-col md="3">
                        <v-card-title>
                          <v-avatar size="75">
                            <v-img :src="userData.profile"> </v-img>
                          </v-avatar>
                        </v-card-title>
                      </v-col>

                      <v-col md="8">
                        <v-textarea
                          placeholder="Type something here"
                          background-color="white"
                          rounded
                          v-model="postCaption"
                        />
                      </v-col>
                    </v-row>

                    <v-divider color="white"> </v-divider>
                    <v-card-actions>
                      <v-file-input
                        label="File input"
                        prepend-icon="mdi-image"
                        v-model="imgData"
                      />
                      <v-spacer></v-spacer>
                      <v-btn color="primary" @click="post"> post </v-btn>
                    </v-card-actions>
                  </v-card-text>
                </v-card>

                <!-- TABS -->
                <v-card>
                  <v-tabs
                    class="mt-5 rounded"
                    v-model="tab"
                    background-color="#222831"
                    centered
                    dark
                    icons-and-text
                  >
                    <v-tabs-slider></v-tabs-slider>

                    <v-tab href="#tab-1">
                      Feeds
                      <v-icon>mdi-home</v-icon>
                    </v-tab>

                    <v-tab href="#tab-2">
                      Share
                      <v-icon>mdi-share-variant</v-icon>
                    </v-tab>
                  </v-tabs>

                  <!-- TAB FOR FEEDS -->
                  <v-tabs-items v-model="tab">
                    <v-tab-item :value="'tab-' + 1">
                      <div v-if="this.postData.length == 0">
                        <v-card>
                          <v-layout justify-center>
                            <v-img
                              max-width="500"
                              max-height="400"
                              src="../assets/frown4.png"
                            >
                            </v-img>
                          </v-layout>
                          <h1 align="center">No post yet</h1>
                        </v-card>
                      </div>

                      <div else>
                        <v-card flat>
                          <div v-if="postData != null">
                            <div
                              v-for="(data, i) in postData"
                              :key="i"
                              id="postDiv"
                            >
                              <v-card
                                class="mx-auto my-10 rounded-xl"
                                max-width="800"
                                v-if="data.photo != ''"
                              >
                                <v-card-title>
                                  <v-avatar size="70">
                                    <v-img :src="data.profile"></v-img>
                                  </v-avatar>

                                  <p class="ml-5">{{ data.fullname }}</p>
                                  <v-spacer></v-spacer>

                                  <v-card-actions>
                                    <v-btn
                                      icon
                                      @click="
                                        bookmark(data.id, data.isBookmarked)
                                      "
                                      v-if="data.isBookmarked == null"
                                    >
                                      <v-icon color="black" x-large>
                                        mdi-bookmark
                                      </v-icon>
                                    </v-btn>
                                    <v-btn
                                      icon
                                      @click="
                                        bookmark(data.id, data.isBookmarked)
                                      "
                                      v-else
                                    >
                                      <v-icon color="blue" x-large>
                                        mdi-bookmark
                                      </v-icon>
                                    </v-btn>
                                    <v-menu bottom left>
                                      <template
                                        v-slot:activator="{ on, attrs }"
                                      >
                                        <v-btn
                                          dark
                                          icon
                                          v-bind="attrs"
                                          v-on="on"
                                        >
                                          <v-icon x-large color="black"
                                            >mdi-dots-vertical</v-icon
                                          >
                                        </v-btn>
                                      </template>

                                      <v-list>
                                        <v-list-item
                                          v-if="data.user_id != userId"
                                        >
                                          <v-list-item-title
                                            ><button
                                              @click="reportPost(data.id)"
                                            >
                                              <v-icon color="red"
                                                >mdi-alert</v-icon
                                              >
                                              Report
                                            </button></v-list-item-title
                                          >
                                        </v-list-item>
                                        <v-list-item>
                                          <v-list-item-title
                                            ><button
                                              @click="goToShowPost(data.id)"
                                            >
                                              <v-icon color="blue"
                                                >mdi-page-next</v-icon
                                              >
                                              Go to Original Post
                                            </button></v-list-item-title
                                          >
                                        </v-list-item>
                                      </v-list>
                                    </v-menu>
                                  </v-card-actions>
                                </v-card-title>
                                <v-img :src="data.photo" />
                                <v-card-text>
                                  <v-row no-gutters>
                                    <v-col sm="2">
                                      <h2 class="ml-5">{{ data.fullname }}</h2>
                                    </v-col>
                                    <v-col>
                                      <p>
                                        {{ data.caption }}
                                      </p>
                                    </v-col>
                                  </v-row>

                                  <br />

                                  <v-row>
                                    <v-icon medium class="ml-5 mr-1">
                                      mdi-comment
                                    </v-icon>
                                    <v-col class="ml-n3" sm="1">
                                      {{ data.comment.length }}
                                    </v-col>

                                    <v-icon color="red"> mdi-heart </v-icon>
                                    <v-col class="ml-n2" sm="1">
                                      {{ data.like.length }}
                                    </v-col>
                                    <v-col sm="3">
                                      <a @click="goToShowPost(data.id)">
                                        Show more ...
                                      </a>
                                    </v-col>
                                  </v-row>
                                </v-card-text>
                              </v-card>

                              <!-- FEEDS WITHOUT PHOTO -->
                              <v-card
                                class="mt-5 rounded-xl"
                                max-width="800"
                                v-if="data.photo == ''"
                              >
                                <v-card-title>
                                  <v-avatar class="mt-2 ml-2" size="70">
                                    <v-img :src="data.profile"> </v-img>
                                  </v-avatar>
                                  <p class="ml-3">{{ data.fullname }}</p>
                                  <v-spacer></v-spacer>
                                  <v-menu bottom left>
                                    <template v-slot:activator="{ on, attrs }">
                                      <v-btn dark icon v-bind="attrs" v-on="on">
                                        <v-icon x-large color="black"
                                          >mdi-dots-vertical</v-icon
                                        >
                                      </v-btn>
                                    </template>

                                    <v-list>
                                      <v-list-item
                                        v-if="data.user_id != userId"
                                      >
                                        <v-list-item-title
                                          ><button @click="reportPost(data.id)">
                                            <v-icon color="red"
                                              >mdi-alert</v-icon
                                            >
                                            Report
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                      <v-list-item>
                                        <v-list-item-title
                                          ><button
                                            @click="goToShowPost(data.id)"
                                          >
                                            <v-icon color="blue"
                                              >mdi-page-next</v-icon
                                            >
                                            Go to Original Post
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                    </v-list>
                                  </v-menu>
                                </v-card-title>

                                <v-card-text>
                                  <p class="pt-3 pl-5">
                                    {{ data.caption }}
                                  </p>
                                </v-card-text>

                                <v-card-actions class="ml-4">
                                  <v-row>
                                    <v-icon medium class="ml-5 mr-1">
                                      mdi-comment
                                    </v-icon>
                                    <v-col class="ml-n3" sm="1">
                                      {{ data.comment.length }}
                                    </v-col>

                                    <v-icon color="red"> mdi-heart </v-icon>
                                    <v-col class="ml-n2" sm="1">
                                      {{ data.like.length }}
                                    </v-col>

                                    <v-col sm="3">
                                      <a @click="goToShowPost(data.id)">
                                        Show more ...
                                      </a>
                                    </v-col>
                                  </v-row>
                                </v-card-actions>
                              </v-card>
                            </div>
                          </div>
                        </v-card>
                      </div>
                    </v-tab-item>

                    <!-- TAB FOR SHARE -->
                    <v-tab-item :value="'tab-' + 2">
                      <div v-if="this.shareData.length == 0">
                        <!-- WITH PHOTO -->
                        <v-card>
                          <v-layout justify-center>
                            <v-img
                              max-width="500"
                              max-height="400"
                              src="../assets/frown4.png"
                            >
                            </v-img>
                          </v-layout>
                          <h1 align="center">No one shared post yet</h1>
                        </v-card>
                      </div>

                      <div v-else>
                        <v-card flat>
                          <div v-if="postData != null">
                            <div
                              v-for="(data, i) in shareData"
                              :key="i"
                              id="shareDiv"
                            >
                              <v-card
                                class="mx-auto my-10 rounded-xl"
                                max-width="800"
                                v-if="data.postData.photo != ''"
                              >
                                <v-card-title>
                                  <v-avatar size="70">
                                    <v-img :src="data.postData.profile"></v-img>
                                  </v-avatar>

                                  <p class="ml-5">
                                    {{ data.postData.fullname }}
                                  </p>
                                  <v-spacer></v-spacer>

                                  <v-menu bottom left>
                                    <template v-slot:activator="{ on, attrs }">
                                      <v-btn dark icon v-bind="attrs" v-on="on">
                                        <v-icon x-large color="black"
                                          >mdi-dots-vertical</v-icon
                                        >
                                      </v-btn>
                                    </template>

                                    <v-list>
                                      <v-list-item
                                        v-if="data.postData.user_id != userId"
                                      >
                                        <v-list-item-title
                                          ><button
                                            @click="
                                              reportPost(data.postData.id)
                                            "
                                          >
                                            <v-icon color="red"
                                              >mdi-alert</v-icon
                                            >
                                            Report
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                      <v-list-item>
                                        <v-list-item-title
                                          ><button
                                            @click="
                                              goToShowPost(data.postData.id)
                                            "
                                          >
                                            <v-icon color="blue"
                                              >mdi-page-next</v-icon
                                            >
                                            Go to Original Post
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                    </v-list>
                                  </v-menu>
                                </v-card-title>
                                <v-img :src="data.postData.photo" />
                                <v-card-text>
                                  <v-row no-gutters>
                                    <v-col sm="2">
                                      <h2 class="ml-5">
                                        {{ data.userData.fullname }}
                                      </h2>
                                    </v-col>
                                    <v-col>
                                      <p>
                                        {{ data.postData.caption }}
                                      </p>
                                    </v-col>
                                  </v-row>
                                </v-card-text>
                                <v-card-actions>
                                  <v-list-item class="grow">
                                    <v-list-item-avatar color="grey darken-3">
                                      <v-img
                                        class="elevation-6"
                                        alt=""
                                        :src="data.userData.profile"
                                      ></v-img>
                                    </v-list-item-avatar>

                                    <v-list-item-content>
                                      <v-list-item-subtitle
                                        >Shared by</v-list-item-subtitle
                                      >
                                      <v-list-item-title>{{
                                        data.userData.fullname
                                      }}</v-list-item-title>
                                    </v-list-item-content>

                                    <v-row align="center" justify="end">
                                      <v-icon class="mr-1" color="red">
                                        mdi-heart
                                      </v-icon>
                                      <span class="subheading mr-2">{{
                                        data.postData.like.length
                                      }}</span>
                                      <span class="mr-1">·</span>
                                      <v-icon class="mr-1">
                                        mdi-comment
                                      </v-icon>
                                      <span class="subheading">{{
                                        data.postData.comment.length
                                      }}</span>
                                    </v-row>
                                  </v-list-item>
                                </v-card-actions>
                              </v-card>

                              <!-- SHARE WITHOUT PHOTO -->
                              <v-card
                                class="mt-5 rounded-xl"
                                max-width="800"
                                v-if="data.postData.photo == ''"
                              >
                                <v-card-title>
                                  <v-avatar class="mt-2 ml-2" size="70">
                                    <v-img :src="data.postData.profile">
                                    </v-img>
                                  </v-avatar>
                                  <p class="ml-3">
                                    {{ data.postData.fullname }}
                                  </p>
                                  <v-spacer></v-spacer>
                                  <v-menu bottom left>
                                    <template v-slot:activator="{ on, attrs }">
                                      <v-btn dark icon v-bind="attrs" v-on="on">
                                        <v-icon x-large color="black"
                                          >mdi-dots-vertical</v-icon
                                        >
                                      </v-btn>
                                    </template>

                                    <v-list>
                                      <v-list-item
                                        v-if="data.postData.user_id != userId"
                                      >
                                        <v-list-item-title
                                          ><button
                                            @click="
                                              reportPost(data.postData.id)
                                            "
                                          >
                                            <v-icon color="red"
                                              >mdi-alert</v-icon
                                            >
                                            Report
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                      <v-list-item>
                                        <v-list-item-title
                                          ><button
                                            @click="
                                              goToShowPost(data.postData.id)
                                            "
                                          >
                                            <v-icon color="blue"
                                              >mdi-page-next</v-icon
                                            >
                                            Go to Original Post
                                          </button></v-list-item-title
                                        >
                                      </v-list-item>
                                    </v-list>
                                  </v-menu>
                                </v-card-title>

                                <v-card-text>
                                  <p class="pt-3 pl-5">
                                    {{ data.postData.caption }}
                                  </p>
                                </v-card-text>

                                <v-card-actions>
                                  <v-list-item class="grow">
                                    <v-list-item-avatar color="grey darken-3">
                                      <v-img
                                        class="elevation-6"
                                        alt=""
                                        :src="data.userData.profile"
                                      ></v-img>
                                    </v-list-item-avatar>

                                    <v-list-item-content>
                                      <v-list-item-subtitle
                                        >Shared by</v-list-item-subtitle
                                      >
                                      <v-list-item-title>{{
                                        data.userData.fullname
                                      }}</v-list-item-title>
                                    </v-list-item-content>

                                    <v-row align="center" justify="end">
                                      <v-icon class="mr-1" color="red">
                                        mdi-heart
                                      </v-icon>
                                      <span class="subheading mr-2">{{
                                        data.postData.like.length
                                      }}</span>
                                      <span class="mr-1">·</span>
                                      <v-icon class="mr-1">
                                        mdi-comment
                                      </v-icon>
                                      <span class="subheading">{{
                                        data.postData.comment.length
                                      }}</span>
                                    </v-row>
                                  </v-list-item>
                                </v-card-actions>
                              </v-card>
                            </div>
                          </div>
                        </v-card>
                      </div>
                    </v-tab-item>
                  </v-tabs-items>
                </v-card>

                <br /><br />
              </v-col>

              <!-- CARD FOR SUGGESTIONS -->
              <v-col sm="2">
                <v-card class="rounded-xl" color="#222831" dark>
                  <v-card-title>
                    <v-layout align-center justify-center>
                      <h3>Suggestion</h3>
                    </v-layout>
                  </v-card-title>
                  <v-divider color="white"> </v-divider>

                  <div v-if="suggestions.length > 0">
                    <v-row v-for="x in suggestions" :key="x.ID">
                      <v-card-text>
                        <v-list-item>
                          <v-list-item-avatar>
                            <v-img :src="x.Profile"></v-img>
                          </v-list-item-avatar>

                          <v-list-item-content>
                            <v-list-item-title>
                              <p
                                class="mt-3"
                                @click="goToAccount(x.ID)"
                                style="cursor: pointer"
                              >
                                {{ x.FullName }}
                              </p>
                            </v-list-item-title>
                          </v-list-item-content>
                        </v-list-item>
                      </v-card-text>
                    </v-row>
                  </div>

                  <div class="mt-5" v-else>
                    <v-card-text>
                      <v-card-subtitle color="white"
                        >No Suggestion</v-card-subtitle
                      >
                    </v-card-text>
                  </div>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-main>
      </v-app>
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      tab: null,
      counts: "",
      postData: [],
      feedsIds: [],
      followingObj: "",
      followerObj: "",
      postCaption: "",
      imgData: null,
      withPhoto: true,
      caption: "",
      userId: "",
      userData: [],
      suggestions: [],
      shareData: [],
    };
  },
  methods: {
    async post() {
      if (this.postCaption == "") {
        window.alert("Caption cannot be empty");
      } else {
        if (this.imgData == null) {
          axios
            .post("http://localhost:8081/post", {
              caption: this.postCaption,
              photo: null,
              user_id: parseInt(this.userId),
            })
            .then(() => {
              this.postCaption = "";
              this.reloadFeeds();
            });
        } else {
          let img = null;
          const reader = new FileReader();
          reader.readAsDataURL(this.imgData);
          reader.onload = () => {
            img = reader.result;
            axios
              .post("http://localhost:8081/post", {
                caption: this.postCaption,
                photo: img,
                user_id: parseInt(this.userId),
              })
              .then(() => {
                this.postCaption = "";
                this.imgData = null;
                this.reloadFeeds();
              });
          };
        }
      }
    },

    goToAccount(userID) {
      this.$router.push({ path: "/" + userID + "/profile/" + this.userId });
    },
    goToShowPost(postID) {
      this.$router.push({
        path: "/post/" + postID + "/" + this.userId,
      });
    },

    goToSearch() {
      this.$router.push({
        path: "/home/" + this.userId + "/search",
      });
    },

    goToProfile() {
      this.$router.push({
        path: "/" + this.userId + "/profile/" + this.userId,
      });
    },
    reloadFeeds() {
      this.postData = [];
      this.loadFeeds();
      document.getElementById("postDiv").innerHTML = document.getElementById(
        "postDiv"
      ).innerHTML;
    },
    getUserId() {
      this.userId = this.$route.params.userId;
    },
    getUserData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.userId)
        .then((response) => {
          this.userData = response.data.data;
        });
    },
    comment(id) {
      let param = {
        user_id: parseInt(this.userId),
        post_id: id,
        content: this.caption,
      };
      if (this.caption == "") {
        alert("Field Comment masih Kosong!");
      } else {
        axios.post("http://localhost:8081/postComment", param).then(() => {
          this.caption = "";
          this.reloadFeeds();
        });
      }
    },
    getFollowingAndFollowerData() {
      axios
        .get(`http://localhost:8081/getFollowing/` + this.userId)
        .then((response) => {
          this.followingObj = response.data.data;
        });
      axios
        .get(`http://localhost:8081/getFollowers/` + this.userId)
        .then((response) => {
          this.followerObj = response.data.data;
        });
    },
    async loadFeeds() {
      await axios
        .get(`http://localhost:8081/loadFeeds/` + this.userId)
        .then((response) => {
          this.feedsIds = response.data.data;
        });
      if (this.feedsIds != null) {
        for (let index = 0; index < this.feedsIds.length; index++) {
          axios
            .get("http://localhost:8081/getPost/" + this.feedsIds[index].ID)
            .then((response) => {
              var postId = response.data.data.id;
              var caption = response.data.data.caption;
              var fullname = response.data.data.fullname;
              var user_id = response.data.data.user_id;
              var like = response.data.data.like;
              var comment = response.data.data.comment;
              var photo = response.data.data.photo;
              var isBookmarked = null;
              var profile = response.data.data.profile;

              axios
                .get("http://localhost:8081/getBookmark/" + this.userId)
                .then((response) => {
                  if (response.data.data.length > 0) {
                    for (
                      let index = 0;
                      index < response.data.data.length;
                      index++
                    ) {
                      if (response.data.data[index].post_id == postId) {
                        isBookmarked = response.data.data[index].ID;
                      }
                    }
                  }
                  var postDataObj = {
                    id: postId,
                    caption: caption,
                    fullname: fullname,
                    user_id: user_id,
                    like: like,
                    comment: comment,
                    photo: photo,
                    isBookmarked: isBookmarked,
                    profile: profile
                  };
                  this.postData.push(postDataObj);
                });
            });
        }
      }
    },
    bookmark(postId, bookmarkId) {
      axios
        .post("http://localhost:8081/postBookmark", {
          post_id: postId,
          user_id: parseInt(this.userId),
        })
        .then((response) => {
          if (response.data.message == "") {
            this.reloadFeeds();
          } else if (response.data.message == "bookmark exist") {
            axios
              .delete("http://localhost:8081/deleteBookmark/" + bookmarkId)
              .then(() => {
                this.reloadFeeds();
              });
          }
        });
    },
    getShareData() {
      axios
        .get("http://localhost:8081/loadShare/" + this.userId)
        .then((response) => {
          var responseShare = response.data.data;
          if (responseShare != null) {
            for (let index = 0; index < responseShare.length; index++) {
              axios
                .get(
                  `http://localhost:8081/getUserData/` +
                    responseShare[index].UserID
                )
                .then((responseUser) => {
                  var userData = responseUser.data.data;

                  axios
                    .get(
                      "http://localhost:8081/getPost/" +
                        responseShare[index].PostID
                    )
                    .then((responsePost) => {
                      var postData = responsePost.data.data;

                      var shareDataObj = {
                        id: responseShare[index].ID,
                        userData: userData,
                        postData: postData,
                      };
                      this.shareData.push(shareDataObj);
                    });
                });
            }
          }
        });
    },
    reportPost(postId) {
      axios
        .post("http://localhost:8081/report", {
          post_id: parseInt(postId),
          user_id: parseInt(this.userId),
        })
        .then(() => {
          window.alert("reported post-id: " + postId);
        });
    },
    getSugestion() {
      axios
        .get("http://localhost:8081/getSuggestion/" + this.userId)
        .then((response) => {
          this.suggestions = response.data.data;
        });
    },
  },
  mounted() {
    this.getUserId();
    this.getFollowingAndFollowerData();
    this.getUserData();
    this.loadFeeds();
    this.getShareData();
    this.getSugestion();
  },
};
</script>

<style>
.h1 {
  padding-left: 25px;
  padding-top: 5px;
  padding-bottom: 5px;
  color: white;
}
.p {
  padding-left: 20px;
  padding-top: 5px;
  color: white;
}
#margin {
  margin-left: 280px;
}
#layout-comment {
  margin-left: 20px;
}
h3 {
  color: white;
}
</style>