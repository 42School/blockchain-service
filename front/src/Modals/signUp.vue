<template>
  <div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Login :</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input class="input is-fullwidth is-rounded" name="login" type="text" placeholder="Login de l'étudiant.e" v-model="login">
          </p>
        </div>
      </div>
    </div>
    <span class="error">{{ error_post }}</span>

    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <button class="button is-success is-rounded" :class="(loading == true) ? 'is-loading' : ''" @click="sendSearch">Sign Up</button>
      </div>
      <div class="control">
        <button class="button is-rounded" @click="this.$parent.closeModal">Cancel</button>
      </div>
    </div>

  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'SignUpForm',
    data () {
      return {
        loading: false,
        login: null,
        error_post: null,
      }
    },
    methods: {
      sendSearch () {
        this.loading = true;
        axios.get('/get/' + this.login)
          .then((response) => {
            if (response.data.message === 'Success') {
              this.loading = false;
              this.$parent.closeModal();
              this.$router.push('/search/' + this.login);
            }
          })
          .catch((err) => {
            this.loading = false;
            this.error_post = 'Le login de l\'étudiant n\'a pas été trouvé';
          })
        console.log('lol')
      }
    }
  }
</script>
