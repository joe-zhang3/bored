<template>
    <v-app>
      <login></login>
      <app-navigation></app-navigation>
        <v-main >
          <v-container fluid>
          <router-view />
          </v-container>
        </v-main>
        <v-footer app>
        </v-footer>

    </v-app>
</template>

<script>
import AppNavigation from './components/AppNavigation';
import login from './components/Login';
export default {
    name: 'App',
    components: {
      AppNavigation,
      login
    } 
};
</script>

<style>
</style>