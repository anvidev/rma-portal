<script lang="ts">
	import Filters from '$lib/components/admin/filters.svelte'
	import Table from '$lib/components/admin/table.svelte'
	import {SvelteURLSearchParams} from 'svelte/reactivity'
	import {page} from '$app/state'

	let {data} = $props()

	let searchParams = new SvelteURLSearchParams(page.url.searchParams)
</script>

<h1>Oversigt over sager</h1>
<Filters {searchParams} />
<Table {searchParams} tickets={data.tickets} />
