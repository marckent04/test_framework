Feature: Products search (JUMIA)

  Background:
    Given I open a new private browser tab             # i_open_private_browser_tab.go:10 -> cucumber/frontend/navigation.init.func3.1
    And I navigate to jumia home page                  # i_navigate_to_page.go:12 -> cucumber/frontend/navigation.init.func2.1
    And I click on "newsletter modal close" button     # i_click_on_button_or_element.go:11 -> cucumber/frontend/visual.init.func2.1

  Scenario: A user search an iphone 15 on the web shop # features/search_product.feature:8
    When I fill the search input with "Iphone 15"      # i_fill_input_with.go:11 -> cucumber/frontend/form.init.func1.1
    And I press the enter button                       # i_press_enter_button.go:10 -> cucumber/frontend/keyboard.init.func1.1
    Then I am redirected to iphone 15 results page     # i_m_redirected_to_page.go:13 -> cucumber/frontend/navigation.init.func1.1
    And I must see on page 15 iphone 15                # i_must_see_on_page_x_elements.go:12 -> cucumber/frontend/visual.init.func4.1
    15 iphone 15 expected but 14 iphone 15 found

--- Failed steps:

  Scenario: A user search an iphone 15 on the web shop # features/search_product.feature:8
    And I must see on page 15 iphone 15 # features/search_product.feature:16
      Error: 15 iphone 15 expected but 14 iphone 15 found


1 scenarios (1 failed)
7 steps (6 passed, 1 failed)
6.32586561s
