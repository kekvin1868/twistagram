
<!-- buat yang edit home: line 53 & 239 -->

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

      <v-icon> mdi-magnify </v-icon>

      <v-icon> mdi-search </v-icon>
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
                      <v-card flat>
                        <div v-if="postData != null">
                          <div
                            v-for="(data, i) in postData"
                            :key="i"
                            id="postDiv"
                          >
                            <v-card
                              class="mx-auto my-10 rounded"
                              max-width="800"
                              v-if="data.photo != ''"
                            >
                              <v-card-title>
                                <v-avatar size="70">
                                  <v-img :src="userData.profile"></v-img>
                                </v-avatar>

                                <p class="ml-5">{{ data.fullname }}</p>
                                <v-spacer></v-spacer>

                                <v-card-actions>
                                  <v-btn icon
                                  @click="bookmark(data.id)">
                                    <v-icon color="black" x-large>
                                      mdi-bookmark
                                    </v-icon>
                                  </v-btn>
                                  <v-btn icon>
                                    <v-icon x-large color="black">
                                      mdi-dots-vertical
                                    </v-icon>
                                  </v-btn>
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
                                <!-- show data comment -->
                                <div id="layout-comment">
                                  <h4>Comments</h4>
                                  <div v-if="data.comment.length > 0">
                                    <p>
                                      <b> {{ data.comment[0].FullName }} </b>
                                      {{ data.comment[0].Content }}
                                    </p>
                                    <a @click="goToShowPost(data.id)"
                                      >Show More ...</a
                                    >
                                  </div>
                                </div>
                                <v-divider color="black"></v-divider>
                                <v-row no-gutters>
                                  <v-col sm="1">
                                    <v-btn
                                      class="mt-1 ml-2"
                                      icon
                                      x-large
                                      color="pink"
                                    >
                                      <v-icon>mdi-heart</v-icon>
                                    </v-btn>
                                  </v-col>

                                  <v-col sm="2" md="8">
                                    <v-text-field
                                      required
                                      single-line
                                      display="flex"
                                      rounded
                                      background-color="grey"
                                      color="black"
                                      label="Insert your comment here..."
                                      v-model="caption"
                                    />
                                  </v-col>
                                  <v-col class="mt-3 ml-5">
                                    <v-btn
                                      @click="comment(data.id)"
                                      color="primary"
                                    >
                                      <v-icon> mdi-send </v-icon>
                                    </v-btn>
                                  </v-col>
                                </v-row>
                              </v-card-text>
                            </v-card>

                            <!-- WITHOUT PHOTO -->
                            <v-card
                              class="mt-5 rounded-xl"
                              max-width="800"
                              v-if="data.photo == ''"
                            >
                              <v-card-title>
                                <v-avatar class="mt-2 ml-2" size="70">
                                  <v-img :src="userData.profile"> </v-img>
                                </v-avatar>
                                <p class="ml-3">{{ data.fullname }}</p>
                                <v-spacer></v-spacer>
                                <v-btn icon>
                                  <v-icon large color="black">
                                    mdi-dots-vertical
                                  </v-icon>
                                </v-btn>
                              </v-card-title>

                              <v-card-text>
                                <p class="pt-3 pl-5">
                                  {{ data.caption }}
                                </p>
                              </v-card-text>

                              <v-divider class="ml-8" width="650"></v-divider>
                              <v-card-actions class="ml-4">
                                <v-btn icon>
                                  <v-icon medium> mdi-comment </v-icon>
                                </v-btn>
                                <v-btn icon>
                                  <v-icon color="red"> mdi-heart </v-icon>
                                </v-btn>
                              </v-card-actions>
                            </v-card>
                          </div>
                        </div>
                      </v-card>
                    </v-tab-item>

                    <!-- TAB FOR SHARE -->
                    <v-tab-item :value="'tab-' + 2">
                      <v-card flat>
                        <v-card-text>
                          Lorem ipsum dolor sit amet, consectetur adipiscing
                          elit, sed do eiusmod tempor incididunt ut labore et
                          dolore magna aliqua. Ut enim ad minim veniam, quis
                          nostrud exercitation ullamco laboris nisi ut aliquip
                          ex ea commodo consequat.
                        </v-card-text>
                      </v-card>
                    </v-tab-item>
                  </v-tabs-items>
                </v-card>

                <div v-if="this.postData.length == 0">
                  <v-card>
                    <!-- tampilin sesuatu jika ga ada post -->
                    <h1>LOL</h1>
                  </v-card>
                </div>

                <br /><br />
              </v-col>

              <!-- CARD FOR TRENDING -->
              <v-col sm="2">
                <v-card class="rounded-xl" color="#222831">
                  <v-card-title>
                    <v-layout align-center justify-center>
                      <h3>Suggestion</h3>
                    </v-layout>
                  </v-card-title>
                  <v-divider color="white"> </v-divider>

                  <v-row v-for="x in suggestions" :key="x.image">
                    <v-card-text>
                      <v-col>
                        <v-avatar size="50">
                          <img :src="x.image" />
                        </v-avatar>
                        {{ x.text }}
                      </v-col>
                    </v-card-text>
                  </v-row>
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
      feedsIds: [],
      followingObj: "",
      followerObj: "",
      postCaption: "",
      imgData: null,
      withPhoto: true,
      caption: "",
      userId: "",
      userData: [],
      suggestions: [
        {
          image: "/images/default-profile.jpg",
          text: "Budi",
        },
        {
          image: "../assets/logo.png",
          text: "lol",
        },
      ],
      postData: [],
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

    goToShowPost(postID) {
      this.$router.push({
        path: "/post/" + postID + "/" + this.userId,
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
              this.postData.push(response.data.data);
            });
        }
      }
    },

    bookmark(postId) {
      
      axios
        .post("http://localhost:8081/postBookmark", {
          post_id: postId,
          user_id: parseInt(this.userId),
        })
        .then((response) => {
          if(response.data.message == "" ){
            alert("Saved to bookmark");
          } else if (response.data.message == "bookmark exist") {
            alert("This post is already bookmarked");
          }
        });
    },
  },

  mounted() {
    this.getUserId();
    this.getFollowingAndFollowerData();
    this.getUserData();
    this.loadFeeds();
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