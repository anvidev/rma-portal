<script lang="ts">
	import Filters from '$lib/components/admin/filters.svelte'
	import Table from '$lib/components/admin/table.svelte'

	let { data } = $props()
</script>

<h1>Sager</h1>
<Filters />
<Table tickets={data.tickets} />
