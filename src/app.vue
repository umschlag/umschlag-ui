<template>
  <div>
    <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
      <router-link to="/" exact class="navbar-brand" active-class="active">
        Umschlag
      </router-link>

      <ul class="nav navbar-nav float-xs-right">
        <router-link tag="li" to="/" exact class="nav-item" active-class="active">
          <a class="nav-link">
            Dashboard
          </a>
        </router-link>
        <router-link tag="li" to="/profile" exact class="nav-item" active-class="active">
          <a class="nav-link">
            Profile
          </a>
        </router-link>
      </ul>
    </nav>
    <div class="container">
      <router-view keep-alive></router-view>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'app',
    components: {

    }
  }
</script>

<style>
  @import "index.less";
</style>
