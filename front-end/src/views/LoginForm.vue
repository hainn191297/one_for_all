<template>
  <div class="login-register">
    <div class="login-box card">
      <div class="card-block">
        <h1 class="box-title">User Login</h1>
        <form id="loginform" @submit.prevent="loginWithEmail">
          <div class="form-group input-group">
            <label for="email">Email:</label>
            <input
              type="email"
              id="email"
              v-model="email"
              class="form-control"
              required
            />
          </div>
          <div class="form-group input-group">
            <label for="password">Password:</label>
            <input
              type="password"
              id="password"
              v-model="password"
              class="form-control"
              required
            />
          </div>
          <div class="form-group text-center">
            <button
              id="login_button"
              type="submit"
              class="btn btn-primary btn-md btn-block"
            >
              Login
            </button>
          </div>
        </form>

        <div class="extra-buttons text-center">
          <button
            class="btn btn-secondary btn-md btn-block"
            @click="routeToRegister"
          >
            Register
          </button>
          <button
            class="btn btn-link"
            @click="routeToForgetPassword"
          >
            Forgot Password?
          </button>
        </div>

        <div class="social-login text-center">
          <button
            class="btn btn-google btn-block"
            @click="signInWithGoogle"
          >
            Login with Google
          </button>
          <button
            class="btn btn-facebook btn-block"
            @click="signInWithFacebook"
          >
            Login with Facebook
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import {
  signInWithEmailAndPassword,
  GoogleAuthProvider,
  FacebookAuthProvider,
  signInWithPopup
} from 'firebase/auth';
import { auth } from '@/firebase'; // Firebase initialization
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const router = useRouter();

const loginWithEmail = async () => {
  try {
    await signInWithEmailAndPassword(auth, email.value, password.value);
  } catch (error) {
    console.error('Login Error: ', error.message);
  }
};

const signInWithGoogle = async () => {
  const provider = new GoogleAuthProvider();
  try {
    await signInWithPopup(auth, provider);
  } catch (error) {
    console.error('Google Sign-In Error: ', error.message);
  }
};

const signInWithFacebook = async () => {
  const provider = new FacebookAuthProvider();
  try {
    await signInWithPopup(auth, provider);
  } catch (error) {
    console.error('Facebook Sign-In Error: ', error.message);
  }
};

// Navigate to register page
const routeToRegister = () => {
  router.push('/register'); // Assuming you have a registration route set up
};

// Navigate to forgot password page
const routeToForgetPassword = () => {
  router.push('/forgot-password'); // Assuming you have a forgot password route set up
};
</script>

<style scoped>
.login-register {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  min-height: 100vh;
}

.login-box {
  width: 100%;
  max-width: 400px;
  background-color: white; /* Background color of the login box */
  padding: 2rem; /* Padding inside the login box */
  border-radius: 10px; /* Rounded corners */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); /* Shadow effect */
}

.card-block {
  margin: 0;
  padding: 0;
}

.box-title {
  margin-bottom: 1rem; /* Space below title */
}

/* Other styles for form controls */
.form-group {
  margin-bottom: 1rem; /* Space between form groups */
}

input.form-control {
  width: 100%; /* Full width for inputs */
  padding: 10px; /* Padding inside input */
  border: 1px solid #ccc; /* Border styling */
  border-radius: 5px; /* Rounded corners for inputs */
}

.btn {
  padding: 10px 20px; /* Button padding */
  font-size: 16px; /* Button font size */
}

.btn-block {
  width: 100%; /* Full width for buttons */
  margin-top: 10px; /* Space above buttons */
}

.btn-google {
  background-color: #dd4b39;
  color: white;
}

.btn-facebook {
  background-color: #3b5998;
  color: white;
}
</style>