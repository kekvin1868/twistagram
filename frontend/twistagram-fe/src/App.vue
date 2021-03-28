<template>
  <v-app>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
  
</template>

</template>

<script>
import HelloWorld from './components/profile';

export default {
  name: 'App',

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
};
</script>
