<template>
    <form @submit.prevent="handleSubmit">
      <label>
        Product Name:
        <input v-model="form.product_name" required>
      </label>
      <label>
        SKU:
        <input v-model="form.sku" required>
      </label>
      <label>
        Quantity:
        <input v-model="form.quantity" type="number" required>
      </label>
      <label>
        Location:
        <input v-model="form.location" required>
      </label>
      <label>
        Status:
        <select v-model="form.status">
          <option value="active">Active</option>
          <option value="inactive">Inactive</option>
        </select>
      </label>
      <button type="submit">{{ isEditing ? 'Update' : 'Add' }} Product</button>
    </form>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const props = defineProps({
    product: Object
  })
  
  const emit = defineEmits(['submit'])
  
  const isEditing = ref(false)
  const form = ref({
    product_name: '',
    sku: '',
    quantity: 0,
    location: '',
    status: 'active'
  })
  
  watch(() => props.product, (newProduct) => {
    if (newProduct) {
      isEditing.value = true
      form.value = { ...newProduct }
    }
  })
  
  const handleSubmit = () => {
    emit('submit', form.value)
    if (!isEditing.value) {
      form.value = {
        product_name: '',
        sku: '',
        quantity: 0,
        location: '',
        status: 'active'
      }
    }
  }
  </script>