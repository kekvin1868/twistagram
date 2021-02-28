<template>
  <v-app>
    <v-app-bar
      app
      color="#222831"
      dark>

      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../assets/twistagram-logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
    </v-app-bar>

    <v-main>
      <Body/>
    </v-main>
  </v-app>
</template>

<script>
import Body from '../components/Body';

export default {
    name: 'Home',
    components: {
        Body,
    },

    data: () => ({
      
    })
};
</script>