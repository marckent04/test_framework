import type { SentenceDefinition } from "./sentence-definition.type";

export const formSentences: SentenceDefinition[] = [
  {
    sentence: 'I type "{text}" into the {field name}',
    description: "Type the specified text in form field.",
    variables: [
      { name: "text", type: "string" },
      { name: "field name", type: "string" },
    ],
    gherkinExample: 'When I type "John Doe" into the username field',
  },
  {
    sentence: 'I select "{option}" into the {name} dropdown',
    description:
      "Selects the specified option from the dropdown list identified by its label.",
    variables: [
      { name: "option", type: "string" },
      { name: "name", type: "string" },
    ],
    gherkinExample: 'When I select "Electronics" into the Category dropdown',
  },
  {
    sentence: "the {checkbox name} checkbox should be {checked|unchecked}",
    description:
      "Verifies the checked state of the checkbox identified by its label.",
    variables: [
      { name: "checkbox name", type: "string" },
      { name: "status", type: "string" }, // Note: This could be improved with a more specific type
    ],
    gherkinExample:
      'Then the "Subscribe to newsletter" checkbox should be checked',
  },
  {
    sentence: 'the {field} should be contain "{value}"',
    description: "Checks if the field contains the specified text.",
    variables: [
      { name: "field", type: "string" },
      { name: "value", type: "string" },
    ],
    gherkinExample: 'Then the result field should contain "10.5"',
  },
  {
    sentence: 'the {string} dropdown should have "{string}" selected',
    description:
      "Checks that the specified option is selected in the dropdown list identified by its label.",
    variables: [
      { name: "dropdown name", type: "string" },
      { name: "option", type: "string" },
    ],
    gherkinExample:
      'Then the Category dropdown should have "Electronics" selected',
  },
];
