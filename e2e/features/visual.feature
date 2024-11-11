@VISUAL
Feature: visual e2e tests

  Background:
    Given I open a new browser tab
    Then I navigate to visual e2e page


  Scenario: User should see certain things on page
    When visual test button should be visible
    And I should see "Voici le titre de la page" on the page


  Scenario: User should see certain things on page
    Given I should not see "L'élément a été caché." on the page
    And I should see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    When I click on the button which contains "Cacher l'élément"
    Then I should not see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    And I should see "L'élément a été caché." on the page


