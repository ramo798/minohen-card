<template>
  <div class="home">
<!--    <img alt="Vue logo" src="../assets/logo.png">-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
      <GetComp></GetComp>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import GetComp from "../components/GetComp";

export default {
  name: 'home',
  components: {
    // HelloWorld,
      GetComp
  }
}
</script>
