<template>
    <div class="getcomp">
        <div class="flex form-box">
            <p>your name:</p>
            <input v-model="msg">
            <button v-on:click="serch(msg)">Greet</button>
        </div>

        <div class="flex space">
            <div class="card"  v-for="data in info" :key="data.name">
                <p>名前:{{data.Name}}</p>
                <p>ニックネーム:{{data.Nickname}}</p>
                <p>Mgram:{{data.Mgram}}</p>
                <p>Twitter:{{data.TwitterId}}</p>
                <p>所属チーム:{{data.BlongTeam}}</p>
                <p>住み:{{data.Where}}</p>
                <p>一言:{{data.Word}}</p>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    export default {

        name: "GetComp",
        data () {
            return {
                info: null,
                msg:"",
            }
        },
        // mounted () {
        //     axios
        //         .get('http://localhost:80/test')
        //         .then(response => (this.info = response.data))
        // },
        methods: {
            greet: function (event) {
                this.msg = "anker"
                // `event` は、ネイティブ DOM イベントです
                if (event) {
                    alert(event.target.tagName)
                }
            },
            serch: function () {
                // alert('Hello ' + this.msg + '!')
                axios.get('http://localhost:80/test')
                    .then(response => (this.info = response.data))
                this.msg = ""
            }
        }
    }
</script>

<style scoped>
    .flex{
        display: flex;
        /*flex-direction: column;*/
    }
    .dlex p{
        display: flex;
    }
    .space{
        justify-content: space-evenly;
    }
    .card{
        border: 1px solid;
    }
    .form-box{
        padding-bottom: 10%;
        justify-content: center;
    }

</style>