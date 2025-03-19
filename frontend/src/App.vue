<template>
  <div class="app-container">
    <!-- Main content area where router views will be rendered -->
    <router-view></router-view>
  </div>
</template>

<script setup>
// No need for additional setup logic here since routing is handled by Vue Router
// and state management is handled by Pinia stores
</script>

<style>
/* Basic global styles */
body {
  margin: 0;
  padding: 20px;
  font-family: Arial, sans-serif;
}

.app-container {
  max-width: 1200px;
  margin: 0 auto;
}

/* Table styling */
table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f4f4f4;
}

tr:nth-child(even) {
  background-color: #f9f9f9;
}

/* Form styling */
form {
  margin: 20px 0;
  padding: 20px;
  border: 1px solid #eee;
  border-radius: 4px;
}

label {
  display: block;
  margin-bottom: 10px;
}

input, select {
  width: 100%;
  padding: 8px;
  margin-top: 4px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  background-color: #007bff;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 5px;
}

button:hover {
  background-color: #0056b3;
}

/* Navigation styling */
nav {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #f8f9fa;
  border-bottom: 1px solid #ddd;
}

nav button {
  float: right;
}
</style>