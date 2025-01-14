<template>
    <div class="flex flex-col min-h-screen">
        <NavBar />

        <main class="container mx-auto py-8 flex-grow">
            <slot />
        </main>

        <Footer />
    </div>

</template>

<script setup lang="ts">
import Footer from '../components/Footer.vue';
import NavBar from '../components/Navbar.vue';
</script>
