<template>
    <header class="bg-blue-500 py-4">
        <div class="container mx-auto flex items-center justify-between text-white">
            <router-link tag="h1" to="/" class="text-3xl font-bold">EToolsE</router-link>
            <nav>
                <ul class="flex space-x-4">
                    <li><router-link :to="{ name: routeNames.home }" class="hover:underline">Home</router-link></li>
                    <li><router-link :to="{ name: routeNames.getStarted }" class="hover:underline">Get
                            started</router-link></li>
                    <li><router-link :to="{ name: routeNames.quickstart }" class="hover:underline">Quick
                            Start</router-link></li>
                    <li><router-link :to="{ name: routeNames.config }"
                            class="hover:underline">Configuration</router-link>
                    </li>
                    <li><router-link :to="{ name: routeNames.sentences }"
                            class="hover:underline">Sentences</router-link>
                    </li>
                    <li><a target="_blank" :href="githubUrl" class="hover:underline">Github</a>
                    </li>
                    <!-- <li><a href="examples.html" class="hover:underline">Examples</a></li> -->
                    <!-- <li><a href="faq.html" class="hover:underline">FAQ</a></li> -->
                </ul>
            </nav>
        </div>
    </header>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
import { routeNames } from "../views";

const githubUrl = "https://github.com/marckent04/test_framework";
</script>