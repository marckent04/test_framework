import type { SentenceDefinition } from "./sentence-definition.type";

export const keyboardSentences: SentenceDefinition[] = [
  {
    sentence: 'I press the "{key}" button',
    description: "Simulates pressing the specified keyboard key.",
    variables: [
      {
        name: "key",
        type: "arrow up | tab | delete | escape | space | arrow right | arrow down | arrow left | enter",
      },
    ],
    gherkinExample: 'When I press the "tab" key',
  },
];
