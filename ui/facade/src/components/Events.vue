<template>
<v-data-table
    :headers="headers"
    :items="desserts"
    :items-per-page="5"
    class="elevation-1"
  ></v-data-table>
</template>

<script>
import axios from 'axios'
export default {
    name: 'events',
    data() {
        return {
            headers: [
                { text: 'NAME', value: 'title'},
                { text: 'ID', value: 'id'},
                { text : 'DESCRIPTION', value: 'description'}
            ],
            desserts: []
        }
    },
    mounted () {
        axios
        .get('http://localhost:8090/events')
        .then(response => {
            this.desserts = response.data
        })
    }
};
</script>

<style >
</style>