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

	let searchQuery = $state(stringQueryState(searchParams, 'query', { default: '' }))
	let statusQuery = $state(stringArrayQueryState(searchParams, 'status'))
	let categoriesQuery = $state(stringArrayQueryState(searchParams, 'categories'))
	let pageQuery = $state(numberQueryState(searchParams, 'page'))
	let limitQuery = $state(numberQueryState(searchParams, 'limit'))
	let sortQuery = $state(stringQueryState(searchParams, 'sort'))
	let directionQuery = $state(stringQueryState(searchParams, 'direction'))

	function clearFilters() {
		searchQuery.set('')
		statusQuery.set([])
		categoriesQuery.set([])
	}
</script>

<svelte:head>
	<meta name="robots" content="noindex, nofollow" />
	<title>RMA Sags Oversigt - Skancode RMA Service Portal</title>
</svelte:head>

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
		bind:statusQuery
		bind:searchQuery
		bind:categoriesQuery
		{clearFilters}
	/>
	<Table
		{searchParams}
		tickets={data.tickets}
		bind:statusQuery
		bind:categoriesQuery
		bind:sortQuery
		bind:directionQuery
	/>
	<Pagination total={data.total} limit={data.limit} {limitOptions} bind:pageQuery bind:limitQuery />
</section>
