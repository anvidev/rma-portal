<script lang="ts">
	import * as Table from '$lib/components/ui/table/index.js'
	import { goto } from '$app/navigation'
	import type { SvelteURLSearchParams } from 'svelte/reactivity'
	import StatusBadge from '../common/status-badge.svelte'
	import CategoryBadge from '../common/category-badge.svelte'
	import type { Ticket } from '$lib/types'
	import { stringArrayQueryState, type QueryState } from '$lib/query-state.svelte'

	let {
		tickets,
		searchParams,
		statusQuery,
		categoriesQuery,
	}: {
		tickets: Ticket[]
		searchParams: SvelteURLSearchParams
		statusQuery: QueryState<string[]>
		categoriesQuery: QueryState<string[]>
	} = $props()

	function handleSortingChange(key: string, direction: 'asc' | 'desc') {
		searchParams.set('sorting', key)
		searchParams.set('direction', direction)
		goto(`?${searchParams.toString()}`)
	}
</script>

<div class="rounded-lg border bg-white shadow-sm">
	<Table.Root>
		<Table.Caption class="m-0 p-4">Oversigt over alle RMA sager</Table.Caption>
		<Table.Header>
			<Table.Row>
				<Table.Head onclick={() => handleSortingChange('id', 'asc')}>Sags ID</Table.Head>
				<Table.Head>Status</Table.Head>
				<Table.Head>Kategorier</Table.Head>
				<Table.Head class="max-w-60 truncate">Afsender firma</Table.Head>
				<Table.Head>Afsender e-mail</Table.Head>
				<Table.Head>Oprettet</Table.Head>
				<Table.Head>Sidste ændring</Table.Head>
			</Table.Row>
		</Table.Header>
		<Table.Body class="border-b">
			{#each tickets as ticket (ticket.id)}
				<Table.Row>
					<Table.Cell class="font-medium">
						<a href={`/admin/tickets/${ticket.id}`}>#{ticket.id}</a>
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
					<Table.Cell class="max-w-60 truncate">{ticket.sender.name}</Table.Cell>
					<Table.Cell>{ticket.sender.email}</Table.Cell>
					<Table.Cell>{new Date(ticket.inserted).toLocaleString('en-GB')}</Table.Cell>
					<Table.Cell>{new Date(ticket.updated).toLocaleString('en-GB')}</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
