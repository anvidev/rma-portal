<script lang="ts">
	import Table from '$lib/components/tickets/table.svelte'
	import Filters from '$lib/components/tickets/filters.svelte'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import { page } from '$app/state'
	import Pagination from '$lib/components/tickets/pagination.svelte'

	let { data } = $props()

	const searchParams = new SvelteURLSearchParams(page.url.searchParams)
</script>

<section class="space-y-4">
	<div>
		<h1 class="text-lg font-semibold leading-tight">Sags oversigt</h1>
		<p class="text-sm text-muted-foreground">
			Se alle dine RMA-sager samlet ét sted, og brug filtre til hurtigt at finde det, du søger
		</p>
	</div>
	<Filters {searchParams} />
	<Table {searchParams} tickets={data.tickets} />
	<Pagination {searchParams} total={data.total} />
</section>
