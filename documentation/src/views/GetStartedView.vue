<template>
    <DefaultLayout>
        <section id="prerequisites" class="mb-12">
            <h2 class="text-2xl font-bold mb-4">Prerequisites</h2>

            <p class="mb-4">
                Before you get started with EToolsE, make sure you have the following:
            </p>

            <ul class="list-disc list-inside mb-4">
                <li>
                    <b>Chromium installed:</b> EToolsE uses Chromium for web testing.
                    <ul class=" list-inside ml-4">
                        <li>EtoolsE install automatically chromium but you can install yourself chormium</li>
                        <li>
                            <a href="https://www.chromium.org/getting-involved/download-chromium/" target="_blank"
                                rel="noopener noreferrer" class="text-blue-500 hover:underline">
                                Click here for instructions on how to install Chromium.
                            </a>
                        </li>
                    </ul>
                </li>
                <li>
                    <b>Understanding of Gherkin:</b> Gherkin is the language used to write your test scenarios.
                    <a href="https://cucumber.io/docs/gherkin/reference/" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        Refer to this guide to learn more about Gherkin.
                    </a>
                </li>
                <li>
                    <b>Understanding of YAML:</b> EToolsE uses YAML for configuration.
                    <a href="https://learnxinyminutes.com/docs/yaml/" target="_blank" rel="noopener noreferrer"
                        class="text-blue-500 hover:underline">
                        This tutorial provides a quick introduction to YAML.
                    </a>
                </li>
                <li>
                    <b>Command Line Interface (CLI):</b> You'll need to be comfortable using a CLI to run EToolsE
                    commands.
                </li>
            </ul>
        </section>

        <section id="download" class="mb-12">
            <h2 class="text-2xl font-bold mb-4">Download EToolsE</h2>

            <p class="mb-4">
                EToolsE is available for Windows, macOS, and Linux. Choose the download link that matches your operating
                system:
            </p>

            <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-8">
                <a href="#"
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/windows8/windows8-original.svg"
                        alt="Windows logo" class="h-6 mr-2"> Download for Windows
                </a>
                <a href="#"
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/apple/apple-original.svg"
                        alt="macOS logo" class="h-6 mr-2"> Download for macOS
                </a>
                <a href="#"
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex items-center">
                    <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linux/linux-original.svg"
                        alt="Linux logo" class="h-6 mr-2"> Download for Linux
                </a>
            </div>

            <p class="mt-4 text-sm text-gray-600">
                For configuration instructions, please refer to the
                <router-link :to="{ name: routeNames.config }"
                    class="text-blue-500 hover:underline">Configuration</router-link> guide.
            </p>

            <p class="mt-4 text-sm text-gray-600">
                For quick start tutorial, please refer to the
                <router-link :to="{ name: routeNames.quickstart }" class="text-blue-500 hover:underline">Quick
                    start</router-link>.
            </p>
            <p class="mt-4 text-sm text-gray-600">
                For available gherkin phrases, please refer to the
                <router-link :to="{ name: routeNames.sentences }"
                    class="text-blue-500 hover:underline">Sentences</router-link> dictionary.
            </p>
        </section>
    </DefaultLayout>
</template>

<script setup lang="ts">
import DefaultLayout from '../layouts/DefaultLayout.vue';
import { routeNames } from '../views';

</script>
