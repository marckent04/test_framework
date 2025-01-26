import type { SentenceDefinition } from "./sentence-definition.type";

export const visualSentences: SentenceDefinition[] = [
  {
    sentence: "I click on the row containing the following elements",
    description:
      "Clicks on the row of a table containing the specified elements.",
    variables: [
      {
        name: "elements",
        type: "gherkin map array",
      },
    ],
    gherkinExample: `
    When I click on the row containing the following elements
      | name      | description                        | price   |
      | Produit 1 | Description détaillée du produit 1 | 19.99 € |
    `,
  },
  {
    sentence: "I should see a row containing the following elements",
    description:
      "Checks that a row in a table contains the specified elements.",
    variables: [
      {
        name: "elements",
        type: "gherkin map array",
      },
    ],
    gherkinExample: `
       Then I should see a row containing the following elements
        | name      | description                        | price   |
        | Produit 1 | Description détaillée du produit 1 | 19.99 € |
      `,
  },
  {
    sentence: "I should not see a row containing the following elements",
    description:
      "Checks that no row in a table contains the specified elements.",
    variables: [
      {
        name: "elements",
        type: "gherkin map array",
      },
    ],
    gherkinExample: `
      Then I should not see a row containing the following elements
      | name      | description                        | price   |
      | Produit 1 | Description détaillée du produit 0 | 19.99 € |
      `,
  },
  {
    sentence: "I should see a table with the following headers",
    description: "Checks that the table has the specified headers.",
    variables: [
      {
        name: "elements",
        type: "gherkin map array",
      },
    ],
    gherkinExample: `Then I should see a table with the following headers
        | product name        | Produit     |
        | product description | Description |
        | product price       | Prix        |
      `,
  },
  {
    sentence: "{element} should be visible",
    description: "Checks that the specified element is visible on the page.",
    variables: [{ name: "element", type: "string" }],
    gherkinExample: "Then confirmation alert should be visible",
  },
  {
    sentence: "{element} should not be visible",
    description:
      "Checks that the specified element is not visible on the page.",
    variables: [{ name: "element", type: "string" }],
    gherkinExample: "Then error alert should not be visible",
  },
  {
    sentence: "I click on {element}",
    description: "Clicks on the element",
    variables: [{ name: "element", type: "string" }],
    gherkinExample: "When I click on login button",
  },
  {
    sentence: 'I click on {element} which contains "{text}"',
    description:
      "Clicks on the specified element that contains the specified text.",
    variables: [
      { name: "element", type: "string" },
      { name: "text", type: "string" },
    ],
    gherkinExample: 'When I click on card which contains "See more"',
  },
  {
    sentence: 'I should see "{text}" on the page',
    description: "Checks that the specified text is present on the page.",
    variables: [{ name: "text", type: "string" }],
    gherkinExample: 'Then I should see "Welcome to our website" on the page',
  },
  {
    sentence: 'I should not see "{string}" on the page',
    description: "Checks that the specified text is not present on the page.",
    variables: [{ name: "text", type: "string" }],
    gherkinExample: 'Then I should not see "Connection error" on the page',
  },
  {
    sentence: "I should see a (link|button|element) which contains {text}",
    description:
      "Checks that there is a link, button or element containing the specified text.",
    variables: [{ name: "text", type: "string" }],
    gherkinExample: 'Then I should see a button which contains "Login"',
  },
  {
    sentence: "I should see {number} {item}",
    description:
      "Checks that the specified number of elements containing the specified text are present on the page.",
    variables: [
      { name: "number", type: "number" },
      { name: "item", type: "string" },
    ],
    gherkinExample: "Then I should see 10 products",
  },
  {
    sentence: 'I should see "{string}" details on the page',
    description:
      "Checks that the details for the specified text are present on the page.",
    variables: [
      { name: "text", type: "string" },
      { name: "details", type: "gherkin map" },
    ],
    gherkinExample: `
        Then I should see "computer" details on the page
      | name        | Ordinateur de Bord pour Rameur                                                |
      | description | Cet ordinateur de rameur vous permet de suivre vos performances en temps réel |
      | price       | 149,99 €                                                                      |
      | screen type | LCD rétroéclairé   
    `,
  },
];
