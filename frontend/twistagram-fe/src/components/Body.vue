<template>
  <v-container>
    <v-card class="mx-auto" max-width="800" color="#222831" rounded>
      <v-card-title>
        <v-row>
          <v-col cols="12" row-height="2" offset-md="5">
            <v-avatar row="2" rows-height="10" size="80">
              <v-img src="../assets/kenji.jpg"></v-img>
            </v-avatar>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1" offset-xs="10" sm="8">
            <v-textarea
              row="1"
              row-height="15"
              placeholder="What's happening"
              auto-grow
              rounded
              background-color="white"
            >
            </v-textarea>
          </v-col>
          <v-col cols="12" offset-lg="10">
            <v-btn color="warning"> Post </v-btn>
          </v-col>
        </v-row>
      </v-card-title>
    </v-card>
    <v-card class="mx-auto" max-width="650">
      <v-card-title>
        <v-avatar size="70">
          <v-img src="../assets/kenji.jpg"></v-img>
        </v-avatar>

        <p class="ml-5">Kenji Edrich Lyn</p>

        <v-row>
          <v-col cols="1" sm="2" offset-md="8">
            <v-icon color="black" x-large> mdi-bookmark </v-icon>
          </v-col>
          <v-icon x-large color="black"> mdi-dots-vertical </v-icon>
        </v-row>
      </v-card-title>
      <v-img src="../assets/kenji.jpg" height="750px"></v-img>
      <v-card-text>
        <v-container>
          <v-row no-gutters>
            <v-col cols="12" sm="3">
              <p class="ml-5">Kenji Edrich Lyn</p>
            </v-col>
          </v-row>

          <v-row no-gutters>
            <v-col cols="12" sm="2" lg="1" offset-sm="100">
              <v-btn icon x-large color="pink">
                <v-icon>mdi-heart</v-icon>
              </v-btn>
            </v-col>

            <v-col cols="1" sm="1" md="9">
              <v-text-field
                single-line
                display="flex"
                class="shrink"
                rounded
                background-color="grey"
                color="black"
                label="Insert your comment here..."
              >
              </v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>

      <!-- Sementara di copy, nanti setelah jadi akan di for -->

      <v-card-title>
        <v-avatar size="70">
          <v-img src="../assets/twistagram.png"></v-img>
        </v-avatar>

        <p class="ml-5">Kenji Edrich Lyn</p>

        <v-row>
          <v-col cols="1" sm="2" offset-md="8">
            <v-icon color="black" x-large> mdi-bookmark </v-icon>
          </v-col>
          <v-icon x-large color="black"> mdi-dots-vertical </v-icon>
        </v-row>
      </v-card-title>
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="850px"
      ></v-img>

      <v-card-text>
        <v-container>
          <v-row no-gutters>
            <v-col cols="12" sm="3">
              <p class="ml-5">Kenji Edrich Lyn</p>
            </v-col>
          </v-row>

          <v-row no-gutters>
            <v-col cols="12" sm="2" lg="1" offset-sm="100">
              <v-btn icon x-large color="pink">
                <v-icon>mdi-heart</v-icon>
              </v-btn>
            </v-col>

            <v-col cols="1" sm="1" md="9">
              <v-text-field
                single-line
                display="flex"
                class="shrink"
                rounded
                background-color="grey"
                color="black"
                label="Insert your comment here..."
              >
              </v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
    </v-card>
  </v-container>
</template>