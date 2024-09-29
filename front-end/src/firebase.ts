import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDQvty7BnGLzc0qMfCfjZL3WV9eCaN5mU4",
  authDomain: "website-9f10c.firebaseapp.com",
  projectId: "website-9f10c",
  storageBucket: "website-9f10c.appspot.com",
  messagingSenderId: "485069253009",
  appId: "1:485069253009:web:786c43ea3075eba9e31230",
  measurementId: "G-82GFW2CNGQ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

const auth = getAuth(app);

export { auth };
// const analytics = getAnalytics(app);