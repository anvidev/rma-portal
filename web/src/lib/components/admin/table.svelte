<script lang="ts">
	import { goto } from '$app/navigation'
	import type { SvelteURLSearchParams } from 'svelte/reactivity'

	let { tickets, searchParams }: { tickets: Ticket[]; searchParams: SvelteURLSearchParams } =
		$props()

	function handleSortingChange(key: string, direction: 'asc' | 'desc') {
		searchParams.set('sorting', key)
		searchParams.set('direction', direction)
		goto(`?${searchParams.toString()}`)
	}
</script>

<div class="overflow-x-auto">
	<table class="min-w-full divide-y-2 divide-gray-200 text-sm">
		<thead class="ltr:text-left rtl:text-right">
			<tr>
				<th
					onclick={() => handleSortingChange('id', 'asc')}
					class="cursor-pointer px-4 py-2 font-medium whitespace-nowrap"
					>Sags ID
				</th>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Status</th>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Kategorier</th>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Afsender</th>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Oprettet</th>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Opdateret</th>
			</tr>
		</thead>

		<tbody class="divide-y divide-gray-200">
			{#each tickets as ticket (ticket.id)}
				<tr>
					<td class="px-4 py-2 font-medium whitespace-nowrap">
						<a href={`/admin/tickets/${ticket.id}`}>{ticket.id}</a>
					</td>
					<td class="px-4 py-2 whitespace-nowrap capitalize">{ticket.status}</td>
					<td class="space-x-2 px-4 py-2 whitespace-nowrap">
						{#each ticket.categories as category}
							<span
								class="rounded-sm border border-purple-600 bg-purple-950 px-1.5 py-1 text-purple-600 uppercase"
								>{category.substring(0, 1)}</span
							>
						{/each}
					</td>
					<td class="px-4 py-2 whitespace-nowrap">{ticket.sender.name}</td>
					<td class="px-4 py-2 whitespace-nowrap"
						>{new Date(ticket.inserted).toLocaleString('en-GB')}</td
					>
					<td class="px-4 py-2 whitespace-nowrap"
						>{new Date(ticket.updated).toLocaleString('en-GB')}</td
					>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
