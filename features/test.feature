Feature: test

  Scenario: User login
    Given I open a new private browser tab
    And I navigate to home page
    When I fill the search input with "typeorm"
    And I click on "submit" button
    And I am redirected to results page
    And I click on "typeorm" element
    #And I click on element which contains "typeorm"
    Then I am redirected to typeorm details page