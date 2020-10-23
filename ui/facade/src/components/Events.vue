<template>
<v-data-table
    :headers="headers"
    :items="desserts"
    :items-per-page="5"
    class="elevation-1"
  ></v-data-table>
</template>

<script>
export default {
    name: 'events',
    data() {
        return {
            headers: [
                { text: 'NAME', value: 'name'},
                { text: 'ID', value: 'id'},
                { text : 'DESCRIPTION', value: 'description'}
            ],
            desserts: [
                {
                    name: 'joe',
                    id: '1',
                    description: 'This is the description'
                },
                {
                    name: 'joe',
                    id: '2',
                    description: 'This is the description'
                },
                {
                    name: 'joe',
                    id: '3',
                    description: 'This is the description'
                },
                {
                    name: 'joe',
                    id: '4',
                    description: 'This is the description'
                }
            ]
        }
    }
};
</script>

<style >
</style>