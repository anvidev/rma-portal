<script lang="ts">
	import { numberQueryState } from '$lib/query-state.svelte'
	import type { SvelteURLSearchParams } from 'svelte/reactivity'
	import Button from '../ui/button/button.svelte'
	import { Minus, Plus } from '@lucide/svelte'
	import * as Select from '$lib/components/ui/select/index.js'

	let { searchParams, total }: { searchParams: SvelteURLSearchParams; total: number } = $props()

	const page = numberQueryState(searchParams, 'page', { default: 1 })
	const limit = numberQueryState(searchParams, 'limit', { default: 25 })
	const totalPages = $derived(Math.max(Math.floor(total / (limit.value ?? 25)), 1))
</script>

<div class="flex items-center justify-between">
	<div>
		<Select.Root type="single" bind:value={limit.value}>
			<Select.Trigger class="w-20 rounded-lg capitalize"
				>{limit.value ? limit.value : '50'}</Select.Trigger
			>
			<Select.Content>
				{#each ['25', '50', '100'] as limit}
					<Select.Item class="capitalize" value={limit}>{limit}</Select.Item>
				{/each}
			</Select.Content>
		</Select.Root>
	</div>

	<div class="flex items-center gap-2">
		<Button
			size="icon"
			variant="outline"
			disabled={page.value <= 1}
			onclick={() => page.decrement()}
		>
			<Minus class="size-3.5" />
		</Button>
		<Button
			size="icon"
			variant="outline"
			disabled={page.value == totalPages}
			onclick={() => page.increment()}
		>
			<Plus class="size-3.5" />
		</Button>
	</div>
</div>
