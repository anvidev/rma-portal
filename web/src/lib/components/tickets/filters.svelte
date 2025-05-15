<script lang="ts">
	import { goto } from '$app/navigation'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import { URLParams } from '$lib/state/url.svelte'
	import Input from '../ui/input/input.svelte'
	import { page } from '$app/state'
	import * as Select from '$lib/components/ui/select/index.js'

	let { searchParams }: { searchParams: SvelteURLSearchParams } = $props()

	let query = $state(searchParams.get('query') || '')
	let status = $state(searchParams.get('status') || '')
	let categories = $state(searchParams.getAll('categories') || [])
	let debounceTimer: ReturnType<typeof setTimeout>

	function handleFilterChange(key: string, value: any, waitMS: number = 300) {
		clearTimeout(debounceTimer)

		if (Array.isArray(value) && value.length == 0) {
			searchParams.delete(key)
		} else if (typeof value == 'string' && value == '') {
			searchParams.delete(key)
		} else {
			searchParams.set(key, value)
		}

		debounceTimer = setTimeout(() => {
			goto(`?${searchParams.toString()}`, { keepFocus: true })
		}, waitMS)
	}

	const urlParams = new URLParams(page.url.searchParams)
</script>

<input
	class="h-9 rounded-sm border border-slate-700 px-2"
	placeholder="Søg efter sager..."
	type="text"
	bind:value={query}
	onkeyup={() => handleFilterChange('query', query)}
/>

<Input bind:value={urlParams.params.query} />

<select class="h-9 rounded-sm border border-slate-700 px-2" bind:value={urlParams.params.status}>
	<option value="">Alle</option>
	<option value="åben">Åben</option>
	<option value="lukket">Lukket</option>
</select>

<select
	class="h-9 rounded-sm border border-slate-700 px-2"
	multiple
	bind:value={urlParams.params.categories}
>
	{#each ['software', 'hardware'] as category}
		<option value={category}>{category}</option>
	{/each}
</select>

<Select.Root type="multiple" bind:value={urlParams.params.categories}>
	<Select.Trigger>Vælg</Select.Trigger>
	<Select.Content>
		{#each ['software', 'hardware'] as category}
			<Select.Item value={category}>{category}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>

<!-- <select -->
<!-- 	class="h-9 rounded-sm border border-slate-700 px-2" -->
<!-- 	bind:value={status} -->
<!-- 	onchange={() => handleFilterChange('status', status, 0)} -->
<!-- > -->
<!-- 	<option value="">Alle</option> -->
<!-- 	<option value="åben">Åben</option> -->
<!-- 	<option value="lukket">Lukket</option> -->
<!-- </select> -->

<!-- <div> -->
<!-- 	{#each ['software', 'hardware'] as category} -->
<!-- 		<label> -->
<!-- 			<input -->
<!-- 				type="checkbox" -->
<!-- 				name="flavours" -->
<!-- 				value={category} -->
<!-- 				bind:group={categories} -->
<!-- 				onchange={() => handleFilterChange('categories', categories)} -->
<!-- 			/> -->
<!-- 			{category} -->
<!-- 		</label> -->
<!-- 	{/each} -->
<!-- </div> -->
