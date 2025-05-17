<script lang="ts">
	import Input from '../ui/input/input.svelte'
	import * as Select from '$lib/components/ui/select/index.js'
	import { type QueryState } from '$lib/query-state.svelte'
	import Button from '../ui/button/button.svelte'

	let {
		statuses,
		categories,
		statusQuery,
		searchQuery,
		categoriesQuery,
		clearFilters,
	}: {
		statuses: string[]
		categories: string[]
		statusQuery: QueryState<string[]>
		categoriesQuery: QueryState<string[]>
		searchQuery: QueryState<string>
		clearFilters: () => void
	} = $props()

	const statusText = $derived(statusQuery.value && statusQuery.value.join(', '))
	const categoriesText = $derived(categoriesQuery.value && categoriesQuery.value.join(', '))
</script>

<div class="flex items-center gap-2">
	<Input class="max-w-52 rounded-lg" placeholder="Søg..." bind:value={searchQuery.value} />

	<Select.Root type="multiple" bind:value={statusQuery.value}>
		<Select.Trigger class="w-fit min-w-40 rounded-lg capitalize"
			>{statusText ? statusText : 'Status'}</Select.Trigger
		>
		<Select.Content>
			{#each statuses as status}
				<Select.Item class="capitalize" value={status}>{status}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<Select.Root type="multiple" bind:value={categoriesQuery.value}>
		<Select.Trigger class="w-fit min-w-40 rounded-lg capitalize"
			>{categoriesText ? categoriesText : 'Kategori'}</Select.Trigger
		>
		<Select.Content>
			{#each categories as category}
				<Select.Item class="capitalize" value={category}>{category}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<Button onclick={clearFilters} variant="outline">Nulstil</Button>
</div>
