<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js'
	import { page } from '$app/stores'
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import {
		stringQueryState,
		stringArrayQueryState,
		booleanQueryState,
		numberQueryState,
	} from '$lib/query-state.svelte'
	import Input from '$lib/components/ui/input/input.svelte'
	import Checkbox from '$lib/components/ui/checkbox/checkbox.svelte'

	const searchParams = new SvelteURLSearchParams($page.url.searchParams)

	const search = stringQueryState(searchParams, 'q', { default: '' })
	const categories = stringArrayQueryState(searchParams, 'categories')
	const activeOnly = booleanQueryState(searchParams, 'active')
	const age = numberQueryState(searchParams, 'age')
</script>

<Input bind:value={search.value} placeholder="Search..." />

<Select.Root type="multiple" bind:value={categories.value}>
	<Select.Trigger>Vælg</Select.Trigger>
	<Select.Content>
		{#each ['software', 'hardware'] as category}
			<Select.Item value={category}>{category}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>

<label>
	<Checkbox bind:checked={activeOnly.value} />
	Active only
</label>

<Input bind:value={age.value} placeholder="Type age..." />

<pre>
Search: {search.value}
Categories: {JSON.stringify(categories.value)}
Active Only: {activeOnly.value}
Age: {age.value}
</pre>
