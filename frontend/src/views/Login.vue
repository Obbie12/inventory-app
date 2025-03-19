<template>
  <div class="login-container">
    <div class="login-box">
      <form @submit.prevent="handleSubmit">
        <h2>{{ isLogin ? 'Login' : 'Register' }}</h2>
        
        <div class="form-group">
          <label>Username</label>
          <input v-model="credentials.username" required>
        </div>

        <div class="form-group">
          <label>Password</label>
          <input v-model="credentials.password" type="password" required>
        </div>

        <div class="form-group" v-if="!isLogin">
          <label>Email</label>
          <input v-model="credentials.email" type="email" required>
        </div>

        <button type="submit" class="primary-btn">{{ isLogin ? 'Login' : 'Register' }}</button>
        <button type="button" class="secondary-btn" @click="isLogin = !isLogin">
          {{ isLogin ? 'Create account' : 'Back to login' }}
        </button>
      </form>
    </div>
  </div>
</template>
  
  <script setup>
  import { ref } from 'vue'
  import { useAuthStore } from '../stores/auth'
  import { useRouter } from 'vue-router'
  
  const authStore = useAuthStore()
  const router = useRouter()
  const isLogin = ref(true)
  const credentials = ref({
    username: '',
    password: '',
    email: ''
  })
  
  const handleSubmit = async () => {
    try {
      if (isLogin.value) {
        await authStore.login(credentials.value)
      } else {
        await authStore.register(credentials.value)
      }
      router.push('/dashboard')
    } catch (error) {
      alert('Authentication failed')
    }
  }
  </script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: white;
}

.login-box {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #333;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #666;
}

input {
  width: 90%;
  margin: 0 auto;
  padding: 0.8rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  display: block;
}

button {
  width: 100%;
  padding: 0.8rem;
  margin-top: 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
}

.primary-btn {
  background-color: #007bff;
  color: white;
}

.primary-btn:hover {
  background-color: #0056b3;
}

.secondary-btn {
  background-color: transparent;
  color: #007bff;
  border: 1px solid #007bff;
  margin-top: 1rem;
}

.secondary-btn:hover {
  background-color: #f0f8ff;
}

@media (max-width: 480px) {
  .login-box {
    margin: 1rem;
    padding: 1.5rem;
  }
}
</style>