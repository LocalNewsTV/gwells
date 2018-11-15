/*
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package pages

class AquifersPage extends BaseAppPage {
  static at = { pageTitle.text() == 'Aquifer Search' }
  static url = 'aquifers/#/'
  static content = {
    pageTitle { $('main h5') }

    aquiferNameField { $('#aquifers-name') }
    aquiferNumberField { $('#aquifers-number') }

    aquiferSearchButton { $('#aquifers-search') }

    aquiferResultsTable { $('#aquifers-results') }
  }

  void setAquiferName(String name) {
    aquiferNameField.value(name)
  }

  void setAquiferNumber(int number) {
    aquiferNumberField.value(number)
  }

  void clickSearchButton() {
    aquiferSearchButton.click()
  }

  Boolean foundSearchResults() {
    waitFor { aquiferResultsTable.$('tbody tr').size() != 0 }
  }
}
