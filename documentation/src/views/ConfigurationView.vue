<template>
    <DefaultLayout>
        <section id="configuration" class="mb-12">
            <h2 class="text-2xl font-bold mb-4">Tool Configuration</h2>

            <p class="mb-4">
                Customize EToolsE's behavior to fit your needs with configuration options through a YAML file and
                command-line arguments.
            </p>

            <div>

                <h3 class="text-xl font-bold mb-2">Configuration via YAML File</h3>

                <p class="mb-4">
                    The YAML configuration file allows you to centralize EToolsE's settings.
                    Place this file, usually named <code>etoolse.yaml</code>, at the root of your project. Here's an
                    example
                    with explanations for each option:
                </p>

                <pre class="bg-gray-200 p-4 rounded mb-4"><code>
          configuration:
            # Maximum timeout for test execution (in seconds)
            timeout: 15s  
            # Number of tests to run in parallel
            parralel: 5  
            # Slow down test execution (in seconds) - useful for debugging
            slowMotion: 0s  
            # Path to the directory containing the Gherkin feature files
            gherkin_location: "e2e/features" 
          
          application:
            # Name of the application under test
            app_name: "MyApplication"  
            # Description of the application
            app_description: "my cool app"  
            # Version of the application
            app_version: "1.0.0"  
          
          reporting:
            # Format of the test report (HTML, JSON...)
            report_format: "html"  
            </code></pre>

                <p class="mb-4">
                    Feel free to modify these default values to adapt EToolsE to your environment and needs.
                </p>
            </div>


            <div>
                <h3 class="text-xl font-bold mb-2">Configuration via CLI Arguments</h3>

                <p class="mb-4">
                    You can also configure EToolsE by using command-line arguments. These arguments provide flexibility
                    and allow you to override settings from the YAML configuration file.
                </p>

                <ul class="list-disc list-inside mb-4">
                    <li><code>-l, --location &lt;path&gt;</code>: Specifies the path to the directory containing your
                        Gherkin feature files. (e.g., <code>--location ./features</code>)</li>
                    <li><code>-c, --config &lt;path&gt;</code>: Sets the path to the main EToolsE configuration YAML
                        file. Defaults to "cli.yml". (e.g., <code>--config config.yaml</code>)</li>
                    <li><code>-f, --front-config &lt;path&gt;</code>: Sets the path to a YAML file specifically for
                        frontend testing configuration. Defaults to "frontend.yml". (e.g.,
                        <code>--front-config frontend-tests.yaml</code>)
                    </li>
                    <li><code>-t, --tags &lt;tags&gt;</code>: Filters tests to run based on specified tags. (e.g.,
                        <code>--tags "@smoke,@regression"</code>)
                    </li>
                    <li><code>-p, --parallel &lt;number&gt;</code>: Defines the number of tests to execute concurrently.
                        (e.g., <code>--parallel 4</code>)</li>
                    <li><code>--timeout &lt;duration&gt;</code>: Sets the maximum duration for the entire test suite to
                        run before timing out. (e.g., <code>--timeout 30s</code>)</li>
                    <li><code>--headless &lt;bool&gt;</code>: Controls whether to run the browser in headless mode
                        (without a visible UI). Defaults to "true" (headless). (e.g., <code>--headless false</code> to
                        show the browser)</li>
                    <li><code>-v, --version &lt;string&gt;</code>: Specifies the version of the application under test.
                        Defaults to "1.0". (e.g., <code>--version 2.1.0</code>)</li>
                </ul>

                <p class="mb-4">
                    You can combine multiple CLI arguments for a more tailored configuration. For example, to run tests
                    from a specific directory with a 60-second timeout in parallel, you would use:
                    <br><code>etoolse -l ./my_features -p 3 --timeout 60s</code>
                </p>
            </div>
        </section>
        <div class="mb-12">
            <h2 class="text-2xl font-bold mb-4">Variables declarations</h2>
            <p>We have 2 variables types: <code>page</code> and <code>element</code></p>

            <p>All variables types are declared in <code>frontend.yml</code></p>
            <div>
                <h3 class="text-xl font-bold">pages variables</h3>
                <p>Pages variables are declared in the <code>pages</code> section</p>
                <p>Example:</p>
                <p>
                    With this sentence: <code>I navigate to FAQ page</code>
                </p>
                <p>The <code>FAQ</code> variable is defined like that :</p>
                <pre class="bg-gray-200 p-4 rounded mb-4">
                    <code>
                    pages:
                        FAQ: "https://www.example.com/faq"
                    </code>
                </pre>

            </div>
            <div>
                <h3 class="text-xl font-bold">elements variables</h3>
                <p>Elements variables are declared in the <code>elements</code> section</p>
                <p>Example:</p>
                <p>
                    With this sentence: <code>I click on the "Login" button</code>
                </p>
                <p>The <code>Login</code> variable is defined like that :</p>
                <pre class="bg-gray-200 p-4 rounded mb-4">
                    <code>
                    elements:
                        login: 
                    </code>
                </pre>
            </div>
        </div>
    </DefaultLayout>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import DefaultLayout from '../layouts/DefaultLayout.vue';
onMounted(() => {
    const menuLinks = document.querySelectorAll('#sentences-menu a');
    const sections = document.querySelectorAll('main section');

    menuLinks.forEach(link => {
        link.addEventListener('click', (event) => {
            event.preventDefault();
            const targetId = link.dataset.target;
            sections.forEach(section => {
                section.classList.add('hidden');
            });
            document.getElementById(targetId)?.classList.remove('hidden');
        });
    });
})
</script>

<style scoped>
.sentences-grid {
    @apply grid grid-cols-1 md:grid-cols-2 gap-4;
}
</style>