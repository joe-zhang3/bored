<template>
  <div id="app">
    <AppNavigation />
      <router-view />
  </div>
</template>

<script>
import AppNavigation from "./components/AppNavigation";
export default {
  name: "App",
  components: {
    AppNavigation
  },
};
</script>
