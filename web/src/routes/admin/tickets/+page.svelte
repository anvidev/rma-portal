<script lang="ts">
	import Table from '$lib/components/tickets/table.svelte'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import { page } from '$app/stores'
	import { stringArrayQueryState, stringQueryState } from '$lib/query-state.svelte'
	import Input from '$lib/components/ui/input/input.svelte'
	import * as Select from '$lib/components/ui/select/index.js'

	let { data } = $props()

	const searchParams = new SvelteURLSearchParams($page.url.searchParams)

	const search = stringQueryState(searchParams, 'q', { default: '' })
	const status = stringQueryState(searchParams, 'status')
	const categories = stringArrayQueryState(searchParams, 'categories')
</script>

<h1>Oversigt over sager</h1>

<Input bind:value={search.value} placeholder="Search..." />

<Select.Root type="multiple" bind:value={categories.value}>
	<Select.Trigger>Vælg</Select.Trigger>
	<Select.Content>
		{#each ['software', 'hardware'] as category}
			<Select.Item value={category}>{category}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>

<Table {searchParams} tickets={data.tickets} />
