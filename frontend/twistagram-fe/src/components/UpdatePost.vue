<template>
  <v-app id="app">
    <v-app-bar app color="#222831" dark>
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../assets/twistagram-logo.png"
          transition="scale-transition"
          width="200"/>
      </div>
      <v-spacer> </v-spacer>
      <v-avatar size="50">
        <v-img src="../assets/kenji.jpg"></v-img>
      </v-avatar>
      <p>JohnDoe</p>
    </v-app-bar>

    <v-main>
      <v-card
        class="mx-auto mt-8"
        max-width="1000"
        outlined
        rounded
        elevation="5">
        
        <v-row class="ma-3">
          <v-col>
            <v-img
              src="../assets/kenji.jpg"
              max-height="800"
              aspect-ratio="1"/>
          </v-col>
          
          <v-col class="mt-5 ml-3">
            <v-row>
              <h2>Update Post:</h2>
            </v-row>
            <v-row>
              <v-textarea
                class="mt-5"
                id="post-caption"
                label="Caption"
                placeholder="Post Caption"
                v-model="this.postCaption"
                clearable/>
            </v-row>
            <v-row>
              <v-col class="ml-n3">
                <v-btn 
                  class="mr-3"
                  width="120px"
                  color="primary"
                  @click="savePost">Save</v-btn>
                <v-btn
                  width="120px"
                  color="error"
                  @click="goToPost">Cancel</v-btn>
              </v-col>
            </v-row>
          </v-col>  
        </v-row>  
        
      </v-card>
    </v-main>
  </v-app>
</template>

<style scoped>
.button {
  padding-top: 180px;
  margin-left: 90px;
}

.table {

  margin-top: 150px;
  margin-left: 390px;
}

.caption {
  margin-top: 5px;
  margin-left: 100px;
  padding-bottom: 100px;
}
</style>

<script>
import axios from 'axios'

export default {
    name: 'updatePost',
    mounted(){
      this.getPostId();
      this.getPostData();
    },
    data() {
        return {
          userId: "",
          postId: "",
          postLike: [],
          postComment: [],
          postCaption: "",
          postFullname: "",
        }
    },
    methods: {
        getPostId(){
          this.postId = this.$route.params.postId;
        },
        getPostData(){
          axios.get(`http://localhost:8081/getPost/`+this.postId)
            .then(response=>{
                this.postLike = response.data.data.like;
                this.postComment = response.data.data.comment;
                this.postCaption = response.data.data.caption;
                this.userId = response.data.data.user_id;
                this.postFullname = response.data.data.fullname;
          });
        },
        savePost(){
          // var pid = this.postId;
          // var uid = this.userId;
          // var likes = this.postLike;
          // var comments = this.postComment;
          // var fullname = this.postFullname;
          // var caption = document.getElementById("post-caption").value;

          // var postObj = {
          //   id: pid,
          //   caption: caption,
          //   user_id: uid,
          //   fullname: fullname,
          //   like: likes,
          //   comment: comments
          // }
          
          // axios
            // .patch()
            // .then((response) => {
          //     console.log(response);
          //     this.goToPost();
          //   })
          //   .catch(function (error) {
          //     window.alert("Update Data Failed");
          //     console.log(error);
          //   });
        },
        goToPost(){
          this.$router.push({path: "/post/"+this.postId+this.userId})
        }
    }
}
</script>