<template>
  <div class="app-background">
    <header>
      <img alt="One-for-all" class="logo" src="./assets/logo.svg" />
      <nav>
        <RouterLink to="/"></RouterLink>
      </nav>
    </header>

    <main>
      <!-- <div class="login-box"> -->
        <RouterView /> 
      <!-- </div> -->
      
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
</script>

<style scoped>
* {
  margin: 0; /* Bỏ margin cho tất cả phần tử */
  padding: 0; /* Bỏ padding cho tất cả phần tử */
  box-sizing: border-box; /* Đảm bảo padding không làm tăng kích thước phần tử */
}

.app-background {
  background: url('./assets/background/bg_login.jpg') no-repeat center center fixed; 
  background-size: cover; /* Để hình nền phủ kín */
  min-height: 100vh; /* Để chiều cao tối thiểu là 100% chiều cao của viewport */
  display: flex; /* Dùng flexbox để căn giữa */
  flex-direction: column; /* Căn giữa theo chiều dọc */
  /* color: pink; */
}

header {
  text-align: center;
  margin-bottom: 2rem; /* Bạn có thể điều chỉnh khoảng cách này nếu cần */
}

nav {
  margin-top: 1rem; /* Bạn có thể điều chỉnh khoảng cách này nếu cần */
}

nav a {
  margin: 0 10px;
  text-decoration: none;
}

.logo {
  display: block;
  margin: 0 auto;
  max-width: 150px; /* Logo width */
}

main {
  flex: 1; /* Fill remaining space */
  display: flex; /* Flexbox for centering */
  justify-content: center; /* Center horizontally */
  align-items: flex-start; /* Align items to the start */
  padding-top: 10px; /* Optional: control space above the login box */
  margin-top: -150px;
}

/* Add styles for your login box */
.login-box {
  width: 100%;
  max-width: 400px; /* Adjust the max width as necessary */
  background-color: white; /* Background color of the login box */
  padding: 2rem; /* Padding inside the login box */
  border-radius: 10px; /* Rounded corners */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); /* Shadow effect */
  margin-top: 10px; /* Adjust space from the header */
}
</style>
