<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
  <div>
    <header>
      <img alt="Vue logo" class="logo" />
      <nav>
        <RouterLink to="/">Login</RouterLink>
        <RouterLink to="/register">Register</RouterLink>
        <RouterLink to="/dashboard">Dashboard</RouterLink>
      </nav>
    </header>

    <RouterView /> <!-- This is where routed components will be displayed -->
  </div>
</template>

<style scoped>
header {
  text-align: center;
  margin-bottom: 2rem;
}

nav {
  margin-top: 1rem;
}

nav a {
  margin: 0 10px;
  text-decoration: none;
}

.logo {
  display: block;
  margin: 0 auto;
}
</style>
