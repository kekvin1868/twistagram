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
      <v-spacer> </v-spacer>
      <v-avatar size="50">
        <v-img :src="this.visitorAvatar"></v-img>
      </v-avatar>
      <p class="mt-3 ml-3 mr-13">
        <a
          href=""
          class="text-decoration-none"
          style="color: white"
          @click.prevent="goToProfile()"
          >{{ this.visitorFullname }}</a
        >
      </p>
    </v-app-bar>

    <v-main>
      <div class="userProfile">
        <v-card class="mx-auto mt-8 mb-5 py-5" max-width="900" elevate="0">
          <v-list-item>
            <v-avatar class="my-auto ml-7" size="140">
              <img
                lazy-src="../assets/default-profile.jpg"
                :src="this.userAvatar"
                alt="profilePicture"
              />
            </v-avatar>

            <v-list-item-content class="ml-10 my-auto">
              <v-row>
                <v-col cols="auto">
                  <p class="mt-1 display-1 text--primary">
                    {{ this.userFullName }}
                  </p>
                </v-col>
                <v-col cols="auto" v-if="this.visitorId == this.userId">
                  <v-btn
                    class="mt-4 font-weight-black"
                    color="#FFD369"
                    elevation="2"
                    x-small
                    @click="editProfile"
                    >Edit Profile</v-btn
                  >
                </v-col>
                <v-col cols="auto" v-if="this.visitorId == this.userId">
                  <v-btn
                    class="ml-n3 mt-4 font-weight-black"
                    color="#4CAF50"
                    elevation="2"
                    x-small
                    width="95px"
                    @click="goToBookmark"
                    >Bookmarked</v-btn
                  >
                </v-col>
                <v-col cols="auto" v-if="this.visitorId == this.userId">
                  <v-btn
                    class="ml-n3 mt-4 font-weight-black"
                    color="#F85151"
                    elevation="2"
                    x-small
                    width="95px"
                    @click="logout"
                    >Logout</v-btn
                  >
                </v-col>
                <v-col
                  cols="auto"
                  v-if="
                    this.visitorId != this.userId && this.isFollowed == 'False'
                  "
                >
                  <v-btn
                    class="mt-4 font-weight-black"
                    color="#FFD369"
                    elevation="2"
                    x-small
                    width="120px"
                    @click="follow"
                    >Follow</v-btn
                  >
                </v-col>
                <v-col
                  cols="auto"
                  v-if="
                    this.visitorId != this.userId && this.isFollowed == 'True'
                  "
                >
                  <v-btn
                    class="mt-4 font-weight-black"
                    color="#4CAF50"
                    elevation="2"
                    x-small
                    width="120px"
                    @click="unfollow"
                    >Followed</v-btn
                  >
                </v-col>
                <v-col class="mt-n4" cols="12">
                  <a
                    href=""
                    class="text-decoration-none"
                    @click.prevent="goToPosts"
                    >{{ this.userPosts.length }} Posts</a
                  >
                  <a
                    href=""
                    class="ml-5 text-decoration-none"
                    @click.prevent="goToFollow()"
                    >{{ this.followingCount }} Following</a
                  >
                  <a
                    href=""
                    class="ml-5 text-decoration-none"
                    @click.prevent="goToFollow()"
                    >{{ this.followersCount }} Follower</a
                  >
                </v-col>
                <v-col class="mt-n2">
                  <p class="text--secondary mt-3">{{ this.userBio }}</p>
                </v-col>
              </v-row>
            </v-list-item-content>
          </v-list-item>
        </v-card>
      </div>

      <div class="feeds">
        <v-row
          class="px-5 py-3"
          v-for="content in this.userPosts"
          :key="content.id"
        >
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 16 : 2"
              :class="{ 'on-hover': hover }"
              class="mx-auto px-3"
              color="#FFFFFF"
              elevate="0"
              width="900"
              @click.native="viewPost(content.id)"
            >
              <v-card-title class="ml-n3">
                <v-list-item-avatar color="grey darken-3">
                  <v-img class="elevation-6" alt="" :src="userAvatar" />
                </v-list-item-avatar>
                <p class="pt-5" style="color: #393e46">
                  <b>{{ content.fullname }}</b>
                </p>
              </v-card-title>

              <v-card-text
                class="headline font-weight-normal"
                style="color: #393e46"
                v-if="content.photo == ''"
              >
                <v-expand-transition>
                  <div
                    v-if="hover"
                    class="d-flex transition-fast-in-fast-out orange darken-3 v-card--reveal display-1 white--text ml-n7"
                    style="height: 100%"
                  >
                    Show Post
                  </div>
                </v-expand-transition>
                {{ content.caption }}
              </v-card-text>

              <v-container v-if="content.photo != ''">
                <v-card-text
                  class="mt-n10 ml-n3 font-weight-normal"
                  style="color: #393e46"
                >
                  {{ content.caption }}
                </v-card-text>
                <v-img
                  class="mx-1"
                  :src="content.photo"
                  aspect-ratio="1"
                  max-height="400"
                >
                  <v-expand-transition>
                    <div
                      v-if="hover"
                      class="d-flex transition-fast-in-fast-out orange darken-3 v-card--reveal display-1 white--text"
                      style="height: 100%"
                    >
                      Show Post
                    </div>
                  </v-expand-transition>
                </v-img>
              </v-container>

              <v-card-actions>
                <v-list-item class="grow">
                  <v-row align="center" justify="end">
                    <v-icon class="mr-1" color="pink">mdi-heart</v-icon>

                    <span class="subheading mr-2">{{
                      content.like.length
                    }}</span>
                  </v-row>
                </v-list-item>
              </v-card-actions>
            </v-card>
          </v-hover>
        </v-row>
      </div>
    </v-main>
  </v-app>
</template>


<script>
import axios from "axios";
export default {
  name: "Profile",
  mounted() {
    this.getId();
    this.getUserData();
    this.getAllPostsID();
    this.getVisitorId();
    this.getVisitorData();
    this.getFollowers();
    this.getFollowing();
    this.getFollowStatus();
  },
  data() {
    return {
      visitorId: "",
      visitorFullname: "",
      visitorAvatar: "",
      userId: "",
      userFullName: "",
      userPassword: "",
      userGender: "",
      userPhone: "",
      userBio: "",
      userPostsID: [],
      userPosts: [],
      userAvatar: "",
      followersCount: "",
      followingCount: "",
      userFollower: [],
      followID: "",
      isFollowed: "False",
    };
  },
  methods: {
    getVisitorId() {
      this.visitorId = this.$route.params.vistId;
    },
    getId() {
      this.userId = this.$route.params.userId;
    },
    getUserData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.userId)
        .then((response) => {
          this.userId = response.data.data.id;
          this.userFullName = response.data.data.fullname;
          this.userBio = response.data.data.bio;
          this.userAvatar = response.data.data.profile;
        });
    },
    getVisitorData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.visitorId)
        .then((response) => {
          this.visitorFullname = response.data.data.fullname;
          this.visitorAvatar = response.data.data.profile;
        });
    },
    getAllPostsID() {
      axios
        .get(`http://localhost:8081/getAllUserPost/` + this.userId)
        .then((response) => {
          this.userPostsID = response.data.data;
          this.userPosts = this.getAllPosts(this.userPostsID);
          // console.log(this.userPosts);
        });
    },
    getAllPosts(postsID) {
      var posts = [];
      var length = postsID.length;
      for (var i = 0; i < length; i++) {
        axios
          .get(`http://localhost:8081/getPost/` + postsID[i].ID)
          .then((response) => {
            posts.push(response.data.data);
          });
      }
      return posts;
    },
    getFollowers() {
      axios
        .get(`http://localhost:8081/getFollowers/` + this.userId)
        .then((response) => {
          this.followersCount = response.data.data.Count;
        });
    },
    getFollowing() {
      axios
        .get(`http://localhost:8081/getFollowing/` + this.userId)
        .then((response) => {
          this.followingCount = response.data.data.Count;
        });
    },
    getFollowStatus() {
      axios
        .get(`http://localhost:8081/getFollowers/` + this.userId)
        .then((response) => {
          this.userFollower = response.data.data.Followers;

          var follower = this.userFollower.find(
            (fol) => fol.user_id == this.visitorId
          );
          if (follower != null) {
            this.followID = follower.ID;
            this.isFollowed = "True";
          }
        });
    },
    goToPosts() {
      this.$router.push({
        path: "/" + this.userId + "/profile/" + this.visitorId,
      });
    },
    goToBookmark() {
      this.$router.push({ path: "/" + this.userId + "/bookmark" });
    },
    editProfile() {
      this.$router.push({ path: "/" + this.userId + "/editProfile" });
    },
    logout() {
      this.$router.push({ path: "/" });
    },
    goHome() {
      this.$router.push({ path: "/home/" + this.visitorId });
    },
    follow() {
      var followObj = {
        user_id: parseInt(this.visitorId),
        follow_id: parseInt(this.userId),
      };

      axios
        .post(`http://localhost:8081/postFollow`, followObj)
        .then((response) => {
          console.log(response);
          window.alert("You Followed " + this.userFullName);
          this.reloadPage();
        })
        .catch(function (error) {
          window.alert("Followed Fail");
          console.log(error);
        });
    },
    unfollow() {
      axios
        .delete(`http://localhost:8081/deleteFollow/` + this.followID)
        .then((response) => {
          console.log(response);
          window.alert("Unfollowed " + this.userFullName);
          this.reloadPage();
        });
    },
    reloadPage() {
      window.location.reload();
    },
    viewPost(postid) {
      this.$router.push({ path: "/post/" + postid + "/" + this.visitorId });
    },
    goToProfile() {
      this.$router.push({
        path: "/" + this.visitorId + "/profile/" + this.visitorId,
      });
      this.reloadPage();
    },
    goToFollow() {
      this.$router.push({
        path: "/" + this.userId + "/showFollow/" + this.visitorId,
      });
    },
  },
};
</script>

<style>
.v-card--reveal {
  align-items: center;
  bottom: 0;
  justify-content: center;
  opacity: 0.7;
  position: absolute;
  width: 100%;
  cursor: pointer;
}
</style>