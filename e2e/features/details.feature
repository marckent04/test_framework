@ELEMENT_DETAILS
Feature: product details e2e tests

  Scenario: a user must see computer details
    Given I open a new private browser tab
    When I navigate to details e2e page
    Then I should see "computer" details on the page
      | name        | Ordinateur de Bord pour Rameur                                                |
      | description | Cet ordinateur de rameur vous permet de suivre vos performances en temps réel |
      | price       | 149,99 €                                                                      |
      | screen type | LCD rétroéclairé                                                              |
