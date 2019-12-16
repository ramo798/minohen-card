<template>
    <div class="getcomp">
        {{ info }}
    </div>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "GetComp",
        data () {
            return {
                info: null
            }
        },
        mounted () {
            axios
                .get('http://localhost:80/test')
                .then(response => (this.info = response))
        }
    }
</script>

<style scoped>

</style>