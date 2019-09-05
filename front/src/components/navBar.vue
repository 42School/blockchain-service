<template>
  <nav class="navbar is-fixed-top" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <router-link class="navbar-item" to="/">
        <img src="../assets/logo_42_blockchain.svg" alt="logo_42_blockchain" id="logo"/>
      </router-link>

      <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="navbarBasicExample" class="navbar-menu">
      <div class="navbar-start">
        <div class="buttons">
          <router-link class="button is-rounded is-medium is-success" to="/">
            Home
          </router-link>
        </div>
      </div>
    </div>

    <!--    Starting: Change Button if one user is connected !!    -->
    <div class="navbar-end" v-if="this.$session.exists() === false || userConnect === false">
      <div class="navbar-item">
        <div class="buttons">
          <button class="button is-rounded is-medium is-success" @click="showModal('SignUpForm')">
            <strong>Rechercher un.e étudiant.e</strong>
          </button>
          <a class="button is-rounded is-medium is-light" @click="showModal('SignInForm')">
            Student
          </a>
        </div>
      </div>
      <Modal :isShowArg="modalIsShow" :childArg="modalContentName" @changeModalContentName="modalContentName = $event" @changeModalIsShow="modalIsShow = $event"></Modal>
    </div>

    <div class="navbar-end" v-else-if="this.$session.exists() === true || userConnect === true">
      <div class="navbar-item">
        <div class="buttons">

          <div class="navbar-item has-dropdown is-hoverable is-hidden-touch">
            <button class="button is-rounded is-medium" :class="[notifsRead ? 'is-success' : 'is-danger']">
              <font-awesome-icon icon="bell" size="1x"/>
            </button>
            <div class="navbar-dropdown is-right is-hidden-mobile" id="div-notifs">
              <div v-for="(notif, index) in notifs" class="notification" :class="notif.vu ? 'white' : 'is-danger'"  @click="seeNotif(notif.id, index)">
                <a :href="notif.link">{{ notif.notif }}</a>
              </div>
            </div>
          </div>
          <div class="navbar-item is-hidden-desktop">
            <router-link to="/notifs-mobile/" class="button is-rounded" :class="[notifsRead ? 'is-success' : 'is-danger']">
              <font-awesome-icon icon="bell" size="1x"/>
            </router-link>
          </div>

          <div class="is-mobile">

          </div>

          <router-link class="button is-rounded is-medium is-success" :to="'/profile/' + this.$session.get('username')">
            <strong>{{ this.$session.get('username') }}</strong>
          </router-link>

          <button class="button is-rounded is-medium is-danger" @click="signOut">
            <font-awesome-icon icon="sign-out-alt" size="1x" color="white"/>
          </button>
        </div>
      </div>
    </div>
    <!--    End: Change Button if one user is connected !!   -->

  </nav>
</template>

<script>
    import Modal from '../Modals/modal'
    import axios from 'axios'

    export default {
        name: 'NavBar',
        data () {
            return {
                userConnect: null,
                modalIsShow: false,
                modalContentName: null,
                isShowNotification: false,
                notifs: null,
                notifsRead: true,
                notifsMsgRead: true,
                nbNotif: null
            }
        },
        components: {
            Modal
        },
        methods: {
            checkNotifRead () {
                this.notifsRead = true
                for (let i = 0; i < this.notifs.length; i++) {
                    if (this.notifs[i].vu === 0) {
                        this.notifsRead = false
                        break
                    }
                }
            },
            seeNotif (id, index) {
                axios.put('/users-notifs/set-see/' + id, {
                    data: {
                        username: this.$session.get('username'),
                    }
                })
                    .then((response) => {
                        if (response.status === 200) {
                            this.notifs[index].vu = 1
                            this.checkNotifRead()
                        }
                    })
            },
            getNotifs () {
                if (this.$session.get('username')) {
                    axios.get('/users-notifs/get-notifs/' + this.$session.get('username'))
                        .then((response) => {
                            if (response.status === 200) {
                                this.notifs = response.data
                                this.checkNotifRead()
                            }
                        })
                }
            },
            /**
             *  showModal show the Modal
             *
             * @param ContentName it's the name of component child that the Modal must choose (ref in Modal.vue)
             */
            showModal (ContentName) {
                this.modalIsShow = true
                this.modalContentName = ContentName
            },

            /**
             *  signOut function destroy all cookies of Session for logout user, and return to index !
             */
            signOut () {
                axios.get('/users-token/disconnect')
                    .then(res => {
                        if (res.status === 200) {
                            this.$session.destroy()
                            this.userConnect = this.$session.exists()
                            this.$router.push('/')
                        }
                    })
            }
        },
        updated () {
            if (this.$session.exists()) {
                setTimeout(() => {
                    this.getNotifs()
                }, 5000)
            }
        },
        created () {


            document.addEventListener('DOMContentLoaded', () => {

                // Get all "navbar-burger" elements
                const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

                // Check if there are any navbar burgers
                if ($navbarBurgers.length > 0) {

                    // Add a click event on each of them
                    $navbarBurgers.forEach( el => {
                        el.addEventListener('click', () => {

                            // Get the target from the "data-target" attribute
                            const target = el.dataset.target;
                            const $target = document.getElementById(target);

                            // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
                            el.classList.toggle('is-active');
                            $target.classList.toggle('is-active');

                        });
                    });
                }

            });

            if (this.$session.exists()) {
                this.getNotifs()
            }

        },
        watch: {
            '$route' () {
                if (this.$session.exists()) {
                    this.getNotifs()
                }

            }
        }
    }
</script>

<style>
  /* #class-new-notif {
    background: red;
    border-radius: 50%;
    height: 20px;
    width: 20px;
    font-size: small;
    position: relative;
    top: -7px;
    left: -2px;
  } */

  #div-notifs {
    height: 50vh;
    width: 21vh;
    overflow: auto;
  }

  #logo {
    min-height: 70px;
    min-width: 70px;
  }

</style>
