<template>
    <v-app style="background-color:#222831" dark>
        <v-app-bar app color = "#393E46" dark>
            <v-img
                max-height="100"
                max-width="200"
                src="../assets/twistagram-logo.png"/>
        </v-app-bar>
        <v-main>
            <v-container fluid fill-height>
                <v-layout align-center justify-center>
                    <v-flex xs12 sm8 md4>
                        <v-row>
                            <v-col>
                                <v-img
                                    class="mx-auto mb-5"
                                    max-height="150"
                                    max-width="250"
                                    src="../assets/twistagram-logo.png"/>
                                
                                <v-form>
                                    <v-text-field
                                        id="user-email"
                                        label="E-mail"
                                        background-color="white"
                                        required
                                        single-line
                                        outlined 
                                        clearable />

                                    <v-text-field
                                        id="user-password"
                                        type="password"
                                        label="Password"
                                        background-color="white"
                                        required
                                        single-line
                                        outlined
                                        clearable/>
                                </v-form>

                                <v-btn
                                    class="px-15"
                                    type="submit"
                                    color="primary"
                                    elevation="5"
                                    @click.prevent="login">Login</v-btn>
                                <p class="mt-5 white--text">Don't have an account? <a @click="navRegister"><strong>Sign up</strong></a></p>
                            </v-col>
                        </v-row>
                    </v-flex>
                </v-layout>
                
            </v-container>
        </v-main>  
    </v-app>
</template>


<script>
import axios from 'axios'
export default {
    name: 'Login',

    data() {
        return {
            userId: ""
        };
    },
    methods: {
        async login() {

            var email = document.getElementById('user-email').value;
            var password = document.getElementById('user-password').value;

            var loginObj = {"email":email, "password":password};

            axios.post(`http://localhost:8081/login`,loginObj)
            .then(response=>{
                this.userId = response.data.data.id;
                this.goToHome();
            })
            .catch(function (error) { 
                window.alert("Email or Password is incorrect");
                console.log(error);
            })

        },
        navRegister() {
            this.$router.push({path: "register"})
        },
        goToHome(){
            this.$router.push({path:"/home/"+this.userId})
        }
    }
}
</script>