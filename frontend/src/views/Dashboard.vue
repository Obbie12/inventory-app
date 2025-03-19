<template>
  <div class="dashboard-container">
    <div class="header">
      <h1>Inventory Management</h1>
      <div class="header-actions">
        <button class="btn-add" @click="showAddModal = true">
          + Add Product
        </button>
        <button class="btn-logout" @click="authStore.logout()">
          Logout
        </button>
      </div>
    </div>

    <AddProductModal 
      :show="showAddModal"
      @submit="handleAddProduct"
      @close="showAddModal = false"
    />

    <ProductList 
      :products="productsStore.products"
      @edit="handleEditProduct"
      @delete="handleDeleteProduct"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useProductsStore } from '../stores/products'
import { useAuthStore } from '../stores/auth'
import ProductList from '../components/ProductList.vue'
import AddProductModal from '../components/AddProductModal.vue'

const productsStore = useProductsStore()
const authStore = useAuthStore()
const showAddModal = ref(false)

onMounted(async () => {
  await productsStore.fetchProducts()
})

const handleAddProduct = async (productData) => {
  try {
    await productsStore.addProduct(productData)
    showAddModal.value = false
  } catch (error) {
    console.error('Error adding product:', error)
  }
}

const handleEditProduct = (product) => {
  // Implement edit logic if needed
}

const handleDeleteProduct = (productId) => {
  productsStore.deleteProduct(productId)
}
</script>

<style scoped>
.dashboard-container {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.header-actions {
  display: flex;
  gap: 1rem;
}

.btn-add {
  background-color: #007bff;
  color: white;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-logout {
  background-color: #dc3545;
  color: white;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-add:hover {
  background-color: #0056b3;
}

.btn-logout:hover {
  background-color: #bb2d3b;
}

@media (max-width: 768px) {
  .dashboard-container {
    padding: 1rem;
  }
  
  .header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .header-actions {
    width: 100%;
    flex-direction: column;
  }
  
  .btn-add,
  .btn-logout {
    width: 100%;
  }
}
</style>