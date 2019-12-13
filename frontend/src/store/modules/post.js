export default {
    actions: {
        async fetchPlayers({commit}) {
            let res = await fetch(
                "/api/get"
            )
            let players = await res.json()
            commit('updatePlayers', players)
        },
        updateForDelete({commit}, id) {
            commit('updateForDelete', id)
        },
        async deletePlayers({commit, state}) {
            let res = await fetch(
                "/api/delete/" + state.forDelete.join(','),
                {
                    method: "POST",
                    mode: 'no-cors', // no-cors, cors, *same-origin
                }
            )
            // eslint-disable-next-line no-console
            console.log(res.status)

            commit('deletePlayers')
        }
    },

    mutations: {
        updatePlayers(state, players) {
            state.players = players
        },
        updateForDelete(state, id) {
            let index = state.forDelete.indexOf(id)
            if (index > -1) {
                state.forDelete.splice(index, 1)
            } else {
                state.forDelete.push(id)
            }
            // eslint-disable-next-line no-console
            console.log(state.forDelete, id, index)
        },
        deletePlayers(state) {
            for (let i = 0; i < state.forDelete.length; i++){
                let index = state.players.findIndex(element => element.id === state.forDelete[i])
                state.players.splice(index, 1)

            }
            state.forDelete = []
        }
    },

    state: {
        players: [],
        forDelete: []
    },

    getters: {
        allPlayers(state) {
            return state.players
        },
        onePlayer(state) {
            return function (id) {
                return state.players.find(player => player.id == id)
            }
        },

    }


}