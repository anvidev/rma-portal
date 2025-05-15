<script lang="ts">
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import Input from '../ui/input/input.svelte'
	import * as Select from '$lib/components/ui/select/index.js'
	import { stringArrayQueryState, stringQueryState } from '$lib/query-state.svelte'

	let { searchParams }: { searchParams: SvelteURLSearchParams } = $props()
	const query = stringQueryState(searchParams, 'query', { default: '' })
	const status = stringQueryState(searchParams, 'status')
	const categories = stringArrayQueryState(searchParams, 'categories')
</script>

<!-- <input -->
<!-- 	class="h-9 rounded-sm border border-slate-700 px-2" -->
<!-- 	placeholder="Søg efter sager..." -->
<!-- 	type="text" -->
<!-- 	bind:value={query.value} -->
<!-- /> -->

<Input bind:value={query.value} />

<!-- <select class="h-9 rounded-sm border border-slate-700 px-2" bind:value={status.value}> -->
<!-- 	<option value="">Alle</option> -->
<!-- 	<option value="åben">Åben</option> -->
<!-- 	<option value="lukket">Lukket</option> -->
<!-- </select> -->

<Select.Root type="multiple" bind:value={status.value}>
	<Select.Trigger>Vælg</Select.Trigger>
	<Select.Content>
		{#each ['åben', 'lukket'] as status}
			<Select.Item value={status}>{status}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>

<!-- <select class="h-9 rounded-sm border border-slate-700 px-2" multiple bind:value={status.value}> -->
<!-- 	{#each ['software', 'hardware'] as category} -->
<!-- 		<option value={category}>{category}</option> -->
<!-- 	{/each} -->
<!-- </select> -->

<Select.Root type="multiple" bind:value={categories.value}>
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
