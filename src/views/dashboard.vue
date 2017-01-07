<template>
  <div>
    <h1>
      Dashboard
    </h1>

    Clicked: {{ $store.state.count }} times, count is {{ evenOrOdd }}.

    <button @click="increment">+</button>
    <button @click="decrement">-</button>
    <button @click="incrementAsync">Increment async</button>
  </div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex'

  export default {
    name: 'dashboard-view',
    components: {

    },
    computed: mapGetters([
      'evenOrOdd'
    ]),
    methods: mapActions([
      'increment',
      'decrement',
      'incrementAsync'
    ])
  }
</script>

<style>
</style>
