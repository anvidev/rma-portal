<script lang="ts">
	import * as Table from '$lib/components/ui/table/index.js'
	import StatusBadge from '../common/status-badge.svelte'
	import CategoryBadge from '../common/category-badge.svelte'
	import type { Ticket } from '$lib/types'
	import { type QueryState } from '$lib/query-state.svelte'
	import { formatDate } from '$lib/utils'
	import { ArrowDown, ArrowUp } from '@lucide/svelte'

	let {
		tickets,
		statusQuery = $bindable(),
		categoriesQuery = $bindable(),
		sortQuery = $bindable(),
		directionQuery = $bindable(),
	}: {
		tickets: Ticket[]
		statusQuery: QueryState<string[]>
		categoriesQuery: QueryState<string[]>
		sortQuery: QueryState<string>
		directionQuery: QueryState<string>
	} = $props()

	function handleSortingChange(key: string) {
		if (sortQuery.value === key) {
			if (directionQuery.value === undefined) {
				directionQuery.set('desc')
			} else if (directionQuery.value === 'desc') {
				directionQuery.set('asc')
			} else {
				directionQuery.set(undefined)
				sortQuery.set(undefined)
			}
		} else {
			sortQuery.set(key)
			directionQuery.set('desc')
		}
	}
</script>

{#snippet sortingArrows(
	key: string,
	sortQuery: QueryState<string>,
	directionQuery: QueryState<string>,
)}
	{#if sortQuery.value === key}
		{#if directionQuery.value === 'desc'}
			<ArrowDown class="size-4" />
		{:else if directionQuery.value === 'asc'}
			<ArrowUp class="size-4" />
		{/if}
	{/if}
{/snippet}

<div class="rounded-lg border bg-white shadow-sm">
	<Table.Root>
		<Table.Caption class="m-0 p-4">Oversigt over alle RMA sager</Table.Caption>
		<Table.Header>
			<Table.Row>
				<Table.Head onclick={() => handleSortingChange('id')} class="group cursor-pointer">
					<div class="flex items-center gap-1">
						<span class="inline-block group-hover:underline">Sags ID</span>
						{@render sortingArrows('id', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head onclick={() => handleSortingChange('status')} class="group cursor-pointer">
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Status</span>
						{@render sortingArrows('status', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head onclick={() => handleSortingChange('categories')} class="group cursor-pointer">
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Kategorier</span>
						{@render sortingArrows('categories', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head
					class="group max-w-44 cursor-pointer truncate"
					onclick={() => handleSortingChange('sender_company')}
				>
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Firma</span>
						{@render sortingArrows('sender_company', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head
					class="group max-w-44 cursor-pointer truncate"
					onclick={() => handleSortingChange('sender_name')}
				>
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Kontaktperson</span>
						{@render sortingArrows('sender_name', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head onclick={() => handleSortingChange('inserted')} class="group cursor-pointer">
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Oprettet</span>
						{@render sortingArrows('inserted', sortQuery, directionQuery)}
					</div>
				</Table.Head>
				<Table.Head onclick={() => handleSortingChange('updated')} class="group cursor-pointer">
					<div class="flex items-center gap-1">
						<span class="group-hover:underline">Sidste ændring</span>
						{@render sortingArrows('updated', sortQuery, directionQuery)}
					</div>
				</Table.Head>
			</Table.Row>
		</Table.Header>
		<Table.Body class="border-b">
			{#each tickets as ticket (ticket.id)}
				<Table.Row>
					<Table.Cell class="font-medium">
						<a href={`/admin/sager/${ticket.id}`}>#{ticket.id.split('-')[0]}</a>
					</Table.Cell>
					<Table.Cell
						><StatusBadge
							onclick={() => statusQuery.set([ticket.status])}
							status={ticket.status}
						/></Table.Cell
					>
					<Table.Cell class="flex items-center gap-1">
						{#each ticket.categories as category}
							<CategoryBadge onclick={() => categoriesQuery.set([category])} {category} />
						{/each}
					</Table.Cell>
					<Table.Cell class="max-w-44 truncate">{ticket.sender.company}</Table.Cell>
					<Table.Cell class="max-w-44 truncate">{ticket.sender.name}</Table.Cell>
					<Table.Cell>{formatDate(ticket.inserted)}</Table.Cell>
					<Table.Cell>{formatDate(ticket.updated)}</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
