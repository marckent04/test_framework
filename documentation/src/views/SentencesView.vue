<template>
    <DefaultLayout>
        <h2 class="text-2xl font-bold">Filter by</h2>

        <div class="flex items-center space-x-4">
            <div>
                <h2 class="font-bold">Categories</h2>
                <div class="flex space-x-4">
                    <a href="#" class="menu-link" :class="{ 'menu-active-link': currentCategory === 'visual' }"
                        @click="filterByCategory('visual')">Visual</a>

                    <a href="#" class="menu-link" :class="{ 'menu-active-link': currentCategory === 'form' }"
                        @click="filterByCategory('form')">Form</a>

                    <a href="#" class="menu-link" @click="filterByCategory('navigation')"
                        :class="{ 'menu-active-link': currentCategory === 'navigation' }">Navigation</a>

                    <a href="#" class="menu-link" @click="filterByCategory('keyboard')"
                        :class="{ 'menu-active-link': currentCategory === 'keyboard' }">Keyboard</a>
                </div>
            </div>
            <h2>OR</h2>
            <div>
                <h2 class="font-bold">Keyword</h2>
                <input type="search" placeholder="Search for a sentence" ref="searchInput" class="search-input"
                    @input="(v) => filterByCategory('search', (v.target as any).value as string)">
            </div>
        </div>

        <section v-if="currentSection" class="mt-8">
            <h2 class="text-xl font-bold mb-2" v-text="currentSection.title"></h2>
            <p v-text="currentSection.description"></p>
            <p>Number of available phrases: {{ currentSection.sentences.length }}</p>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <SentenceDefinitionCard v-for="definition in currentSection.sentences" v-bind="definition"
                    :key="definition.sentence" />
            </div>
        </section>
    </DefaultLayout>
</template>


<script lang="ts">
type SentencesSectionDefinition = {
    title: string;
    description: string;
    sentences: SentenceDefinition[];
};

type SentencesCategory = 'visual' | 'form' | 'navigation' | 'keyboard' | 'search';


</script>
<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import DefaultLayout from '../layouts/DefaultLayout.vue';
import SentenceDefinitionCard from '../components/SentenceDefinitionCard.vue';
import { formSentences, keyboardSentences, navigationSentences, visualSentences, type SentenceDefinition } from '../data';

const currentCategory = ref<SentencesCategory>();
const currentSentences = ref<SentenceDefinition[]>([]);
const currentSection = ref<SentencesSectionDefinition | null>(null);
const searchInput = ref<HTMLInputElement>();

const allSentences = [...formSentences, ...keyboardSentences, ...navigationSentences, ...visualSentences];
function filterByCategory(category: SentencesCategory, searchValue?: string) {
    currentCategory.value = category;
    const defaultcategoriesSentencesSections: Record<string, SentencesSectionDefinition> = {
        visual: {
            title: 'Visual',
            description: 'Description of visual interactions.',
            sentences: visualSentences,
        },
        form: {
            title: 'Form',
            description: 'Description of interactions with forms.',
            sentences: formSentences,
        },
        navigation: {
            title: 'Navigation',
            description: 'Description of navigation interactions.',
            sentences: navigationSentences,
        },
        keyboard: {
            title: 'Keyboard',
            description: 'Description of keyboard interactions.',
            sentences: keyboardSentences,
        },
        search: {
            title: 'Search results',
            description: `Search results for "${searchValue}"`,
            sentences: allSentences,
        },
    };

    const section = defaultcategoriesSentencesSections[category];

    if (searchValue) {
        const sentences = section.sentences.filter(sentence => {
            const isSentenceMatches = sentence.sentence.toUpperCase().includes(searchValue.toUpperCase());
            const isDescriptionMatches = sentence.description.toUpperCase().includes(searchValue.toUpperCase());

            return isSentenceMatches || isDescriptionMatches;
        });

        currentSection.value = {
            title: 'Search results',
            description: `Search results for "${searchValue}"`,
            sentences,
        };
        return;
    }

    if (searchInput.value)
        searchInput.value.value = '';

    currentSection.value = section;
}


const searchValue = ref('');
watchEffect(() => {
    if (!searchValue.value) {
        currentSentences.value = [];
        return;
    }

    const sentences = allSentences.filter(sentence => {
        const isSentenceMatches = sentence.sentence.toUpperCase().includes(searchValue.value.toUpperCase());
        const isDescriptionMatches = sentence.description.toUpperCase().includes(searchValue.value.toUpperCase());

        return isSentenceMatches || isDescriptionMatches;
    });

    currentSection.value = {
        title: 'Search results',
        description: `Search results for "${searchValue.value}"`,
        sentences,
    };
});
</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}

#sentences-menu {
    @apply mb-8 flex justify-center;
}

.menu-link {
    @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}

.menu-active-link {
    @apply bg-blue-700;
}

.search-input {
    @apply block min-w-0 grow py-1.5 pl-1 pr-3 text-base text-gray-900 placeholder:text-gray-400 focus:outline focus:outline-0 sm:text-sm/6;
}
</style>