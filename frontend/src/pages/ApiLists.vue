<template>
    <div class="api-docs-container">
        <h1>APIs Documentation</h1>
        <div class="swagger-container">
            <iframe 
                src="http://localhost:8000/swagger/index.html"
                frameborder="0"
                width="100%"
                height="800px"
            ></iframe>
        </div>
    </div>
</template>

<script setup>

</script>

<style scoped>
.api-docs-container {
    padding: 20px;
    max-width: 1200px;
    margin: 0 auto;
}

h1 {
    margin-bottom: 20px;
    color: #333;
}

.swagger-container {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    overflow: hidden;
}

iframe {
    display: block;
}
</style>