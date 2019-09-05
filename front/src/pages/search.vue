<template>
  <section class="particles full-height">
    <div class="card overlay card-search">
      <progress class="progress is-small is-success" max="100" v-if="signature === null"></progress>
      <div class="information" v-else>
        <p class="title has-text-centered">Cette étudiant.e est certifié par 42</p>
        <br>
        <p class="subtitle has-text-centered">Le hash de son diplôme:</p>
        <p class="has-text-centered">{{ login }}</p>
        <br>
        <p class="has-text-centered">{{ signature }}</p>
        <br>
        <div class="field is-grouped is-grouped-centered">
          <button v-if="checkerHash === false" class="button is-success is-rounded" :class="(loading == true) ? 'is-loading' : ''" @click="checkHash">Verifier le hash</button>
          <button v-else class="button is-success is-rounded" @click="closeCheckHash">Close</button>
        </div>

        <div class="check-hash" v-if="checkerHash === true">
          <div id="sign1Div" v-if="signData1 != null">
            <p class="subtitle is-6 has-text-centered has-text-white">Première signature by <br>{{ signData1.address }} :</p>
            <p class="is-italic is-size-8 has-text-centered">{{ signData1.signature }}</p>
          </div>
          <br>
          <div id="signHDiv" v-if="hashChecker != null">
            <p class="subtitle is-6 has-text-centered has-text-white">Signatures en Blockchain:</p>
            <p class="is-italic is-size-7 has-text-centered" :class="(signature != hashChecker) ? 'has-text-danger' : 'has-text-primary'">{{ hashChecker }}</p>
          </div>
        </div>

        <br>
        <p class="subtitle has-text-centered">D'autres informations concernant son cursus:</p>
        <p class="has-text-centered">Level: {{ data.intraLevel }}</p>
        <p class="has-text-centered">Promo: {{ data.promoYears }}</p>
        <p class="has-text-centered">Date de certification: {{ data.graduateYears }}</p>
      </div>
    </div>
    <vue-particles color="#dedede" shapeType="polygon" class="full-height"></vue-particles>
  </section>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'Search',
    data () {
      return {
        checkerHash: false,
        loading: false,
        login: null,
        signature: null,
        data: null,
        signData1: null,
        signData2: null,
        signData3: null,
        hashChecker: null
      }
    },
    methods: {

      async checkHash () {
        this.loading = true;
        this.checkerHash = true;
        await axios.get('/sign/sophie/' + this.login)
          .then((response) => {
            console.log(response);
            if (response.data.message === 'Success') {
              this.signData1 = response.data.signData;
              console.log(this.signData1);
            }
          })
        await axios.get('/sign/benny/' + this.login)
          .then((response) => {
            if (response.data.message === 'Success') {
              this.signData2 = response.data.signData;
            }
          })
        await axios.get('/sign/niel/' + this.login)
          .then((response) => {
            if (response.data.message === 'Success') {
              this.signData3 = response.data.signData;
            }
          })
        await axios.post('/sign/get-hash/'  + this.login, {
          data: {
            signData1: this.signData1,
            signData2: this.signData2,
            signData3: this.signData3,
          }
        }).then((response) => {
          if (response.data.message === 'Success') {
            this.hashChecker = response.data.hash;
          }
        })
        this.loading = false;
      },

      closeCheckHash () {
        this.checkerHash = false;
      }

    },
    created () {
      if (this.$route.params.login === undefined) {
        this.$router.push('/');
      } else {
        this.login = this.$route.params.login;
        axios.get('/get/' + this.login)
          .then((response) => {
            if (response.data.message === 'Success') {
              this.signature = response.data.signature;
              this.data = response.data.data;
            }
          })
          .catch(() => {
            this.$router.push('/');
          })
      }
    }
  }
</script>

<style>
  .card-search {
    scrollbar-width: none;
    overflow-x: hidden;
    overflow-y: visible;
    min-width: 45vh;
    max-height: 45vh;
    max-width: 45vh;
    margin: -22.5vh 0 0 -22.5vh;
  }

  .card-search::-webkit-scrollbar {
    display: none;
  }
</style>
