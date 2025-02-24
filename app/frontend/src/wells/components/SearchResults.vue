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
<template>
  <div id="searchResults">
    <b-row align-h="between" class="mb-2">
      <b-col sm="4" class="form-inline">
        Show <b-form-select class="mx-1" :value="limit" @input="setLimit($event)" :options="limitOptions" /> results
      </b-col>
      <b-col sm="4">
        <search-column-select
          id="searchColumnSelect"
          class="float-right"
          :column-data="searchFields" />
      </b-col>
    </b-row>
    <div class="table-responsive">
      <table id="searchResultsTable" class="table table-striped">
        <thead>
          <tr>
            <th
              v-for="column in columns"
              :key="column.id"
              class="text-nowrap vertical-align-middle"
              scope="col">
              {{ column.resultLabel ? column.resultLabel : column.label }}
              <b-button v-if="column.sortable"
                class="sort-button px-0"
                :class="{active: column.sortParam === orderingParam}"
                variant="link"
                @click="sortResults({ param: column.sortParam || column.sortParam, desc: (column.sortParam === orderingParam) ? !orderingDesc : false })">
                {{ (column.sortParam === orderingParam) ? orderingDesc ? '&#x2191;' : '&#x2193;' : '&#x2195;' }}
              </b-button>
            </th>
          </tr>
          <tr class="filters">
            <th v-for="column in columns" :key="column.id" :class="`search-result-filter-${column.type}`">
              <search-result-filter
                :type="column.type"
                :id="`${column.id}ResultFilter`"
                :errors="resultErrors[column.param]"
                :param-names="column.params"
                :options="column.options || filterSelectOptions[column.id]"
                :value-field="column.valueField"
                :text-field="column.textField"
                :value="filterParams[column.id]"
                @input="applyFilter(column, $event)" />
            </th>
          </tr>
        </thead>
        <tbody role="rowgroup" v-if="(isBusy || isEmpty || isReset)">
          <tr :class="{loading: isBusy}">
            <td :colspan="columnCount">
              <div class="position-relative my-1 mx-1">
                <div v-if="isBusy" class="d-flex">
                  <div class="spinner-border m-3" role="status" aria-hidden="true"></div>
                  <strong class="align-middle my-3 pt-1">
                    Loading...
                  </strong>
                </div>
                <strong v-else-if="hasResultErrors" class="text-danger align-middle my-3 pt-1">
                  An error occurred processing the search.
                </strong>
                <strong v-else-if="isReset" class="align-middle my-3 pt-1">
                  Enter some criteria in the form above to search.
                </strong>
                <strong v-else-if="isEmpty" class="align-middle my-3 pt-1">
                  No results matched your search criteria.
                </strong>
              </div>
            </td>
          </tr>
        </tbody>
        <tbody role="rowgroup" v-else>
          <tr v-for="row in results" :key="row.well_tag_number" @mousedown="searchResultsRowClicked(row)">
            <td v-for="column in columns" :key="column.id" class="data">
              <template v-if="column.param === 'well_tag_number'">
                <router-link :to="{ name: 'wells-detail', params: {id: row.well_tag_number} }">{{ row.well_tag_number }}</router-link>
              </template>
              <template v-else-if="column.param === 'street_address'">
                {{ row | streetAddressFormat }}
              </template>
              <template v-else-if="column.type === 'select' && Array.isArray(row[column.param])">
                <template v-for="(value, index) in row[column.param]">
                  <span :key="`${row.well_tag_number}-${column.param}-${index}`">
                    {{ value | selectOptionFormat(column, filterSelectOptions[column.id]) }}<span v-if="index < row[column.param].length - 1">, </span>
                  </span>
                </template>
              </template>
              <template v-else-if="column.type === 'select' || column.type === 'radio'">
                {{ row[column.param] | selectOptionFormat(column, filterSelectOptions[column.id]) }}
              </template>
              <template v-else-if="column.param === 'legal_pid'">
                {{ row[column.param] }}
              </template>
              <template v-else-if="column.param === 'licence_number'">
                <span v-for="(licence, index) in row.licence_number">
                  <a :href="LICENCE_URL + licence">
                    {{ licence }}{{ index + 1 < row.licence_number.length ? ', ' : '' }}
                  </a>
                </span>
              </template>
              <template v-else>
                {{ row[column.param] | defaultFormat }}
              </template>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="my-3" v-if="resultCount > 0">
      <div>Showing {{ currentRecordsCountStart }} to {{ currentRecordsCountEnd }} of {{ resultCount }} {{ resultCount === 1 ? 'record' : 'records'}}.</div>
      <search-result-exports class="my-3" :field-data="searchFields" />
    </div>
    <b-pagination
      class="mt-3"
      size="md"
      :disabled="isBusy"
      :total-rows="resultCount"
      :value="currentPage"
      :per-page="limit"
      @input="changePage($event)"/>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { SEARCH_WELLS } from '@/wells/store/actions.types.js'
import {
  SET_SEARCH_LIMIT,
  SET_SEARCH_OFFSET,
  SET_SEARCH_ORDERING,
  SET_SEARCH_RESULT_FILTERS
} from '@/wells/store/mutations.types.js'
import { FILTER_TRIGGER } from '@/wells/store/triggers.types.js'
import SearchResultExports from '@/wells/components/SearchResultExports.vue'
import SearchResultFilter from '@/wells/components/SearchResultFilter.vue'
import SearchColumnSelect from '@/wells/components/SearchColumnSelect.vue'
import filterMixin from '@/wells/components/mixins/filters.js'

export default {
  mixins: [filterMixin],
  components: {
    'search-column-select': SearchColumnSelect,
    'search-result-filter': SearchResultFilter,
    'search-result-exports': SearchResultExports
  },
  data () {
    return {
      limitOptions: [
        { value: 10, text: '10' },
        { value: 25, text: '25' },
        { value: 50, text: '50' }
      ],
      LICENCE_URL: 'https://j200.gov.bc.ca/pub/ams/Default.aspx?PossePresentation=AMSPublic&PosseObjectDef=o_ATIS_DocumentSearch&PosseMenuName=WS_Main&Criteria_LicenceNumber='
    }
  },
  computed: {
    ...mapGetters({
      limit: 'searchLimit',
      errors: 'searchErrors',
      params: 'searchParams',
      offset: 'searchOffset',
      ordering: 'searchOrdering',
      pending: 'pendingSearch',
      resultFilters: 'searchResultFilters',
      resultColumns: 'searchResultColumns',
      resultCount: 'searchResultCount',
      results: 'searchResults',
      selectedWells: 'selectedWells'
    }),
    columns () {
      return this.getFilterFields(this.resultColumns)
    },
    currentPage () {
      return Math.ceil(this.offset / this.limit) + 1
    },
    // currentRecordsCountStart is the starting record number in the table of wells
    // (e.g. the 1 in 'showing 1 to 10 of 25 records')
    currentRecordsCountStart () {
      return (this.currentPage - 1) * this.limit + 1
    },
    // currentRecordsCountEnd is the last visible record number in the table of wells
    // (e.g. the 10 in 'showing 1 to 10 of 25 records')
    currentRecordsCountEnd () {
      if (this.results === null) {
        return 0
      }
      return (this.currentPage - 1) * this.limit + this.results.length
    },
    orderingParam () {
      if (this.orderingDesc) {
        return this.ordering.substr(1)
      } else {
        return this.ordering
      }
    },
    orderingDesc () {
      return this.ordering.startsWith('-')
    },
    columnCount () {
      return this.resultColumns.length
    },
    resultErrors () {
      return this.errors.filter_group || {}
    },
    hasResultErrors () {
      return Object.entries(this.resultErrors).length > 0
    },
    isBusy () {
      return (this.pending !== null)
    },
    isEmpty () {
      return (!this.isBusy && this.results !== null && this.resultCount === 0)
    },
    isReset () {
      return (!this.isBusy && this.results === null)
    }
  },
  methods: {
    setLimit (limit) {
      this.$store.commit(SET_SEARCH_LIMIT, limit)
      this.$emit('limit-changed', limit)

      this.$store.dispatch(SEARCH_WELLS, { trigger: FILTER_TRIGGER })
    },
    changePage (page) {
      const offset = this.limit * (page - 1)
      this.$store.commit(SET_SEARCH_OFFSET, offset)
      this.$emit('page-changed', page)

      this.$store.dispatch(SEARCH_WELLS, { trigger: FILTER_TRIGGER })
    },
    sortResults ({ param, desc }) {
      const sort = `${desc ? '-' : ''}${param}`
      this.$store.commit(SET_SEARCH_ORDERING, sort)
      this.$emit('sort-changed', sort)

      this.$store.dispatch(SEARCH_WELLS, { trigger: FILTER_TRIGGER })
    },
    applyFilter ({ id }, values) {
      this.filterParams[id] = values
      const filterGroup = { ...this.searchQueryParams }
      this.$store.commit(SET_SEARCH_RESULT_FILTERS, filterGroup)
      this.$emit('filter-changed', filterGroup)

      this.$store.dispatch(SEARCH_WELLS, { trigger: FILTER_TRIGGER })
    },
    initFilterParams () {
      const filterParams = { ...this.emptyFilterParams }

      Object.entries(this.resultFilters).forEach(([param, value]) => {
        this.columns.forEach(column => {
          if (column.params.includes(param)) {
            filterParams[column.id] = { [param]: value }
          }
        })
      })

      this.filterParams = filterParams
    },
    searchResultsRowClicked (data) {
      this.$emit('rowClicked', data)
    }
  },
  filters: {
    streetAddressFormat (row) {
      if (row.city !== undefined && row.city !== null && row.city.toString().trim() !== '') {
        return `${row.street_address}, ${row.city}`
      } else {
        return row.street_address
      }
    },
    selectOptionFormat (value, column, options = null) {
      if (value === undefined || value === null || value === '') {
        return ''
      }

      const valueProp = column.valueField || 'value'
      const textProp = column.textField || 'text'
      const selectOptions = column.options || options
      const optionsMatch = selectOptions.find(option => option[valueProp] === value.toString())
      if (optionsMatch) {
        return optionsMatch[textProp]
      }

      return value
    },
    defaultFormat: (value) => {
      if (value === null) {
        return ''
      }
      return value
    }
  },
  watch: {
    resultColumns (newColumns, oldColumns) {
      const removedColumnIds = oldColumns.filter(colId => !newColumns.includes(colId))
      const removedWithFilters = removedColumnIds.filter(columnId => {
        return this.filterParams[columnId] !== undefined && Object.entries(this.filterParams[columnId]).length > 0
      })
      if (removedWithFilters.length > 1) {
        removedWithFilters.forEach(columnId => {
          this.filterParams[columnId] = {}
        })
      }
      if (removedWithFilters.length > 0) {
        const lastFilterRemoved = removedWithFilters.slice(-1)[0]
        // Call applyFilter to trigger reload
        this.applyFilter({ id: lastFilterRemoved }, {})
      }
    },
    resultFilters (newFilters) {
      // on reset (empty filters), clear state
      if (Object.entries(newFilters).length === 0 && newFilters.constructor === Object) {
        this.clearFilterParams()
      }
    }
  },
  created () {
    this.initFilterParams()
  }
}
</script>

<style lang="scss">

#searchResultsTable {
  min-height: 36rem;

  td.data {
    white-space: wrap;
    overflow-wrap: break-word;
  }

 .sort-button {
    opacity: 0.4;

    &:hover,
    &:active,
    &:visited,
    &:focus {
      text-decoration: none;
    }
  }

  .sort-button.active {
    opacity: 1;
  }

  .vertical-align-middle {
    vertical-align: middle;
  }
}

/* Spinner styles — these can be removed when moving to bootstrap 4.3 */

$spinner-width:         2rem !default;
$spinner-height:        $spinner-width !default;
$spinner-border-width:  .25em !default;

$spinner-width-sm:        1rem !default;
$spinner-height-sm:       $spinner-width-sm !default;
$spinner-border-width-sm: .2em !default;

@keyframes spinner-border {
  to { transform: rotate(360deg); }
}

.spinner-border {
  display: inline-block;
  width: $spinner-width;
  height: $spinner-height;
  vertical-align: text-bottom;
  border: $spinner-border-width solid currentColor;
  border-right-color: transparent;
  // stylelint-disable-next-line property-blacklist
  border-radius: 50%;
  animation: spinner-border .75s linear infinite;
}

.spinner-border-sm {
  width: $spinner-width-sm;
  height: $spinner-height-sm;
  border-width: $spinner-border-width-sm;
}
</style>
