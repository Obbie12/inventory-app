<template>
  <div class="container">
    <h2>Products List</h2>
    
    <div class="filters">
      <div class="filter-group">
        <select v-model="selectedStatus" class="status-filter">
          <option value="">All Status</option>
          <option value="active">Active</option>
          <option value="inactive">Inactive</option>
        </select>
      </div>
      
      <div class="filter-group">
        <label class="lowstock-filter">
          <input type="checkbox" v-model="showLowStock">
          Show Low Stock (≤10)
        </label>
      </div>

      <div class="filter-group">
        <button class="export-btn" @click="exportToCSV">
          Export to CSV
        </button>
      </div>
    </div>

    <div class="table-container">
      <table class="product-table">
        <thead>
          <tr>
            <th>Product Name</th>
            <th>SKU</th>
            <th>Quantity</th>
            <th>Location</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        
        <tbody>
          <tr v-for="product in filteredProducts" :key="product.id">
            <td><router-link :to="`/products/${product.id}`">
            {{ product.product_name }}
            </router-link>
            </td>
            <!-- <td>{{ product.product_name }}</td> -->
            <td>{{ product.sku }}</td>
            <td :class="{ 'low-stock': product.quantity <= 10 }">
              {{ product.quantity }}
            </td>
            <td>{{ product.location }}</td>
            <td>
              <span class="status-tag" :class="product.status">
                {{ product.status }}
              </span>
            </td>
            <td>
              <button class="action-btn edit" @click="$emit('edit', product)">Edit</button>
              <button class="action-btn delete" @click="$emit('delete', product.id)">Delete</button>
              <button class="action-btn barcode" @click="generateBarcode(product.id)">Barcode</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()

const API_BASE = import.meta.env.VITE_API_BASE_URL;

const props = defineProps({
  products: Array
})

const emit = defineEmits(['edit', 'delete'])

const selectedStatus = ref('')
const showLowStock = ref(false)

const filteredProducts = computed(() => {
  return props.products.filter(product => {
    const statusMatch = !selectedStatus.value || product.status === selectedStatus.value
    const stockMatch = !showLowStock.value || product.quantity <= 10
    return statusMatch && stockMatch
  })
})
const exportToCSV = async () => {
  try {
    const response = await fetch(`${API_BASE}/api/v1/export/products`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (!response.ok) throw new Error('Export failed')

    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'products_export.csv'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    window.URL.revokeObjectURL(url)
    
  } catch (error) {
    console.error('Export error:', error)
    alert('Failed to export data')
  }
}
const generateBarcode = async (productId) => {
  try {
    const response = await fetch(`${API_BASE}/api/v1/products/${productId}/barcode`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (!response.ok) throw new Error('Failed to generate barcode')

    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    
    // Buka di tab baru
    const newWindow = window.open()
    newWindow.document.write(`
      <img src="${url}" style="max-width: 100%; height: auto;">
    `)
    
    // Bersihkan memori setelah 10 menit
    setTimeout(() => {
      window.URL.revokeObjectURL(url)
    }, 600000)

  } catch (error) {
    console.error('Barcode error:', error)
    alert('Failed to generate barcode')
  }
}
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  color: #333;
  margin-bottom: 20px;
  font-weight: 600;
}

.filters {
  display: flex;
  gap: 20px;
  margin-bottom: 15px;
}

.status-filter {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.lowstock-filter {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
}

.table-container {
  border: 1px solid #eee;
  border-radius: 6px;
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0,0,0,0.05);
}

.product-table {
  width: 100%;
  border-collapse: collapse;
  background: white;
}

.product-table th,
.product-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.product-table th {
  background-color: #f8f9fa;
  font-weight: 500;
  color: #666;
}

.product-table tbody tr:hover {
  background-color: #fafafa;
}

.status-tag {
  display: inline-block;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.status-tag.active {
  background: #e8f5e9;
  color: #2e7d32;
}

.icon-btn.barcode:hover {
  color: #2196F3;
  background: #e3f2fd;
}

.status-tag.inactive {
  background: #ffebee;
  color: #c62828;
}

.low-stock {
  color: #d32f2f;
  font-weight: 500;
}

.action-btn {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
  margin: 2px;
}

.edit {
  background-color: #e3f2fd;
  color: #1976d2;
}

.delete {
  background-color: #ffebee;
  color: #d32f2f;
}

@media (max-width: 768px) {
  .table-container {
    overflow-x: auto;
  }
  
  .product-table {
    min-width: 800px;
  }
  
  .filters {
    flex-direction: column;
  }
}
</style>