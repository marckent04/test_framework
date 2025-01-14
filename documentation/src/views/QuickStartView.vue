<template>
    <DefaultLayout>
        <h1 class="text-3xl font-bold mb-6">Démarrage rapide avec etoolse</h1>

        <p>Ce guide s'adresse aux développeurs ou QA testers qui souhaitent apprendre à utiliser etoolsee pour lancer
            des
            tests e2e sur
            leurs applications web. Il suppose que vous etes a l'aise avec la ligne de commande, yaml et gherkin</p>

        <p>Dans cet exercice, nous allons tester notre propre documentation. Nous allons nous assurer que La recherche
            de phrases fonctionne correctement.</p>


        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">1. Installation et configuration</h2>
            <ol class="list-decimal list-inside mb-4">
                <li>Téléchargez etoolse pour votre système d'exploitation : <a href="[lien de téléchargement]"
                        target="_blank" class="underline">Lien de téléchargement</a>. Placez l'exécutable dans le
                    dossier de votre choix.</li>
                <li>Ouvrez votre terminal et tapez la commande suivante :
                    <code>etoolse init</code>
                    <p class="mt-2">Cette commande crée les fichiers de configuration `frontend.yml` et `cli.yml` à la
                        racine de votre projet, ainsi qu'un dossier `features` vide. </p>
                    <p>Le fichier `frontend.yml` contient les paramètres spécifiques à votre application web, tandis que
                        le fichier `cli.yml` contient les paramètres généraux d'etoolse.</p>
                </li>
            </ol>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md mt-6">
            <h2 class="text-xl font-semibold mb-4">2. Écriture du test</h2>

            <p>Commençons par écrire nos tests en langage naturel. Vous pouvez vous rendre sur la page de referentiel
                pour effectuer les actions et les noter. Une fois terminé vous pouvez derouler notre suggestion</p>
            <div>
                <a class="suggestion-toggle" @click="showSuggestion('natural_suggestion')">Afficher la suggestion</a>
                <div class="hidden" id="natural_suggestion">
                    <p class="italic">Pour vérifier que la recherche fonctionne correctement, nous allons
                        effectuer les actions suivantes :</p>
                    <ol class="list-decimal ml-10">
                        <li>Ouvrir le navigateur.</li>
                        <li>Accéder à la page de référencement des phrases.</li>
                        <li>Écrire "browser" dans le champ de recherche.</li>
                        <li>Vérifier que la phrase suivante apparaît à l'écran :
                            <code>I open a new browser tab</code>.
                        </li>
                    </ol>
                </div>


            </div>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md mt-6">
            <h2 class="text-xl font-semibold mb-4">Interactive Tutorial: Writing a Gherkin Scenario</h2>

            <p>Let's walk through creating a Gherkin scenario step-by-step. Our goal is to write a test that verifies
                the search functionality on the sentence referencing page. </p>

            <p>Here are the actions we want to test:</p>
            <ol class="list-decimal list-inside mb-4">
                <li>Open the browser.</li>
                <li>Go to the sentence referencing page.</li>
                <li>Type "browser" into the search field.</li>
                <li>Verify that the following sentence appears on the screen: <code>I open a new browser tab.</code>
                </li>
            </ol>

            <p>Now, let's translate these actions into Gherkin code.</p>

            <h3 class="text-lg font-medium mt-4">Step 1: Create a Feature File</h3>

            <p>First, create a new file named <code>search.feature</code> in your <code>features</code> directory. This
                file will contain our Gherkin scenarios.</p>

            <h3 class="text-lg font-medium mt-4">Step 2: Define the Feature</h3>

            <p>Start by defining the feature we are testing. In this case, it's the search functionality. Add the
                following line to your <code>search.feature</code> file:</p>

            <highlightjs language="gherkin" :code="`Feature: Search Functionality`" />

            <h3 class="text-lg font-medium mt-4">Step 3: Write the Scenario</h3>

            <p>Next, define a scenario within our feature. A scenario outlines a specific test case. We'll do this in
                stages:</p>

            <h4 class="text-md font-medium mt-3">Stage 1: Setting the Scene</h4>
            <p>Add the following line to your file:</p>
            <highlightjs language="gherkin" :code="`Scenario: Search for a sentence`" />
            <p>This gives our scenario a descriptive name.</p>

            <h4 class="text-md font-medium mt-3">Stage 2: Starting Point</h4>
            <p>Now, let's define the starting point of our test. Add this line:</p>
            <highlightjs language="gherkin"
                :code="`Given I open a new browser tab\nAnd i navigate to sentence referencing page`" />
            <p>This "Given" step sets the initial context for the scenario. It says we begin on the sentence referencing
                page.</p>

            <h4 class="text-md font-medium mt-3">Stage 3: User Action</h4>
            <p>Next, describe the action the user performs. Add this line:</p>
            <highlightjs language="gherkin" :code='`When I type "browser" into the search field`' />
            <p>This "When" step specifies the user action – typing "browser" into the search field.</p>

            <h4 class="text-md font-medium mt-3">Stage 4: Expected Outcome</h4>
            <p>Finally, define the expected result of the user's action. Add this line:</p>
            <highlightjs language="gherkin" :code='`Then I should see on page "I open a new browser tab."`' />
            <p>This "Then" step states the expected outcome – seeing a specific sentence on the page.</p>

            <p>Save the <code>search.feature</code> file.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md mt-6">
            <h2 class="text-xl font-semibold mb-4">4. Defining Variables</h2>
            <p>You can define variables in the <code>frontend.yml</code> file to make your scenarios work smoothly.
                These variables represent elements on your web page and make your Gherkin code more readable and
                maintainable.</p>

            <p>Here's how it works:</p>
            <ul class="list-disc list-inside ml-2">
                <li><b>Page URLs:</b> Store the URLs of frequently used pages, like the sentence referencing page. This
                    way, you can refer to them by name in your scenarios. For example:
                    <code>sentences_reference: "http://&lt;doc-base-url&gt;/sentences"</code>
                </li>
                <li><b>UI Elements:</b> Define variables for interactive elements like buttons, input fields, etc. This
                    makes your Gherkin steps clearer and easier to understand. For example, instead of writing a complex
                    CSS selector in your Gherkin step, you can use a variable like <code>search_field</code>.</li>
            </ul>

            <p><b>Important Tip:</b> You can even specify multiple CSS selectors for a single element! This is helpful
                if the element might be identified differently in various parts of your application. Etools will try
                each selector in order until it finds a match.</p>

            <p><b>Naming Convention:</b> To keep your code organized, use underscores to separate words in variable
                names (e.g., <code>login_button</code> instead of <code>loginButton</code>). This is a common practice
                in YAML files.</p>

            <p>Here's an example of how your <code>frontend.yml</code> file might look:</p>
            <highlightjs language="yml" :code='frontendYml' />

            <p><b>Remember:</b> Replace <code>&lt;doc-base-url&gt;</code> with the actual base URL of your
                documentation.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-md mt-6">
            <h2 class="text-xl font-semibold mb-4">5. Exécution des tests</h2>
            <p>Exécutez la commande suivante dans votre terminal pour lancer les tests e2e :</p>
            <highlightjs language="bash" :code="launchCommand" />


            <p>This command will execute your Gherkin scenario and provide you with a report of the test results.</p>

            <p>Congratulations! You've successfully written a Gherkin scenario to test the search functionality. You can
                now expand on this by adding more scenarios to cover other test cases.</p>
        </div>

    </DefaultLayout>
</template>

<script setup lang="ts">
import DefaultLayout from '../layouts/DefaultLayout.vue';

function showSuggestion(id: string) {
    const suggestion = document.getElementById(id);
    suggestion?.classList.remove('hidden');
}


const frontendYml = `
# frontend.yml
global:
  elements:
    search_field:
      - input[type='search'] 
      - .search-input 
    category_button: 
      - .menu-link 
        
   pages:
    sentences_reference: " http://<doc-base-url>/sentences" `


const finalGherkin = `
Feature: Search feature

    This scenario tests the search feature of the sentences reference page

    Scenario: Search feature works correctly
       Given I open a new browser tab
       And I navigate to sentences reference page
       When I type "browser" in the search field
       Then I should see "I open a new browser tab" on the page
       And I should see "I open a new private browser tab" on the page
`

const launchCommand = "./etoolse run --timeout 10s";

const gherkinReferenceUrl = "https://cucumber.io/docs/gherkin/reference/"

</script>

<style scoped>
.title {
    @apply text-3xl font-bold mb-6;
}

.suggestion-toggle {
    @apply cursor-pointer text-blue-500 hover:underline;
}
</style>