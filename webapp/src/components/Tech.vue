<template>
  <v-app id="inspire">
    <v-navigation-drawer
        v-model="drawer"
        :clipped="$vuetify.breakpoint.lgAndUp"
        app
    >
      <v-list dense>
        <template v-for="item in items">
          <v-row
              v-if="item.heading"
              :key="item.heading"
              align="center"
          >
            <v-col cols="6">
              <v-subheader v-if="item.heading">
                {{ item.heading }}
              </v-subheader>
            </v-col>
            <v-col
                cols="6"
                class="text-center"
            >
              <a
                  href="#!"
                  class="body-2 black--text"
              >EDIT</a>
            </v-col>
          </v-row>
          <v-list-group
              v-else-if="item.children"
              :key="item.text"
              v-model="item.model"
              :prepend-icon="item.model ? item.icon : item['icon-alt']"
              append-icon=""
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>
                  {{ item.text }}
                </v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item
                v-for="(child, i) in item.children"
                :key="i"
                link
            >
              <v-list-item-action v-if="child.icon">
                <v-icon>{{ child.icon }}</v-icon>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>
                  {{ child.text }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
          <v-list-item
              v-else
              :key="item.text"
              link
          >
            <v-list-item-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>
                {{ item.text }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
        :clipped-left="$vuetify.breakpoint.lgAndUp"
        app
        color="blue darken-3"
        dark
        height="80%"
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title
          style="width:300px"
          class="pa-3"
      >
        <v-img src="https://www.lider.cl/images/logo.svg" height="50px" width="150px" aspect-ratio="1" />
      </v-toolbar-title>

      <v-spacer/>
      <v-text-field
          flat
          solo-inverted
          hide-details
          prepend-inner-icon="mdi-magnify"
          label="¿Qué estás buscando?"
          class="hidden-sm-and-down align-center"
          dark
          v-model="token"
          @keyup.enter="getProduct()"
      ></v-text-field>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-apps</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-bell</v-icon>
      </v-btn>
      <v-btn
          icon
          large
      >
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container
          class="fill-height mx-8"
          fluid
      >
        <v-row dense>

          <v-col
              cols="3"
              v-for="product in products"
              :key="product.id"
          >
            <Product
              :brand="product.brand"
              :description="product.description"
              :imageLink="product.image"
              :price="product.price"
              :palindrome="isTokenPalindrome"
            />
          </v-col>
          </v-row>
        <v-row
            align="center"
            justify="center"
        >
          <v-btn
              @click="getProductById('1')"
          >Get By ID</v-btn>
          <v-btn
              @click="getProductsByToken('brh')"
          >Get By Token</v-btn>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>


<script>
import Product from '@/components/Product'
import axios from 'axios'

const API_BYID = "/byId/"
const API_BYTOKEN = "/byToken/"
export default {
  name: 'Tech',
  components: {
    Product,
  },
  props: {
    source: String,
  },
  data: () => ({
    dialog: false,
    drawer: null,
    products: [],
    isTokenPalindrome: false,
    token: "",
    cards: [
      { title: 'Pre-fab homes', src: 'https://cdn.vuetifyjs.com/images/cards/house.jpg', flex: 4 },
      { title: 'Favorite road trips', src: 'https://cdn.vuetifyjs.com/images/cards/road.jpg', flex: 4 },
      { title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 4 },
    ],
    items: [
      { icon: 'mdi-printer', text: 'Electrónica' },
      { icon: 'mdi-cellphone', text: 'Telefonía y fotografía' },
      { icon: 'mdi-laptop', text: 'Computación' },
      { icon: 'mdi-speaker', text: 'Electrohogar' },
      { icon: 'mdi-bed', text: 'Dormitorio' },
      { icon: 'mdi-home', text: 'Hogar' },
      { icon: 'mdi-soccer', text: 'Deportes' },
    ],
  }),
  methods: {
    clearProducts() {
      this.products = []
    },
    getProductById(id) {
      this.clearProducts()
      axios
      .get(`${process.env.VUE_APP_API_URL}/api/${API_BYID}/${id}`)
      .then(response => (this.products.push(response.data)));
    },
    getProductsByToken(token) {
      this.clearProducts()
      axios
      .get(`${process.env.VUE_APP_API_URL}/api/${API_BYTOKEN}/${token}`)
      .then(response => (this.products = response.data));
    },
    getProduct() {
      const t = this.token
      console.log("Search performed")
      this.isTokenPalindrome = this.isPalindrome(t)
      console.log("Palindrome: " + this.isTokenPalindrome)
      if (!this.checkInput(t)) {
        return
      }
      if (!isNaN(t)) {
        this.getProductById(t)
      } else {
        this.getProductsByToken(t)
      }
    },
    checkInput(input) {
      const i = input.trim()
      if (i.length === 0 || i === "0" || i < 1) {
        return false
      }
      return true
    },
    isPalindrome(x) {
      if (x.trim().length <= 1 ) {
        return false
      }
      const xreversed = x.split("").reverse().join("")
      return x === xreversed;
    },
  },
  mounted() {
    
  }
};
</script>
