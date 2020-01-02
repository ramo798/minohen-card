<template>
    <div class="getcomp">
        <div style="border: 1px solid;display: inline-block" v-on:click="show">入力画面</div>
        <div v-if="form" class="form-box">
            <span>facebook_id:</span>
            <input v-model="facebook_id">
            <br>
            <span>ニックネーム:</span>
            <input v-model="nickname">
            <br>
            <span>入会年:</span>
            <input v-model="year">
            <br>
            <span>入会月:</span>
            <input v-model="month">
            <br>
            <span>twitter_id:</span>
            <input v-model="twitter_id">
            <br>
            <span>所属チーム1:</span>
            <input v-model="team_1">
            <br>
            <span>所属チーム2:</span>
            <input v-model="team_2">
            <br>
            <span>mgram1:</span>
            <input v-model="mgram_1">
            <br>
            <span>mgram2:</span>
            <input v-model="mgram_2">
            <br>
            <span>エリア:</span>
            <input v-model="area">
            <br>
            <span>一言:</span>
            <input v-model="word">
            <br>
            <button v-on:click="send()">送信</button>
        </div>

        <div class="space" style="padding-top: 50px" v-for="data in info" :key="data.name">
            <div class="mycard">
                <div class="">
                    <div class="top">
                        <h1>ニックネーム:{{data.nickname}}</h1>
                    </div>
                    <div class="under">
                        <div class="flex">
                            <div class="left">

                                <div class="prof-pic">
                                    <img class="img" src="../assets/noimage.png" width="100px"/>
                                </div>

                            </div>
                            <div class="right">
                                <p>入会:{{data.year}}年{{data.month}}日</p>
                                <p>チーム:{{data.team_1}}</p>
                                <p>一言:{{data.word}}</p>
                            </div>
                        </div>

                    </div>
                </div>

            </div>

            <!--            <div class="card"  v-for="data in info" :key="data.name">-->
            <!--                <p>facebook_id:{{data.facebook_id}}</p>-->
            <!--                <p>ニックネーム:{{data.nickname}}</p>-->
            <!--                <p>入会年:{{data.year}}</p>-->
            <!--                <p>入会月:{{data.month}}</p>-->
            <!--                <p>twitter_id:{{data.twitter_id}}</p>-->
            <!--                <p>所属チーム1:{{data.team_1}}</p>-->
            <!--                <p>所属チーム1:{{data.team_2}}</p>-->
            <!--                <p>一言:{{data.word}}</p>-->
            <!--                <p>mgram_1:{{data.mgram_1}}</p>-->
            <!--                <p>mgram_2:{{data.mgram_2}}</p>-->
            <!--                <p>mgram_3:{{data.mgram_3}}</p>-->
            <!--                <p>mgram_4:{{data.mgram_4}}</p>-->
            <!--                <p>mgram_5:{{data.mgram_5}}</p>-->
            <!--                <p>mgram_6:{{data.mgram_6}}</p>-->
            <!--                <p>mgram_7:{{data.mgram_7}}</p>-->
            <!--                <p>mgram_8:{{data.mgram_8}}</p>-->
            <!--                <p>mgram_9:{{data.mgram_9}}</p>-->
            <!--                <p>エリア:{{data.area}}</p>-->
            <!--                <p>card_color:{{data.card_color}}</p>-->
            <!--            </div>-->
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    export default {

        name: "GetComp",
        data () {
            return {
                mes:"",
                info: null,
                form:false,
                facebook_id:"",
                nickname:'',
                year:0,
                month:0,
                twitter_id:"",
                team_1:"",
                team_2:"",
                word:"",
                mgram_1:"",
                mgram_2:"",
                area:"",
            }
        },
        mounted () {
            axios.get('http://ec2-13-115-220-61.ap-northeast-1.compute.amazonaws.com:8080/card/all')
                .then(response => (this.info = response.data))
        },
        methods: {
            show: function(){
                if (this.form){
                    this.form = false;
                }else {
                    this.form = true;
                }

            },
            serch: function () {
                axios.get('http://ec2-13-115-220-61.ap-northeast-1.compute.amazonaws.com:8080/card/all')
                    .then(response => (this.info = response.data))
            },
            send:function () {
                axios.post('http://ec2-13-115-220-61.ap-northeast-1.compute.amazonaws.com:8080/register', {
                    facebook_id: this.facebook_id,
                    nickname: this.nickname,
                    year:this.year,
                    month:this.month,
                    twitter_id:this.twitter_id,
                    team_1:this.team_1,
                    team_2:this.team_2,
                    word:this.word,
                    mgram_1:this.mgram_1,
                    mgram_2:this.mgram_2,
                    area:this.area,
                }).then(res => {
                    // console.log(res.status);
                    this.mes =res.status;
                    this.serch();
                    this.facebook_id="";
                    this.nickname='';
                    this.year=0;
                    this.month=0;
                    this.twitter_id="";
                    this.team_1="";
                    this.team_2="";
                    this.word="";
                    this.mgram_1="";
                    this.mgram_2="";
                    this.area="";
                });
            }

        }
    }
</script>

<style scoped>
    .flex{
        display: flex;
        /*flex-direction: column;*/
    }
    .jc{
        justify-content: center;
        align-items: center;
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
    .mycard{
        padding: 1%;
    }
    .prof-pic{
        width: 100px;
        height: 100px;
        background-color: antiquewhite;
    }
    .left{
        padding: 2%;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .right {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
</style>