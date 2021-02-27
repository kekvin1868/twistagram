<template>
    <v-app id="app">
        <v-app-bar app color = "#393E46" dark></v-app-bar>

        <div class="mt-n5">
            <v-main>
                <v-form ref="form" lazy-validation>
                    <v-container>
                        <v-row>
                            <v-col cols="12" md="6">
                                <v-img
                                    class="mx-auto mb-5"
                                    max-height="150"
                                    max-width="250"
                                    src="../assets/twistagram-logo.png"/>
                                
                                <v-text-field
                                    id="user-email"
                                    label="E-mail"
                                    placeholder="E-mail"
                                    background-color="white"
                                    required
                                    single-line
                                    outlined 
                                    clearable />

                                <v-text-field
                                    id="user-name"
                                    label="Username"
                                    placeholder="Username"
                                    background-color="white"
                                    required
                                    single-line
                                    outlined
                                    clearable />

                                <v-text-field
                                    id="user-password"
                                    type="password"
                                    label="Password"
                                    placeholder="Password"
                                    background-color="white"
                                    required
                                    single-line
                                    outlined
                                    clearable/>

                                <v-text-field
                                    id="user-phone"
                                    label="Phone"
                                    placeholder="Phone"
                                    background-color="white"
                                    required
                                    single-line
                                    outlined
                                    clearable/>

                                <v-radio-group
                                    class="mt-n4"
                                    row>
                                        <v-radio
                                            name="gender"
                                            label="Male"
                                            value="Male"
                                            color="indigo"/>
                                        <v-radio
                                            name="gender"
                                            label="Female"
                                            value="Female"
                                            color="indigo"/>
                                    </v-radio-group>

                                <v-btn
                                    type="submit"
                                    color="primary"
                                    elevation="5"
                                    @click.prevent="register"
                                    >Register This Account</v-btn>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-form>
                <p class="ml-16 pl-7 mt-5 white--text">Already have an account? <a href="#" @click="navLogin"><strong>Log In</strong></a></p>
            </v-main>
        </div>
    </v-app>  
</template>


<script>
export default {
    name: 'Register',

    data() {
        return {
            found: null
        }
    },
    methods: {
        register() {
            var email = document.getElementById('user-email').value;
            var username = document.getElementById('user-name').value;
            var password = document.getElementById('user-password').value;
            var phone = document.getElementById('user-phone').value;
            var gender = "";
            var genders = document.getElementsByName('gender');

            for (var i=0, length=genders.length; i<length; i++) {
                if (genders[i].checked) {
                    gender = genders[i].value;
                }
            }

            var userObj = {"email":email, "username":username, "password":password, "phone":phone, "gender":gender};
            console.log(userObj)


        },
        navLogin() {
            
        }
    }
}
</script>

<style>
    #app {
        background-color: var(--v-background-base);
    }
</style>