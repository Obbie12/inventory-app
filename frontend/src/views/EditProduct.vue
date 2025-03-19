<template>
    <div class="edit-product-container">
      <h2>Edit Product</h2>
      
      <form @submit.prevent="handleSubmit">
        <!-- Product Name -->
        <div class="form-group">
          <label>Product Name</label>
          <input v-model="form.product_name" required>
        </div>
  
        <!-- SKU -->
        <div class="form-group">
          <label>SKU</label>
          <input v-model="form.sku" required>
        </div>
  
        <!-- Quantity -->
        <div class="form-group">
          <label>Quantity</label>
          <input 
            v-model.number="form.quantity" 
            type="number" 
            min="0" 
            required
          >
        </div>
  
        <!-- Location -->
        <div class="form-group">
          <label>Location</label>
          <input v-model="form.location" required>
        </div>
  
        <!-- Status -->
        <div class="form-group">
          <label>Status</label>
          <select v-model="form.status">
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </select>
        </div>
  
        <div class="form-actions">
          <button type="button" class="btn-cancel" @click="router.back()">
            Cancel
          </button>
          <button type="submit" class="btn-submit">
            Save Changes
          </button>
        </div>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useProductsStore } from '../stores/products'
  
  const route = useRoute()
  const router = useRouter()
  const productsStore = useProductsStore()
  
  const form = ref({
    product_name: '',
    sku: '',
    quantity: 0,
    location: '',
    status: 'active'
  })
  
  // Load product data
  onMounted(async () => {
    await productsStore.fetchProducts()
    const product = productsStore.products.find(p => p.id === route.params.id)
    if (product) {
      form.value = { 
        product_name: product.product_name,
        sku: product.sku,
        quantity: product.quantity,
        location: product.location,
        status: product.status
      }
    }
  })
  
  const handleSubmit = async () => {
    try {
      await productsStore.updateProduct(route.params.id, form.value)
      router.push(`/products/${route.params.id}`)
    } catch (error) {
      console.error('Error updating product:', error)
      alert('Failed to update product')
    }
  }
  </script>
  
  <style scoped>
  .edit-product-container {
    max-width: 600px;
    margin: 2rem auto;
    padding: 2rem;
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    align-items: center;
  }
  
  .form-group {
    margin-bottom: 1.5rem;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #333;
  }
  
  .form-group input,
  .form-group select {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 1rem;
  }
  
  .form-group input[type="number"] {
    -moz-appearance: textfield;
  }
  
  .form-group input[type="number"]::-webkit-outer-spin-button,
  .form-group input[type="number"]::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }
  
  .form-actions {
    margin-top: 2rem;
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
  }
  
  .btn-cancel {
    background-color: #f8f9fa;
    color: #333;
    padding: 0.75rem 1.5rem;
    border: 1px solid #ddd;
    border-radius: 6px;
    cursor: pointer;
  }
  
  .btn-submit {
    background-color: #3b82f6;
    color: white;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 6px;
    cursor: pointer;
  }
  
  .btn-submit:hover {
    background-color: #2563eb;
  }
  
  @media (max-width: 768px) {
    .edit-product-container {
      margin: 1rem;
      padding: 1.5rem;
    }
    
    .form-actions {
      flex-direction: column;
    }
    
    .btn-cancel,
    .btn-submit {
      width: 100%;
    }
  }
  </style>