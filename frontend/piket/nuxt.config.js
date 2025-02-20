export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'piket',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css' },
    ],
  },

  // Global CSS
  css: [
    'bootstrap/dist/css/bootstrap.min.css',
  ],

  // Plugins to run before rendering page


  // Auto import components
  components: true,

  // Modules for dev and build
  buildModules: [],

  // Axios module configuration
  modules: [
    '@nuxtjs/axios',
    'bootstrap-vue/nuxt',
  ],
  axios: {
    baseURL: 'http://localhost:8080', // Ganti dengan URL backend Anda
  },

  // Build Configuration
  build: {},

  // Router Configuration
  
}
