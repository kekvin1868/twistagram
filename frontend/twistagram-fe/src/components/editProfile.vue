<template>
    <v-app id="app">
        <v-app-bar app color = "#393E46" dark>
            <v-img
                class="ml-3 my-auto"
                max-height="100"
                max-width="130"
                src="../assets/twistagram-logo.png"/>
        </v-app-bar>

        <v-main>
            <div class="userProfile">
                <v-card 
                    class="mx-auto mt-n8 py-5"
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
                                        x-small>Edit Profile</v-btn>
                                </v-col>
                                <v-col cols="auto">
                                    <v-btn
                                        class="ml-n3 mt-4 font-weight-black"
                                        color="#F85151"
                                        elevation="2"
                                        x-small
                                        width="95px">Logout</v-btn>
                                </v-col>
                                <v-col class="mt-n4" cols="12">
                                    <a href=# class="text-decoration-none">10 Posts</a>
                                    <a href=# class="ml-5 text-decoration-none">10 Following</a>
                                    <a href=# class="ml-5 text-decoration-none">10 Follower</a>
                                </v-col>
                                <v-col class="mt-n2">
                                    <p class="text--secondary mt-3">relating to or dependent on charity; charitable.
                                        "an eleemosynary educational institution."</p>
                                </v-col>
                            </v-row>    
                        </v-list-item-content>
                    </v-list-item>
                </v-card>
            </div>

            <div class="feeds">
                <v-card  
                    class="mx-auto mt-5 py-5"
                    max-width="900"
                    elevate="0">
                    <v-row class="mx-16">
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
                            clearable
                            disabled/>
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
                    <v-row class="mx-auto">
                        <div class="form-button">
                            <v-btn
                                class="mr-3"
                                depressed
                                color="primary"
                                width="120px">Save</v-btn>
                            <v-btn
                                depressed
                                color="error"
                                width="120px">Cancel</v-btn>
                        </div>
                    </v-row>
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
    },
    data() {
        return {
            userId: "",
            userFullName: "",
            userEmail: "",
            userPassword: "",
            userGender: "",
            userPhone:"",
            userBio: ""
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