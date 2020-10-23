<template>
    <v-app>
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
export default {
    name: 'App',
    components: {
      AppNavigation
    } 
};
</script>

<style>
</style>