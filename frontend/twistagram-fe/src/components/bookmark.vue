<template>
    <v-app id="app">
        <v-app-bar app color = "#393E46" dark>
            <a href="">
                <v-img
                    class="ml-3 my-auto"
                    max-height="100"
                    max-width="130"
                    src="../assets/twistagram-logo.png"
                    @click.prevent="goHome"/>
            </a>
        </v-app-bar>

        <v-main>
            <div class="userProfile">
                <v-card 
                    class="mx-auto mt-8 mb-5 py-5"
                    max-width="900"
                    elevate="0">
                    <v-list-item>
                        <v-avatar
                            class="my-auto ml-7"
                            size="140">
                                <img
                                    lazy-src="../assets/default-profile.jpg"
                                    src="../assets/default-profile.jpg"
                                    alt="profilePicture"/>
                        </v-avatar>
                            
                        <v-list-item-content class="ml-10 my-auto">
                            <v-row>
                                <v-col cols="auto">
                                    <p class="mt-1 display-1 text--primary">{{this.userFullName}}</p>
                                </v-col>
                                <v-col cols="auto">
                                    <v-btn
                                        class="mt-4 font-weight-black"
                                        color="#FFD369"
                                        elevation="2"
                                        x-small
                                        @click="editProfile">Edit Profile</v-btn>
                                </v-col>
                                <v-col cols="auto">
                                    <v-btn
                                        class="ml-n3 mt-4 font-weight-black"
                                        color="#4CAF50"
                                        elevation="2"
                                        x-small
                                        width="95px"
                                        @click="goToBookmark">Bookmarked</v-btn>
                                </v-col>
                                <v-col cols="auto">
                                    <v-btn
                                        class="ml-n3 mt-4 font-weight-black"
                                        color="#F85151"
                                        elevation="2"
                                        x-small
                                        width="95px"
                                        @click="logout">Logout</v-btn>
                                </v-col>
                                <v-col class="mt-n4" cols="12">
                                    <a href="" class="text-decoration-none" @click="goToPosts">{{this.userPosts.length}} Posts</a>
                                    <a href="" class="ml-5 text-decoration-none">{{this.followingCount}} Following</a>
                                    <a href="" class="ml-5 text-decoration-none">{{this.followersCount}} Follower</a>
                                </v-col>
                                <v-col class="mt-n2">
                                    <p class="text--secondary mt-3">{{this.userBio}}</p>
                                </v-col>
                            </v-row>    
                        </v-list-item-content>
                    </v-list-item>
                </v-card>
            </div>

            <div class="bookmark">
                <v-row class="px-5 py-3" v-for="content in this.userBookmark" :key="content">
                    <v-card
                        class="mx-auto px-3"
                        color="#FFFFFF"
                        elevate="0"
                        width="900">
                        
                        <v-card-title class="ml-n3">
                            <v-list-item-avatar color="grey darken-3">
                                <v-img
                                    class="elevation-6"
                                    alt=""
                                    src="../assets/default-profile.jpg"/>
                            </v-list-item-avatar>
                            <p class="pt-5" style="color:#393E46"><b>{{content.fullname}}</b></p>
                        </v-card-title>
                        
                        <v-card-text class="headline font-weight-normal" style="color:#393E46" v-if="content.photo==''">
                            {{content.caption}}
                        </v-card-text>

                        <v-container v-if="content.photo!=''">
                            <v-card-text class="mt-n10 ml-n3 font-weight-normal" style="color:#393E46">
                                {{content.caption}}
                            </v-card-text>
                            <v-img
                                class="mx-1"
                                src="../assets/kenji.jpg"
                                aspect-ratio="1"
                                max-height="400"/>
                        </v-container>

                        <v-card-actions>
                            <v-list-item class="grow">
                                <v-row
                                    align="center"
                                    justify="end">
                                    <v-btn icon>
                                        <v-icon class="mr-1" disabled>mdi-heart</v-icon>
                                    </v-btn>
                                    <span class="subheading mr-2">{{content.like.length}}</span>
                                </v-row>
                            </v-list-item>
                        </v-card-actions>
                    </v-card>
                </v-row>
            </div>
        </v-main>
    </v-app>  
</template>


<script>
import axios from 'axios'
export default {
    name: 'Profile',
    mounted(){
        this.getId();
        this.getUserData();
        this.getAllPosts();
        this.getBookmarksID();
        this.getFollowers();
        this.getFollowing();
    },
    data() {
        return {
            userId: "",
            userFullName: "",
            userPassword: "",
            userGender: "",
            userPhone:"",
            userBio: "",
            userPosts: [],
            userBookmarkID: [],
            userBookmark: [],
            followersCount: "",
            followingCount: "",
        }
    },
    methods: {
        getId(){
            this.userId = this.$route.params.userId;
        },
        getUserData(){
            axios.get(`http://localhost:8081/getUserData/`+this.userId)
                .then(response=>{
                    this.userId = response.data.data.id;
                    this.userFullName = response.data.data.fullname;
                    this.userBio = response.data.data.bio;
                });
        },
        getAllPosts(){
            axios.get(`http://localhost:8081/getAllUserPost/`+this.userId)
                .then(response=>{
                    this.userPosts = response.data.data;
                });
        },
        getBookmarksID(){
            axios.get(`http://localhost:8081/getBookmark/`+this.userId)
                .then(response=>{
                    this.userBookmarkID = response.data.data;
                    this.userBookmark = this.getBookmarks(this.userBookmarkID);
                    console.log(this.userBookmark);
                });
        },
        getBookmarks(postsID){
            var posts = [];
            var length = postsID.length;
            for(var i=0; i<length; i++) {
                axios.get(`http://localhost:8081/getPost/`+postsID[i].post_id)
                    .then(response=>{
                        posts.push(response.data.data);
                    });
            }
            return posts;
        },
        getFollowers(){
            axios.get(`http://localhost:8081/getFollowers/`+this.userId)
                .then(response=>{
                    this.followersCount = response.data.data.Count;
                });
        },
        getFollowing(){
            axios.get(`http://localhost:8081/getFollowing/`+this.userId)
                .then(response=>{
                    this.followingCount = response.data.data.Count;
                });
        },
        goToPosts(){
            this.$router.push({path:"/"+this.userId+"/profile/"+this.userId})
        },
        goToBookmark(){
            this.$router.push({path:"/"+this.userId+"/bookmark"})
        },
        editProfile(){
            this.$router.push({path:"/"+this.userId+"/editProfile"})
        },
        logout(){
            this.$router.push({path:"/"})
        },
        goHome(){
            this.$router.push({path: "/home/"+this.userId})
        }
    }
}
</script>

<style>
    #app {
        background-color: var(--v-background2-base);
    }
</style>