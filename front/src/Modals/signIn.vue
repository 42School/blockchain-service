<template>
  <div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Username :</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input class="input is-fullwidth is-rounded" :class="{'is-danger': errors.has('username')}" type="text" name="username" placeholder="username" v-model="username" v-validate="'required'">
          </p>
        </div>
      </div>
      <span class="error">{{ errors.first('username') }}</span>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Password :</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input class="input is-fullwidth is-rounded" :class="{'is-danger': errors.has('password')}" name="password" type="password" placeholder="Password" ref="password" v-validate="'required'" v-model="password">
          </p>
        </div>
      </div>
      <span class="error">{{ errors.first('password') }}</span>
    </div>

    <span class="error">{{ error_msg }}</span>

    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <button class="button is-success is-rounded" @click="sendSignIn">Sign In</button>
      </div>
      <div class="control">
        <button class="button is-rounded" @click="changeModalChild">Reset your password</button>
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
        name: 'SignInForm',
        data () {
            return {
                username: null,
                password: null,
                localization: null,
                error_msg: null
            }
        },
        methods: {
            /**
             *  changeModalChild change the content of modal !
             */
            changeModalChild () {
                this.$parent.changeChildArg('ResetPasswordForm')
            },

            getAndPostLocalization () {
                navigator.geolocation.getCurrentPosition((pos) => {
                    this.localization = `latitude:${pos.coords.latitude}, longitude:${pos.coords.longitude}`
                    axios.put('/users-infos/update-loc/' + this.$session.get('username'), {
                        data: {
                            localization: this.localization
                        }
                    })},
                )
            },

            /**
             *  sendSignIn request if the password it's same of password link at this user in db !
             *
             *  If the password and the db password it's same password Api generate a random unique token,
             *  and return this token, saved in the cookies session with this username !
             */
            sendSignIn () {
                let connectData = {
                    'username': this.username,
                    'password': this.password
                }
                axios.post('/users/connect/' + this.username, {
                    data: connectData
                })
                    .then((response) => {
                        if (response.status === 201 && 'token' in response.data) {
                            this.$parent.closeModal()
                            this.$session.start()
                            this.$session.set('username', this.username)
                            this.$session.set('token', response.data.token)
                            axios.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`
                            this.getAndPostLocalization()
                        } else if (response.status === 200 && 'error_message' in response.data) {
                            if (response.data.error_message === 'Failure !! Please Confirme your email !') {
                                this.error_msg = 'Confirmez votre email avant de vous connecter.'
                            } else if (response.data.error_message === 'Failure !! Password invalid.') {
                                this.error_msg = 'Username ou mot de passe invalides.'
                            }
                        }
                    })
            }
        }
    }
</script>
