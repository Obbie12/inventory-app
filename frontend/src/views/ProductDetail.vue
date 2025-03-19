<!-- src/views/ProductDetail.vue -->
<template>
    <div class="product-detail-container">
      <div class="header">
        <button class="back-btn" @click="$router.go(-1)">← Back</button>
        <h1>Product Details</h1>
      </div>
  
      <div v-if="loading" class="loading">Loading...</div>
      
      <div v-else-if="product" class="detail-content">
        <div class="detail-section">
          <div class="detail-row">
            <span class="label">Product Name:</span>
            <span class="value">{{ product.product_name }}</span>
          </div>
          
          <div class="detail-row">
            <span class="label">SKU:</span>
            <span class="value">{{ product.sku }}</span>
          </div>
  
          <div class="detail-row">
            <span class="label">Quantity:</span>
            <span class="value" :class="{ 'low-stock': product.quantity <= 10 }">
              {{ product.quantity }}
            </span>
          </div>
  
          <div class="detail-row">
            <span class="label">Location:</span>
            <span class="value">{{ product.location }}</span>
          </div>
  
          <div class="detail-row">
            <span class="label">Status:</span>
            <span class="status-badge" :class="product.status">
              {{ product.status }}
            </span>
          </div>
  
          <div class="detail-row">
            <span class="label">Created At:</span>
            <span class="value">{{ formatDate(product.created_at) }}</span>
          </div>
  
          <div class="detail-row">
            <span class="label">Last Updated:</span>
            <span class="value">{{ formatDate(product.updated_at) }}</span>
          </div>
        </div>
  
        <div class="action-buttons">
            <button class="generate-barcode-btn" @click="generateBarcode">Generate Barcode</button>
          <button class="edit-btn" @click="handleEdit">Edit Product</button>
          <button class="delete-btn" @click="handleDelete">Delete Product</button>
        </div>
      </div>
      
      <div v-else class="not-found">
        Product not found
      </div>
      <div v-if="barcodeUrl" class="barcode-modal">
      <div class="modal-content">
        <h3>Product Barcode</h3>
        <img :src="barcodeUrl" alt="Product Barcode">
        <button class="close-btn" @click="barcodeUrl = null">Close</button>
      </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useProductsStore } from '../stores/products'
  import { useAuthStore } from '../stores/auth'
  
  const API_BASE = import.meta.env.VITE_API_BASE_URL;
  const authStore = useAuthStore()
  const route = useRoute()
  const router = useRouter()
  const productsStore = useProductsStore()
  const barcodeUrl = ref(null)
  
  const product = ref(null)
  const loading = ref(true)
  
  onMounted(async () => {
    try {
      await productsStore.fetchProducts()
      product.value = productsStore.products.find(
        p => p.id === route.params.id
      )
    } catch (error) {
      console.error('Error fetching product:', error)
    } finally {
      loading.value = false
    }
  })
  
  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }
  
  const handleEdit = () => {
    router.push({
        name: 'EditProduct',
        params: { id: product.value.id }
    })
  }
  
  const handleDelete = async () => {
    if (confirm('Are you sure you want to delete this product?')) {
      await productsStore.deleteProduct(route.params.id)
      router.push('/dashboard')
    }
  }
  const generateBarcode = async () => {
  try {
    const response = await fetch(`${API_BASE}/api/v1/products/${route.params.id}/barcode`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (!response.ok) throw new Error('Failed to generate barcode')

    const blob = await response.blob()
    barcodeUrl.value = URL.createObjectURL(blob)
    
  } catch (error) {
    console.error('Barcode error:', error)
    alert('Failed to generate barcode')
  }
}
  </script>
  
  <style scoped>
  .product-detail-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 2rem;
  }
  
  .header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 2rem;
  }
  
  .back-btn {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1.2rem;
    color: #666;
    padding: 0.5rem;
  }
  
  .detail-content {
    background: white;
    border-radius: 8px;
    padding: 2rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .detail-section {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
  }
  
  .detail-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    border-bottom: 1px solid #eee;
  }
  
  .label {
    color: #666;
    font-weight: 500;
  }
  
  .value {
    color: #333;
  }
  
  .low-stock {
    color: #dc2626;
    font-weight: 500;
  }
  
  .status-badge {
    padding: 0.25rem 0.75rem;
    border-radius: 12px;
    font-size: 0.85rem;
  }
  
  .status-badge.active {
    background: #dcfce7;
    color: #166534;
  }
  
  .status-badge.inactive {
    background: #fee2e2;
    color: #991b1b;
  }

  .generate-barcode-btn {
  background-color: #8b5cf6;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.generate-barcode-btn:hover {
  background-color: #7c3aed;
}

.barcode-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  text-align: center;
  max-width: 90%;
}

.modal-content img {
  max-width: 300px;
  margin: 1rem 0;
}

.close-btn {
    background-color: #ef4444;
    color: white;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
  
  .action-buttons {
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
    margin-top: 2rem;
  }
  
  .edit-btn {
    background-color: #3b82f6;
    color: white;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 6px;
    cursor: pointer;
  }
  
  .delete-btn {
    background-color: #ef4444;
    color: white;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 6px;
    cursor: pointer;
  }
  
  .loading,
  .not-found {
    text-align: center;
    padding: 2rem;
    font-size: 1.2rem;
    color: #666;
  }
  
  @media (max-width: 768px) {
    .detail-section {
      grid-template-columns: 1fr;
    }
    
    .action-buttons {
      flex-direction: column;
    }
    
    .edit-btn,
    .delete-btn {
      width: 100%;
    }
    .action-buttons {
    flex-direction: column;
  }
  
  .generate-barcode-btn,
  .edit-btn,
  .delete-btn {
    width: 100%;
  }
  
  .modal-content {
    padding: 1rem;
  }
  
  .modal-content img {
    max-width: 200px;
  }
  }
  </style>