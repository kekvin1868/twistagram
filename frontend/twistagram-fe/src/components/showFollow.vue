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
        <v-img :src="visitorAvatar"></v-img>
      </v-avatar>
      <p class="mt-3 ml-3 mr-13">
        <a
          href=""
          class="text-decoration-none"
          style="color: white"
          @click.prevent="goToProfile()">{{ visitorFullname }}</a></p>
    </v-app-bar>

    <v-main>
        <v-card
            class="my-8 mx-auto"
            width="900px"
            elevation="0">

            <div class="userProfile">
                <v-card class="mx-auto mt-8 mb-5 py-5" max-width="900" elevate="0">
                <v-list-item>
                    <v-avatar class="my-auto ml-7" size="140">
                        <v-img
                            lazy-src="../assets/default-profile.jpg"
                            :src="userAvatar"
                            alt="profilePicture"/>
                    </v-avatar>

                    <v-list-item-content class="ml-10 my-auto">
                    <v-row>
                        <v-col cols="auto">
                        <p class="mt-1 display-1 text--primary">
                            {{ this.userFullName }}
                        </p>
                        </v-col>
                        <v-col cols="auto" v-if="visitorId == userId">
                        <v-btn
                            class="mt-4 font-weight-black"
                            color="#FFD369"
                            elevation="2"
                            x-small
                            @click="editProfile">Edit Profile</v-btn>
                        </v-col>
                        <v-col cols="auto" v-if="visitorId == userId">
                        <v-btn
                            class="ml-n3 mt-4 font-weight-black"
                            color="#4CAF50"
                            elevation="2"
                            x-small
                            width="95px"
                            @click="goToBookmark">Bookmarked</v-btn>
                        </v-col>
                        <v-col cols="auto" v-if="visitorId == userId">
                        <v-btn
                            class="ml-n3 mt-4 font-weight-black"
                            color="#F85151"
                            elevation="2"
                            x-small
                            width="95px"
                            @click="logout">Logout</v-btn>
                        </v-col>
                        <v-col
                        cols="auto"
                        v-if="
                            visitorId != userId && isFollowed == 'False'">
                        <v-btn
                            class="mt-4 font-weight-black"
                            color="#FFD369"
                            elevation="2"
                            x-small
                            width="120px"
                            @click="follow">Follow</v-btn>
                        </v-col>
                        <v-col
                            cols="auto"
                            v-if="
                                visitorId != userId && isFollowed == 'True'">
                        <v-btn
                            class="mt-4 font-weight-black"
                            color="#4CAF50"
                            elevation="2"
                            x-small
                            width="120px"
                            @click="unfollow">Followed</v-btn>
                        </v-col>
                        <v-col class="mt-n4" cols="12">
                        <a
                            href=""
                            class="text-decoration-none"
                            @click.prevent="goToPosts">{{ this.userPosts.length }} Posts</a>
                        <a href="" class="ml-5 text-decoration-none" @click.prevent="reloadPage()">{{ this.followingCount }} Following</a>
                        <a href="" class="ml-5 text-decoration-none" @click.prevent="reloadPage()">{{ this.followersCount }} Follower</a>
                        </v-col>
                        <v-col class="mt-n2">
                        <p class="text--secondary mt-3">{{ this.userBio }}</p>
                        </v-col>
                    </v-row>
                    </v-list-item-content>
                </v-list-item>
                </v-card>
            </div>

            <v-tabs
                class="rounded"
                v-model="tab"
                background-color="#222831"
                centered
                dark
                icons-and-text>
                <v-tabs-slider></v-tabs-slider>

                <v-tab href="#tab-1">
                    Following
                </v-tab>

                <v-tab href="#tab-2">
                    Follower
                </v-tab>
            </v-tabs>

            <v-tabs-items v-model="tab">
                <v-tab-item :value="'tab-' + 1">
                    <v-card flat>
                        <v-list-item class="mx-7" v-for="following in userFollowing" :key="following.id">
                            <v-avatar class="my-auto" size="40">
                                <v-img
                                    lazy-src="../assets/default-profile.jpg"
                                    :src="following.avatar"
                                    alt="profilePicture"/>
                            </v-avatar>

                            <v-list-item-content class="ml-10 my-auto">
                                <v-row>
                                    <v-col cols="auto">
                                    <a href=""
                                        class="text-decoration-none"
                                        style="color: #222831"
                                        @click.prevent="gotoFeeds(following.id)">
                                        <p class="mt-5 text--primary" style="font-size:20px"><b>
                                            {{ following.fullname }}
                                        </b></p>
                                    </a>
                                    </v-col>
                                    <v-col
                                        v-if="visitorId != following.id && following.isFollowed == false"
                                        align="right">
                                        <v-btn
                                            class="mt-5 font-weight-black"
                                            color="#FFD369"
                                            elevation="2"
                                            x-small
                                            width="120px"
                                            @click="followUser(following)">Follow</v-btn>
                                    </v-col>
                                    <v-col
                                        v-if="visitorId != following.id && following.isFollowed == true"
                                        align="right">
                                        <v-btn
                                            class="mt-4 font-weight-black"
                                            color="#4CAF50"
                                            elevation="2"
                                            x-small
                                            width="120px"
                                            @click="unfollowUser(following)">Followed</v-btn>
                                    </v-col>
                                </v-row>
                            </v-list-item-content>
                        </v-list-item>
                    </v-card>
                </v-tab-item>

                <v-tab-item :value="'tab-' + 2">
                    <v-card flat>
                        <v-list-item class="mx-7" v-for="follower in userFollower" :key="follower.id">
                            <v-avatar class="my-auto" size="40">
                                <v-img
                                    lazy-src="../assets/default-profile.jpg"
                                    :src="follower.avatar"
                                    alt="profilePicture"/>
                            </v-avatar>

                            <v-list-item-content class="ml-10 my-auto">
                                <v-row>
                                    <v-col cols="auto">
                                    <a href=""
                                        class="text-decoration-none"
                                        style="color: #222831"
                                        @click.prevent="gotoFeeds(follower.id)">
                                        <p class="mt-5 text--primary" style="font-size:20px"><b>
                                            {{ follower.fullname }}
                                        </b></p>
                                    </a>
                                    </v-col>
                                    <v-col
                                        v-if="visitorId != follower.id && follower.isFollowed == false"
                                        align="right">
                                        <v-btn
                                            class="mt-5 font-weight-black"
                                            color="#FFD369"
                                            elevation="2"
                                            x-small
                                            width="120px"
                                            @click="followUser(follower)">Follow</v-btn>
                                    </v-col>
                                    <v-col
                                        v-if="visitorId != follower.id && follower.isFollowed == true"
                                        align="right">
                                        <v-btn
                                            class="mt-4 font-weight-black"
                                            color="#4CAF50"
                                            elevation="2"
                                            x-small
                                            width="120px"
                                            @click="unfollowUser(follower)">Followed</v-btn>
                                    </v-col>
                                </v-row>
                            </v-list-item-content>
                        </v-list-item>
                    </v-card>
                </v-tab-item>
            </v-tabs-items>
        </v-card>
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
    this.getFollowStatus();
    this.getFollowing();
    this.getFollowers();
  },
  data() {
    return {
      tab:null,
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
      userFollowing: [],
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
          this.userFollower = this.getFollowersData(response.data.data.Followers);
        });
    },
    getFollowing() {
      axios
        .get(`http://localhost:8081/getFollowing/` + this.userId)
        .then((response) => {
          this.followingCount = response.data.data.Count;
          this.userFollowing = this.getFollowingData(response.data.data.Following);
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
    getFollowersData(followers){
        var followerArr = [];

        for(let i=0; i<followers.length; i++) {
            axios
                .get(`http://localhost:8081/getUserData/` + followers[i].user_id)
                .then((response) => {
                    var fullname = response.data.data.fullname;
                    var avatar = response.data.data.profile;

                     axios
                        .get(`http://localhost:8081/getFollowers/` + followers[i].user_id)
                        .then((response) => {
                            var follows = response.data.data.Followers;
                            var isFollowed = false;

                            var follower = follows.find(
                                (fol) => fol.user_id == this.visitorId
                            );
                            
                            if (follower != null) {
                                isFollowed = true;
                            }
                            
                            var followerObj = {
                                id: followers[i].user_id,
                                fullname: fullname,
                                avatar: avatar,
                                isFollowed: isFollowed
                            }
                            followerArr.push(followerObj);
                        });
                });
        }
        return followerArr;
    },
    getFollowingData(followings){
        var following = [];

        for(let i=0; i<followings.length; i++) {
            axios
                .get(`http://localhost:8081/getUserData/` + followings[i].follow_id)
                .then((response) => {
                    var fullname = response.data.data.fullname;
                    var avatar = response.data.data.profile;

                    axios
                        .get(`http://localhost:8081/getFollowers/` + followings[i].follow_id)
                        .then((response) => {
                            var follows = response.data.data.Followers;
                            var isFollowed = false;

                            var follower = follows.find(
                                (fol) => fol.user_id == this.visitorId
                            );
                            
                            if (follower != null) {
                                isFollowed = true;
                            }
                            
                            var followingObj = {
                                id: followings[i].follow_id,
                                fullname: fullname,
                                avatar: avatar,
                                isFollowed: isFollowed
                            }
                            
                            following.push(followingObj);
                        });
                });
        }
        return following;
    },
    goToPosts() {
      this.$router.push({
        path: "/" + this.userId + "/profile/" + this.visitorId,
      });
    },
    gotoFeeds(uid) {
      this.$router.push({
        path: "/" + uid + "/profile/" + this.visitorId,
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
      this.$router.push({ path: "/home/" + this.userId });
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
    followUser(user) {
      var followObj = {
        user_id: parseInt(this.visitorId),
        follow_id: parseInt(user.id),
      };

      axios
        .post(`http://localhost:8081/postFollow`, followObj)
        .then((response) => {
          console.log(response);
          window.alert("You Followed " + user.fullname);
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
    unfollowUser(user) {
        console.log(user);
        axios
            .get(`http://localhost:8081/getFollowers/` + user.id)
            .then((response) => {
                var follows = response.data.data.Followers;
                console.log(follows);

                var follower = follows.find(
                    (fol) => fol.user_id == this.visitorId
                );

                if(follower != null) {
                    axios
                        .delete(`http://localhost:8081/deleteFollow/` + follower.ID)
                        .then((response) => {
                            console.log(response);
                            window.alert("Unfollowed " + user.fullname);
                            this.reloadPage();
                        });
                }
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
    },
    goToFollow() {
      this.$router.push({ path: "/" + this.userId + "/showFollow/" + this.visitorId });
    },
  },
};
</script>