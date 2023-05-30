<template>
 <div>
    <b-table striped hover :items="desserts"></b-table>
  </div> 
</template>

<script>
// import axios from "axios";
export default {
  name: "events",
  data() {
    return {
      headers: [
        { text: "NAME", value: "title" },
        { text: "ID", value: "id" },
        { text: "DESCRIPTION", value: "description" },
      ],
      desserts: [{"title": "Joe", "id": "12345", "description": "This is a test."}],
    };
  },
  mounted() {
    // axios.get("/events").then((response) => {
    //   this.desserts = response.data;
    // });
  },
};
</script>

<style >
</style>