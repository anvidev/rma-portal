<script lang="ts">
	import { SvelteURLSearchParams } from 'svelte/reactivity'
	import Input from '../ui/input/input.svelte'
	import * as Select from '$lib/components/ui/select/index.js'
	import { stringArrayQueryState, stringQueryState } from '$lib/query-state.svelte'

	let { searchParams }: { searchParams: SvelteURLSearchParams } = $props()
	const query = stringQueryState(searchParams, 'query', { default: '' })
	const status = stringArrayQueryState(searchParams, 'status')
	const categories = stringArrayQueryState(searchParams, 'categories')

	const statusText = $derived(status.value && status.value.join(', '))
	const categoriesText = $derived(categories.value && categories.value.join(', '))
</script>

<div class="flex items-center gap-2">
	<Input class="max-w-52 rounded-lg" placeholder="Søg..." bind:value={query.value} />

	<Select.Root type="multiple" bind:value={status.value}>
		<Select.Trigger class="w-fit min-w-40 rounded-lg capitalize"
			>{statusText ? statusText : 'Status'}</Select.Trigger
		>
		<Select.Content>
			{#each ['åben', 'intern reperation', 'ekstern reperation', 'lukket'] as status}
				<Select.Item class="capitalize" value={status}>{status}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<Select.Root type="multiple" bind:value={categories.value}>
		<Select.Trigger class="w-fit min-w-40 rounded-lg capitalize"
			>{categoriesText ? categoriesText : 'Kategori'}</Select.Trigger
		>
		<Select.Content>
			{#each ['software', 'hardware'] as category}
				<Select.Item class="capitalize" value={category}>{category}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>
</div>
