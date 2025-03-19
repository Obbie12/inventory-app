<template>
    <div v-if="show" class="modal-overlay">
      <div class="modal-content">
        <div class="modal-header">
          <h3>Add New Product</h3>
          <button class="close-btn" @click="$emit('close')">&times;</button>
        </div>
  
        <form @submit.prevent="handleSubmit">
          <div class="form-group">
            <label>Product Name</label>
            <input v-model="form.product_name" required>
          </div>
  
          <div class="form-group">
            <label>SKU</label>
            <input v-model="form.sku" required>
          </div>
  
          <div class="form-group">
            <label>Quantity</label>
            <input v-model="form.quantity" type="number" required>
          </div>
  
          <div class="form-group">
            <label>Location</label>
            <input v-model="form.location" required>
          </div>
  
          <div class="form-group">
            <label>Status</label>
            <select v-model="form.status">
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
            </select>
          </div>
  
          <div class="modal-actions">
            <button type="button" class="btn-cancel" @click="$emit('close')">
              Cancel
            </button>
            <button type="submit" class="btn-submit">
              Add Product
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const emit = defineEmits(['submit', 'close'])
  const props = defineProps({
    show: Boolean
  })
  
  const form = ref({
    product_name: '',
    sku: '',
    quantity: 0,
    location: '',
    status: 'active'
  })
  
  const handleSubmit = () => {
    emit('submit', form.value)
    resetForm()
  }
  
  const resetForm = () => {
    form.value = {
      product_name: '',
      sku: '',
      quantity: 0,
      location: '',
      status: 'active'
    }
  }
  
  watch(() => props.show, (newVal) => {
    if (!newVal) resetForm()
  })
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }
  
  .modal-content {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    width: 90%;
    max-width: 500px;
    animation: modalSlide 0.3s ease-out;
  }
  
  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
    padding-bottom: 1rem;
    border-bottom: 1px solid #eee;
  }
  
  .close-btn {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
    color: #666;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 0.5rem;
    color: #333;
  }
  
  .form-group input,
  .form-group select {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
  }
  
  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 2rem;
  }
  
  .btn-cancel {
    background-color: #f0f0f0;
    color: #333;
    padding: 0.5rem 1rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .btn-submit {
    background-color: #28a745;
    color: white;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  @keyframes modalSlide {
    from {
      transform: translateY(-50px);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }
  
  @media (max-width: 480px) {
    .modal-content {
      width: 95%;
      padding: 1rem;
    }
    
    .modal-actions {
      flex-direction: column;
    }
  }
  </style>