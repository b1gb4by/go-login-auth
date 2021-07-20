<template>
  <nav class="navbar navbar-expand-md navbar-dark bg-dark">
      <ul class="navbar-nav mr-auto">
          <li class="nav-item">
              <router-link to="/" class="nav-link">Home</router-link>
          </li>
      </ul>

      <ul class="navbar-nav my-2 my-lg-0">
          <li class="nav-item">
              <router-link to="/login" class="nav-link">Login</router-link>
          </li>
          <li class="nav-item">
              <router-link to="/register" class="nav-link">Register</router-link>
          </li>       
      </ul>
  </nav>
</template>

<script>
export default {
    name: 'Nav'
}
</script>
