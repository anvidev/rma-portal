<script lang="ts">
	import Table from '$lib/components/tickets/table.svelte'
	import Filters from '$lib/components/tickets/filters.svelte'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import { page } from '$app/state'
	import Pagination from '$lib/components/tickets/pagination.svelte'
	import {
		numberQueryState,
		stringArrayQueryState,
		stringQueryState,
	} from '$lib/query-state.svelte.js'

	let { data } = $props()
	let limitOptions = [25, 50, 100]

	const searchParams = new SvelteURLSearchParams(page.url.searchParams)

	const searchQuery = stringQueryState(searchParams, 'query', { default: '' })
	const statusQuery = stringArrayQueryState(searchParams, 'status')
	const categoriesQuery = stringArrayQueryState(searchParams, 'categories')
	const pageQuery = numberQueryState(searchParams, 'page')
	const limitQuery = numberQueryState(searchParams, 'limit')

	function clearFilters() {
		searchQuery.set('')
		statusQuery.set([])
		categoriesQuery.set([])
	}
</script>

<section class="space-y-4">
	<div>
		<h1 class="text-lg font-semibold leading-tight">Sags oversigt</h1>
		<p class="text-muted-foreground text-sm">
			Se alle dine RMA-sager samlet ét sted, og brug filtre til hurtigt at finde det, du søger
		</p>
	</div>
	<Filters
		statuses={data.statuses}
		categories={data.categories}
		{statusQuery}
		{searchQuery}
		{categoriesQuery}
		{clearFilters}
	/>
	<Table {searchParams} tickets={data.tickets} {statusQuery} {categoriesQuery} />
	<Pagination total={data.total} limit={data.limit} {limitOptions} {pageQuery} {limitQuery} />
</section>
