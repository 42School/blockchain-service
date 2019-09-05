<template>
  <div id="app">
    <NavBar></NavBar>
    <router-view/>
    <Footer></Footer>
  </div>
</template>

<script>
import NavBar from './components/navBar'
import Footer from './components/footer'

export default {
  name: 'App',
  components: {
    NavBar,
    Footer
  }
}
</script>

<style>
@font-face {
  font-family: 'Gill Sans Medium';
  font-style: normal;
  font-weight: normal;
  src: url('assets/fonts/gill-sans/Gill_Sans_Medium.woff') format('woff');
}

@font-face {
  font-family: 'Etna';
  font-style: normal;
  font-weight: normal;
  src: url('assets/fonts/etna-free-font.otf') format('otf');
}

body {
  scrollbar-width: none;
}

body::-webkit-scrollbar {
  display: none;
}
</style>
