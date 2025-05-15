<script lang="ts">
	import Filters from '$lib/components/tickets/filters.svelte'
	import Table from '$lib/components/tickets/table.svelte'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import { page } from '$app/state'

	let { data } = $props()

	let searchParams = new SvelteURLSearchParams(page.url.searchParams)
</script>

<h1>Oversigt over sager</h1>
<Filters {searchParams} />
<Table {searchParams} tickets={data.tickets} />
