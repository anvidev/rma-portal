<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { Clock } from '@lucide/svelte'
	import type { AdminTicket } from '$lib/types'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-3">
	<h1 class="text-2xl font-medium leading-tight">Sag #{ticket.id}</h1>
	<div class="flex items-center gap-1">
		<Badge class="bg-blue-600 capitalize">{ticket.status}</Badge>
		{#each ticket.categories as category}
			<Badge variant="secondary" class="capitalize">{category}</Badge>
		{/each}
	</div>
	<div>
		<div class="flex items-center gap-2">
			<Clock class="text-muted-foreground size-4" />
			<span class="text-muted-foreground">Oprettet: {ticket.inserted}</span>
		</div>
		<div class="flex items-center gap-2">
			<Clock class="text-muted-foreground size-4" />
			<span class="text-muted-foreground">Sidst opdateret: {ticket.updated}</span>
		</div>
	</div>
</div>
