<template>
    <div class="panel panel-info">
        <div class="panel-heading">手札：</div>
        <div class="panel-body">
            <div v-for="(card, index) in this.$store.getters.hands" class="btn-group"  data-toggle="buttons">
                <card :name="card.name" :desc="card.desc" :type="card.type" 
                      :cost="card.cost" :key="index" :value="index" @trigger="click">
                </card>
            </div>
            <button v-show="this.$store.getters.showNoActionButton" 
                    class="btn btn-default"  @click="noAction"> アクションカードを使用しない</button>
            <button v-show="this.$store.getters.showBuyButton" 
                    class="btn btn-default"  @click="buy"> 購入する</button>
            <button v-show="this.$store.getters.showNoBuyButton" 
                    class="btn btn-default"  @click="noBuy"> 購入しない</button>
            <selection v-show="this.$store.getters.showBuySelection" @trigger="touch"></selection>
        </div>
    </div>
</template>
<script>
import Card from './Card.vue'
import Selection from './Selection.vue'
export default {
    components: { Card,Selection },
    methods: {
        click: function(value){
            switch(this.$store.getters.phase){
                case 'Action':
                    this.$store.dispatch('isActionCard', value);
                    break;
                case 'Buy':
                    this.$store.commit('appearBuyButton');
                    this.$store.commit('addChecks', value);
                    break;
                case 'Chapel':
                    this.$store.dispatch('tryChapel', value);
                    break;
            }
        },
        noAction: function(){
            this.$store.dispatch('startBuyPhase');
            this.$store.commit('disappearNoActionButton');
        },
        noBuy: function(){
            this.$store.dispatch('startCleanUpPhase');
            this.$store.commit('disappearNoBuyButton');
        },
        buy: function(){
            this.$store.dispatch('buy');
        },
        touch: function(selected){
            this.$store.commit('appearBuyButton');
            this.$store.commit('selectOption', selected);
        },
    }
}
</script>

<style scoped>
</style>
