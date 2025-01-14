<template>
    <div class="bg-white p-4 rounded-md shadow-md">
        <h3 class="text-lg font-bold mb-2" v-text="sentence"></h3>
        <p><span class="font-bold">Description: </span>{{ description }}</p>
        <p><span class="font-bold">Variables: </span>
            <span v-if="variables.length === 0">None</span>
            <template v-else>
                <span v-for="(variable, i) in variables">{{ variable.name }} ({{ variable.type }}) {{ i ===
                    variables.length - 1 ? "" : " / " }}</span>
            </template>
        </p>

        <highlightjs language='gherkin' :code="gherkinExample.trim()" />
    </div>
    <!-- 

    <div class="bg-white p-4 rounded-md shadow-md">
                    <h3 class="text-lg font-bold mb-2">I should be navigated to {page} page</h3>
                    <p>Variables: page (string)</p>
                    <p>Description: Checks that the user is redirected to the specified page.</p>
                    <pre class="bg-gray-200 p-4 rounded-md"><code>
          Scenario: Verify redirection after login
            Given I am on the login page
            When I enter my credentials and click "Login"
            <mark>Then I should be navigated to "Home" page</mark>
                </code></pre>
                </div> -->
</template>

<script lang="ts" setup>
import type { SentenceDefinition } from '../data';

defineProps<SentenceDefinition>();
</script>
