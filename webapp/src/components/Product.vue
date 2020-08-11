<template>
  <div id="product">
    <v-card
        :loading="loading"
        class="mx-auto my-4 px-4 py-4"
        max-width="300px"
        elevation="3"
    >

        <v-img
            height="100%"
            width="100%"
            :src="'https://' + this.imageLink"
            aspect-ratio="1"
        />

      <v-card-title class="display-1">{{ this.brand }}</v-card-title>
      <v-card-text>
        <div class="text-body-1">{{ this.description }}</div>
        <v-row
            align="center"
            class="mx-0 mt-1"
        >
          <v-rating
              :value="this.randomRating()"
              color="amber"
              dense
              half-increments
              readonly
              size="14"
          ></v-rating>
        </v-row>
      </v-card-text>
      <v-divider class="mx-4"></v-divider>
      <v-card-title class="title">Precio</v-card-title>
      <v-card-text>
        <v-chip-group
            active-class="deep-purple accent-4 white--text"
            column
        >
          <v-chip class="title">${{ this.price }}</v-chip>
          <v-chip v-if="this.palindrome" class="title red" dark>50%</v-chip>
        </v-chip-group>
        <v-chip-group
            v-if="this.palindrome"
            active-class="deep-purple accent-4 white--text"
            column
        >
          <v-chip class="title light-green" dark>${{ this.palindromeRealPrice }}</v-chip>
        </v-chip-group>
      </v-card-text>
      <v-card-actions>
        <v-btn
            color="deep-purple lighten-2"
            text
        >
           + Agregar al carro
        </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>
<script>
export default {
  name: "Product",
  props: {
    brand: String,
    description: String,
    imageLink: String,
    price: Number,
    palindrome: Boolean,
  },
  data: () => ({
    loading: true
  }),
  computed: {
    palindromeRealPrice: () => this.price * 2,
  },
  methods: {
    randomRating: () => {
      return Math.random() * (5 - 2) + 2
    }
  },
  mounted() {

  }
}
</script>
