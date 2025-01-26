import type { SentenceDefinition } from "./sentence-definition.type";

export const navigationSentences: SentenceDefinition[] = [
  {
    sentence: "I navigate to {name} page",
    description: "Navigates the user to the specified page.",
    variables: [{ name: "page name", type: "string" }],
    gherkinExample: 'Given I navigate to "Contact" page',
  },
  {
    sentence: "I open a new browser tab",
    description: "Opens a new tab in the browser.",
    variables: [],
    gherkinExample: "When I open a new browser tab",
  },
  {
    sentence: "I open a new private browser tab",
    description: "Opens a new private tab in the browser.",
    variables: [],
    gherkinExample: "When I open a new private browser tab",
  },
  {
    sentence: "I should be navigated to {string} page",
    description: "Checks that the user is redirected to the specified page.",
    variables: [
      {
        name: "page",
        type: "string",
      },
    ],
    gherkinExample: "Then I should be navigated to Contact page",
  },
];
