<template>
  <div id="app">
    <PlayerList/>
  </div>
</template>

<script>
  import PlayerList from './components/PlayerList.vue'

  export default {
    name: 'app',
    components: {
      PlayerList
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
