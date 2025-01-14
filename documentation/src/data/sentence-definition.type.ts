export type SentenceDefinition = {
  sentence: string;
  description: string;
  variables: SentenceVariable[];
  gherkinExample: string;
};

export type SentenceVariable = {
  name: string;
  type: "string" | "number" | "boolean" | "gherkin map array" | "gherkin map";
};
