import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useAuthStore } from './auth'

const API_BASE = import.meta.env.VITE_API_BASE_URL;

export const useProductsStore = defineStore('products', () => {
  const products = ref([])
  const authStore = useAuthStore()

  const fetchProducts = async () => {
    const response = await fetch(`${API_BASE}/api/v1/products`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    products.value = await response.json()
  }

  const addProduct = async (product) => {
    await fetch(`${API_BASE}/api/v1/products`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(product)
    })
    await fetchProducts()
  }

  const updateProduct = async (id, updates) => {
    await fetch(`${API_BASE}/api/v1/products/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(updates)
    })
    await fetchProducts()
  }

  const deleteProduct = async (id) => {
    await fetch(`${API_BASE}/api/v1/products/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    await fetchProducts()
  }

  return { products, fetchProducts, addProduct, updateProduct, deleteProduct }
})