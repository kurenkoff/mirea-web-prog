<template>
   <div class="card">
       <div class="player">
           Name: {{player.id}}<br/>
           Position: {{player.position}}<br/>
           Nationality: {{player.nationality}}<br/>
       </div>
       <input type="checkbox" @change="sendChange">
   </div>
</template>

<script>
    export default {
        name: 'Player',

        props: {
            id: String,
            player: Object
        },

        mounted() {
            this.player = this.$store.getters.onePlayer(this.id)
        },

        methods: {
          sendChange(){
              this.$store.dispatch('updateForDelete', this.id)
          }
        }
    }
</script>

<style scoped>
.card{
    border-style: solid;
    border-width: 2px;
    text-align: center;
}
</style>