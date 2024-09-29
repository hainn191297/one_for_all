<template>
    <div class="dashboard">
      <h1>Welcome to the Dashboard</h1>
      <p>You are successfully logged in!</p>
      <button @click="logout">Logout</button>
    </div>
  </template>
  
  <script setup>
  import { getAuth, signOut } from 'firebase/auth';
  import { useRouter } from 'vue-router';
  
  const router = useRouter();
  const auth = getAuth();
  
  const logout = async () => {
    try {
      await signOut(auth);
      router.push('/login');
    } catch (error) {
      console.error("Logout Error: ", error.message);
    }
  };
  </script>
  
  <style scoped>
  .dashboard {
    text-align: center;
    margin-top: 50px;
  }
  button {
    margin-top: 20px;
    padding: 10px 20px;
    font-size: 16px;
  }
  </style>
  