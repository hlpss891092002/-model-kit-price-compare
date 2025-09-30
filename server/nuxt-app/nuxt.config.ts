// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  nitro: {
    cors: {
      origin: '*'
    }
  },

  runtimeConfig: {
    apiBaseUrl: process.env.API_BASE_URL || 'http://localhost:8000',
    public: {
      apiBaseUrl: process.env.PUBLIC_API_URL || 'http://localhost:8000'
    }
  },

  ssr: true,

  modules: [],

  app: {
    head: {
      title: 'Kotobukiya Price Comparison',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Compare Kotobukiya product prices across different stores' }
      ]
    }
  }
})
