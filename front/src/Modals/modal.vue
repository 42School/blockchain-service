<template>
  <div class="modal" :class="{'is-active': isShowArg}">
    <div class="modal-background" @click="closeModal"></div>
    <div class="modal-card">
      <section class="modal-card-body">
        <components :is="childArg" @changeChild="changeChildArg"></components>
      </section>
    </div>
  </div>
</template>

<script>
    import SignInForm from './signIn'
    import SignUpForm from './signUp'

    export default {
        name: 'ModalTemplate',
        props: {
            isShowArg: false,
            childArg: null
        },
        components: {
            /**
             *  It's all component name that we can choose by the var `childArg`
             */
            SignUpForm,
            SignInForm,
        },
        methods: {
            /**
             *  closeModal $emit a new value for changing var `isShow` passing in props
             */
            closeModal() {
                this.$emit('changeModalIsShow', !this.isShowArg)
            },

            /**
             *  changeChildArg $emit a new value for changing var `childArg` passing in props
             *
             *  @param nameChild is the value $emit by the child of this Modal and send in $emit
             */
            changeChildArg(nameChild) {
                this.$emit('changeModalContentName', nameChild)
            }
        }
    }
</script>

<style>
  .modal-card-body {
    border-radius: 20px;
  }
</style>

