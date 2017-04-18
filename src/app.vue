<template>
  <div>
    <nav class="navbar navbar-toggleable-md navbar-inverse bg-inverse fixed-top">
      <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbar" aria-controls="navbar" aria-expanded="false" aria-label="Menu">
        <span class="navbar-toggler-icon"></span>
      </button>

      <router-link to="/" exact class="navbar-brand" active-class="active">
        Umschlag
      </router-link>

      <div class="collapse navbar-collapse" id="navbar">
        <ul class="navbar-nav">
          <router-link tag="li" to="/" exact class="nav-item" active-class="active">
            <a class="nav-link">
              <i class="fa fa-home" aria-hidden="true"></i>
              Dashboard
            </a>
          </router-link>
          <router-link tag="li" to="/profile" exact class="nav-item" active-class="active">
            <a class="nav-link">
              <i class="fa fa-user" aria-hidden="true"></i>
              Profile
            </a>
          </router-link>
        </ul>
      </div>
    </nav>
    <div class="container-fluid">
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
