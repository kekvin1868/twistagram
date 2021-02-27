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
                                    id="user-password"
                                    type="password"
                                    label="Password"
                                    placeholder="Password"
                                    background-color="white"
                                    required
                                    single-line
                                    outlined
                                    clearable/>

                                <v-btn
                                    class="px-15"
                                    type="submit"
                                    color="primary"
                                    elevation="5"
                                    @click.prevent="login">Login</v-btn>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-form>
                <p class="ml-16 pl-7 mt-5 white--text">Don't have an account? <a href="#" @click="navRegister"><strong>Sign up</strong></a></p>
            </v-main>
        </div>
    </v-app>  
</template>


<script>

export default {
    name: 'Login',

    data() {
        return {
            found: null
        }
    },
    methods: {
        login() {
            var email = document.getElementById('user-email').value;
            var password = document.getElementById('user-password').value;

            var loginObj = {"email":email, "password":password};
            console.log(loginObj)
        },
        navRegister() {
            this.$router.push({path: "register"})
        }
    }
}
</script>

<style>
    #app {
        background-color: var(--v-background-base);
    }
</style>