<template>
    <div class="login-container">
      <h1>Login</h1>
      <form @submit.prevent="loginWithEmail">
        <div>
          <label for="email">Email:</label>
          <input type="email" v-model="email" required />
        </div>
        <div>
          <label for="password">Password:</label>
          <input type="password" v-model="password" required />
        </div>
        <button type="submit">Login</button>
      </form>
      <div class="social-login">
        <button @click="signInWithGoogle">Login with Google</button>
        <button @click="signInWithFacebook">Login with Facebook</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { signInWithEmailAndPassword, GoogleAuthProvider, FacebookAuthProvider, signInWithPopup } from 'firebase/auth';
  import { auth } from '@/firebase'; 
  import { useRouter } from 'vue-router';
  
  const email = ref('');
  const password = ref('');
  const router = useRouter();
  
  const loginWithEmail = async () => {
    try {
      await signInWithEmailAndPassword(auth, email.value, password.value);
      router.push('/dashboard');
    } catch (error) {
      console.error("Login Error: ", error.message);
    }
  };
  
  const signInWithGoogle = async () => {
    const provider = new GoogleAuthProvider();
    try {
      await signInWithPopup(auth, provider);
      router.push('/dashboard');
    } catch (error) {
      console.error("Google Sign-In Error: ", error.message);
    }
  };
  
  const signInWithFacebook = async () => {
    const provider = new FacebookAuthProvider();
    try {
      await signInWithPopup(auth, provider);
      router.push('/dashboard');
    } catch (error) {
      console.error("Facebook Sign-In Error: ", error.message);
    }
  };
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: auto;
    padding: 1rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  h1 {
    text-align: center;
    margin-bottom: 1.5rem;
  }
  
  form {
    display: flex;
    flex-direction: column;
  }
  
  label {
    margin-bottom: 0.5rem;
  }
  
  input {
    margin-bottom: 1rem;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  button {
    padding: 0.5rem;
    border: none;
    border-radius: 4px;
    background-color: #007bff;
    color: white;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  button:hover {
    background-color: #0056b3;
  }
  
  .social-login {
    display: flex;
    justify-content: space-between;
    margin-top: 1rem;
  }
  
  .social-login button {
    background-color: #fff;
    border: 1px solid #ccc;
    color: #333;
  }
  
  .social-login button:hover {
    background-color: #f0f0f0;
  }
  </style>
  