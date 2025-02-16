<script lang="ts">
	let { tickets }: { tickets: Ticket[] } = $props()
</script>

<div class="overflow-x-auto">
	<table class="min-w-full divide-y-2 divide-gray-200 text-sm">
		<thead class="ltr:text-left rtl:text-right">
			<tr>
				<th class="px-4 py-2 font-medium whitespace-nowrap">Sags ID </th>
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
					<td class="px-4 py-2 whitespace-nowrap">{ticket.status}</td>
					<td class="space-x-2 px-4 py-2 whitespace-nowrap">
						{#each ticket.categories as category}
							<span
								class="rounded-full border border-purple-600 bg-purple-400 px-1.5 py-1 text-purple-600"
								>{category}</span
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
