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
                    class="mx-auto mt-8 py-5"
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

            <div class="form">
                <v-card  
                    class="mx-auto mt-8 py-8"
                    max-width="900"
                    elevate="0">
                    <v-row class="mx-16 mt-2">
                        <p class="py-4 mr-10">Name</p>
                        <v-text-field
                            id="user-name"
                            label="Name"
                            placeholder="User Name"
                            background-color="white"
                            v-model="this.userFullName"
                            required
                            single-line
                            outlined 
                            clearable/>
                    </v-row>    
                    <v-row class="mx-16">
                        <p class="py-4 mr-15">Bio</p>
                        <v-textarea
                            id="user-bio"
                            label="Bio"
                            placeholder="User Bio"
                            background-color="white"
                            v-model="this.userBio"
                            required
                            single-line
                            outlined 
                            clearable/>
                    </v-row>
                    <v-row class="mx-16">
                        <p class="py-4 mr-11">Email</p>
                        <v-text-field
                            id="user-email"
                            label="Email"
                            placeholder="User Email"
                            background-color="white"
                            v-model="this.userEmail"
                            required
                            single-line
                            outlined 
                            clearable
                            disabled/>
                    </v-row>    
                    <v-row class="mx-16">
                        <p class="py-4 mr-9">Phone</p>
                        <v-text-field
                            id="user-phone"
                            label="Phone"
                            placeholder="User Phone"
                            background-color="white"
                            v-model="this.userPhone"
                            required
                            single-line
                            outlined 
                            clearable/>
                    </v-row>    
                    <v-row class="mx-16">
                        <p class="py-4 mr-8">Gender</p>
                        <v-text-field
                            id="user-gender"
                            label="Gender"
                            placeholder="User Gender"
                            background-color="white"
                            v-model="this.userGender"
                            required
                            single-line
                            outlined 
                            clearable
                            disabled/>
                    </v-row>  
                    <v-layout-justify-center>
                        <v-flex offset-sm4 md6>
                            <div class="form-button mt-5">
                                <v-btn
                                    class="mr-3"
                                    depressed
                                    color="primary"
                                    width="120px"
                                    @click.prevent="updateProfile">Save</v-btn>
                                <v-btn
                                    depressed
                                    color="error"
                                    width="120px"
                                    @click="goToPosts">Cancel</v-btn>
                            </div>
                        </v-flex>
                    </v-layout-justify-center>
                </v-card>
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
        this.getFollowers();
        this.getFollowing();
        this.getAllPosts();
    },
    data() {
        return {
            userId: "",
            userFullName: "",
            userEmail: "",
            userPassword: "",
            userGender: "",
            userPhone:"",
            userBio: "",
            userPosts: [],
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
                    this.userGender = response.data.data.gender;
                    this.userPhone = response.data.data.phone;
                    this.userBio = response.data.data.bio;
                    this.userEmail = response.data.data.email;
                    this.userPassword = response.data.data.password;
                });
        },
        getAllPosts(){
            axios.get(`http://localhost:8081/getAllUserPost/`+this.userId)
                .then(response=>{
                    this.userPosts = response.data.data;
                });
        },
        updateProfile(){
            var id = this.userId;
            var email = this.userEmail;
            var username = document.getElementById("user-name").value;
            var password = this.userPassword;
            var phone = document.getElementById("user-phone").value;
            var gender = this.userGender;
            var bio = document.getElementById("user-bio").value;

            var userObj = {
                id: id,
                email: email,
                fullname: username,
                password: password,
                phone: phone,
                gender: gender,
                bio: bio
            };

            axios
                .patch(`http://localhost:8081/updateUserData`, userObj)
                .then((response) => {
                    console.log(response);
                    this.goToPosts();
                })
                .catch(function (error) {
                    window.alert("Update Data Failed");
                    console.log(error);
                });
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

    div.form-button {
        justify-content: center;
    }
</style>