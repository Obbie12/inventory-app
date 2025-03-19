import { defineStore } from 'pinia'
import { ref } from 'vue'

const API_BASE = import.meta.env.VITE_API_BASE_URL;

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const token = ref(localStorage.getItem('authToken') || null)

  const login = async (credentials) => {
    const response = await fetch(`${API_BASE}/api/v1/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credentials)
    })
    const data = await response.json()
    user.value = data.user
    token.value = data.token
    localStorage.setItem('authToken', data.token)
  }

  const register = async (userData) => {
    const response = await fetch(`${API_BASE}/api/v1/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(userData)
    })
    return await response.json()
  }

  const logout = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('authToken')
  }

  return { user, token, login, register, logout }
})