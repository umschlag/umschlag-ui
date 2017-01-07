import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
  count: 0
}

const mutations = {
  increment (currState) {
    currState.count += 1
  },
  decrement (currState) {
    currState.count -= 1
  }
}

const actions = {
  increment: ({ commit }) => commit('increment'),
  decrement: ({ commit }) => commit('decrement'),
  incrementAsync ({ commit }) {
    return new Promise((resolve) => {
      setTimeout(() => {
        commit('increment')
        resolve()
      }, 1000)
    })
  }
}

const getters = {
  evenOrOdd: currState => (currState.count % 2 === 0 ? 'even' : 'odd')
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
