<script lang="ts">
	import { type QueryState } from '$lib/query-state.svelte'
	import Button from '../ui/button/button.svelte'
	import { Minus, Plus } from '@lucide/svelte'
	import * as Select from '$lib/components/ui/select/index.js'

	let {
		total,
		limit,
		limitOptions,
		pageQuery,
		limitQuery,
	}: {
		total: number
		limit: number
		limitOptions: number[]
		pageQuery: QueryState<number | null>
		limitQuery: QueryState<number | null>
	} = $props()

	const totalPages = $derived(total > 0 ? Math.ceil(total / (limitQuery.value ?? 25)) : 1)
</script>

<div class="flex items-center justify-between">
	<div class="flex items-center gap-4">
		<Select.Root type="single" bind:value={limitQuery.value}>
			<Select.Trigger class="w-20 rounded-lg capitalize"
				>{limitQuery.value ? limitQuery.value : limit}</Select.Trigger
			>
			<Select.Content>
				{#each limitOptions as limit}
					<Select.Item class="capitalize" value={limit.toString()}>{limit}</Select.Item>
				{/each}
			</Select.Content>
		</Select.Root>
		<span class="text-muted-foreground text-sm">
			{total} Total rækker
		</span>
	</div>

	<div class="flex items-center gap-2">
		<span class="text-muted-foreground mr-2 text-sm">
			Side {pageQuery.value ?? 1} af {totalPages}
		</span>
		<Button
			size="icon"
			variant="outline"
			disabled={(pageQuery.value ?? 1) <= 1}
			onclick={() => pageQuery.set((pageQuery.value ?? 1) - 1)}
			aria-label="Forrige side"
		>
			<Minus class="size-3.5" />
		</Button>
		<Button
			size="icon"
			variant="outline"
			disabled={(pageQuery.value ?? 1) >= totalPages}
			onclick={() => pageQuery.set((pageQuery.value ?? 1) + 1)}
			aria-label="Næste side"
		>
			<Plus class="size-3.5" />
		</Button>
	</div>
</div>
