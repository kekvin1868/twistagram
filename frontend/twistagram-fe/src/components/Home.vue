<template>
  <v-app id="app">
    <v-app-bar app color="#222831" dark>
      <div class="d-flex align-center">
        <v-img
          alt="twistagram Logo"
          class="shrink mr-2"
          contain
          src="../assets/twistagram-logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>

      <v-spacer></v-spacer>

      <v-row>
        <v-avatar size="60">
          <v-img src="../assets/kenji.jpg"></v-img>
        </v-avatar>
        <v-col md="9">
          <p>{{ userData.fullname }}</p>
        </v-col>
      </v-row>
    </v-app-bar>

    <v-main>
      <v-app id="app">
        <v-main class="grey lighten-3">
          <v-container>
            <!-- CARD FOR PROFILE -->
            <v-row id="margin">
              <v-col cols="2">
                <v-card class="rounded-xl" color="#222831">
                  <v-card-title>
                    <v-layout align-center justify-center>
                      <v-avatar size="95">
                        <v-img src="../assets/kenji.jpg"> </v-img>
                      </v-avatar>
                    </v-layout>
                  </v-card-title>
                  <v-card-text>
                    <v-layout align-center justify-center>
                      <h3>
                        {{ userData.fullname }}
                      </h3>
                    </v-layout>
                    <v-divider class="mt-3" color="white"> </v-divider>
                    <h3 class="text-center" v-for="item in profile" :key="item">
                      {{ item }}
                      <v-divider class="mt-3" color="white"> </v-divider></h3
                  ></v-card-text>
                </v-card>

                <!-- CARD FOR SUGGESTION -->
                <v-card class="mt-5 rounded-xl" color="#222831">
                  <v-card-title>
                    <v-layout align-center justify-center>
                      <h3>Suggestion</h3>
                    </v-layout>
                  </v-card-title>
                  <v-divider color="white"> </v-divider>

                  <v-row v-for="x in suggestions" :key="x.image">
                    <v-card-text>
                      <v-col>
                        <v-avatar size="50">
                          <img :src="x.image" />
                        </v-avatar>
                      </v-col>
                    </v-card-text>
                  </v-row>
                </v-card>
              </v-col>

              <!-- CARD FOR UPLOAD CONTENT & POST-->
              <v-col sm="6">
                <v-card
                  class="mx-auto rounded-xl"
                  max-width="800"
                  color="#222831"
                >
                  <v-card-text>
                    <v-row no-gutters>
                      <v-col md="3">
                        <v-card-title>
                          <v-avatar size="75">
                            <v-img src="../assets/kenji.jpg"> </v-img>
                          </v-avatar>
                        </v-card-title>
                      </v-col>

                      <v-col md="8">
                        <v-textarea
                          placeholder="Type something here"
                          background-color="white"
                          rounded
                        >
                        </v-textarea>
                      </v-col>
                    </v-row>

                    <v-divider color="white"> </v-divider>
                    <v-card-actions>
                      <v-file-input
                        label="File input"
                        prepend-icon="mdi-image"
                        v-model="imgData"
                      >
                      </v-file-input>
                      <v-spacer></v-spacer>
                      <v-btn color="primary" @click="post"> post </v-btn>
                    </v-card-actions>
                  </v-card-text>
                </v-card>

                <!-- posting card -->
                <div v-if="this.postData != null">
                  <div v-for="(data, i) in postData" :key="i">
                    <v-card
                      class="mx-auto my-10 rounded-xl"
                      max-width="800"
                      v-if="withPhoto"
                    >
                      <v-card-title>
                        <v-avatar size="70">
                          <v-img src="../assets/kenji.jpg"></v-img>
                        </v-avatar>

                        <p class="ml-5">{{ data.fullname }}</p>
                        <v-spacer></v-spacer>

                        <v-card-actions>
                          <v-btn icon>
                            <v-icon color="black" x-large>
                              mdi-bookmark
                            </v-icon>
                          </v-btn>
                          <v-btn icon>
                            <v-icon x-large color="black">
                              mdi-dots-vertical
                            </v-icon>
                          </v-btn>
                        </v-card-actions>
                      </v-card-title>
                      <v-img src="../assets/kenji.jpg" height="750px"></v-img>
                      <v-card-text>
                        <v-row no-gutters>
                          <v-col sm="5">
                            <h2 class="ml-5">{{ data.fullname }}</h2>
                          </v-col>
                        </v-row>

                        <v-row>
                          <v-col>
                            <h2 class="ml-5">
                              {{ data.caption }}
                            </h2>
                          </v-col>
                        </v-row>
                        <br />

                        <v-divider color="black"></v-divider>
                        <v-row no-gutters>
                          <v-col sm="2">
                            <v-btn class="ml-2" icon x-large color="pink">
                              <v-icon>mdi-heart</v-icon>
                            </v-btn>
                          </v-col>

                          <v-col cols="1" sm="1" md="9">
                            <v-text-field
                              single-line
                              display="flex"
                              class="shrink ml-n5"
                              rounded
                              background-color="grey"
                              color="black"
                              label="Insert your comment here..."
                              v-model="caption"
                            />
                            <v-btn @click="comment(data.id)" color="primary"
                              >Send it bitches</v-btn
                            >
                          </v-col>
                        </v-row>
                      </v-card-text>
                    </v-card>
                    <div v-if="data.comment != null">
                      <div v-for="(comment, i) in data.comment" :key="i">
                        {{ comment.FullName }}
                        {{ comment.Content }}
                      </div>
                    </div>
                  </div>
                </div>
                <div v-if="this.postData == null">
                  <!-- tampilin sesuatu jika ga ada post -->
                </div>

                <br /><br />
              </v-col>

              <!-- CARD FOR TRENDING -->
              <v-col sm="2">
                <v-card class="rounded-xl" color="#222831">
                  <h1 class="h1">Trending</h1>
                  <v-divider class="ml-7" width="220" color="white"></v-divider>
                  <div class="p">
                    <p>#ITHB</p>
                    <p>#ITHB</p>
                    <p>#ITHB</p>
                  </div>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-main>
      </v-app>
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      imgData: null,
      withPhoto: true,
      caption: "",
      userId: "",
      userData: [],
      profile: {
        Following: "Following" + " " + 115,
        Followers: "Followers" + " " + 120,
        ViewProfile: "View Profile",
      },
      suggestions: [
        {
          image: "/images/kenji.jpg",
          text: "Budi",
        },
        {
          image: "../assets/logo.png",
          text: "lol",
        },
      ],
      model: 1,
      postDataTemp: [],
      postData: [],
    };
  },

  methods: {
    async post() {
      let img = null;
      const reader = new FileReader();
      reader.readAsDataURL(this.imgData);
      reader.onload = () => {
        img = reader.result;
        console.log(img);
        axios.post('http://localhost:8081/post',{
          caption:"test",
          photo: img,
          user_id: parseInt(this.userId)
        })
      };
    },
    getUserId() {
      this.userId = this.$route.params.userId;
    },
    getUserData() {
      axios
        .get(`http://localhost:8081/getUserData/` + this.userId)
        .then((response) => {
          this.userData = response.data.data;
        });
    },
    async getPostData() {
      await axios
        .get("http://localhost:8081/getAllUserPost/" + this.userId)
        .then((response) => {
          this.postDataTemp = response.data.data;
        });
      if (this.postDataTemp != null) {
        for (let index = 0; index < this.postDataTemp.length; index++) {
          axios
            .get("http://localhost:8081/getPost/" + this.postDataTemp[index].ID)
            .then((response) => {
              this.postData.push(response.data.data);
            });
        }
      }
    },
    comment(id) {
      let param = {
        user_id: parseInt(this.userId),
        post_id: id,
        content: this.caption,
      };

      axios.post("http://localhost:8081/postComment", param);
    },
  },

  mounted() {
    this.getUserId();
    this.getUserData();
    this.getPostData();
  },
};
</script>

<style>
#app {
  background-color: var(--v-background-base);
}

.h1 {
  padding-left: 25px;
  padding-top: 5px;
  padding-bottom: 5px;
  color: white;
}
.p {
  padding-left: 20px;
  padding-top: 5px;
  color: white;
}

#margin {
  margin-left: 280px;
}

h3 {
  color: white;
}
</style>