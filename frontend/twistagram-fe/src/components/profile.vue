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
                                    <p class="mt-1 display-1 text--primary">username</p>
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
                    class="mx-auto mt-5"
                    max-width="900"
                    elevate="0">
                    <v-row class="px-5 py-3">
                        <v-col
                            v-for="n in 11"
                            :key="n"
                            class="d-flex child-flex"
                            cols="4">
                            <a :href="`https://picsum.photos/500/300?image=${n * 5 + 10}`">
                                <v-img
                                    :src="`https://picsum.photos/500/300?image=${n * 5 + 10}`"
                                    :lazy-src="`https://picsum.photos/10/6?image=${n * 5 + 10}`"
                                    aspect-ratio="1"
                                    class="grey lighten-2">
                                    <template v-slot:placeholder>
                                        <v-row
                                            class="fill-height ma-0"
                                            align="center"
                                            justify="center">
                                                <v-progress-circular
                                                indeterminate
                                                color="grey lighten-5"/>
                                        </v-row>
                                    </template>
                                </v-img>
                            </a>
                        </v-col>
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
                });
        }
    }
}
</script>

<style>
    #app {
        background-color: var(--v-background2-base);
    }
</style>