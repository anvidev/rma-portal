<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { Clock } from '@lucide/svelte'
	import type { AdminTicket } from '$lib/types'
	import { formatDate } from '$lib/utils.js'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-3">
	<h1 class="text-2xl font-semibold leading-tight">Sag #{ticket.id}</h1>
	<div class="flex items-center gap-1">
		<Badge class="flex items-center gap-2 bg-sky-100 capitalize text-sky-700">
			<span class="block size-1.5 rounded-full bg-sky-700"></span>
			{ticket.status}</Badge
		>
		{#each ticket.categories as category}
			<Badge variant="secondary" class="flex items-center gap-2 capitalize">
				<span class="bg-secondary-foreground block size-1.5 rounded-full"></span>
				{category}</Badge
			>
		{/each}
	</div>
	<div>
		<div class="flex items-center gap-2">
			<Clock class="text-muted-foreground size-3.5" />
			<span class="text-muted-foreground text-sm">Oprettet: {formatDate(ticket.inserted)}</span>
		</div>
		<div class="flex items-center gap-2">
			<Clock class="text-muted-foreground size-3.5" />
			<span class="text-muted-foreground text-sm"
				>Sidst opdateret: {formatDate(ticket.updated)}</span
			>
		</div>
	</div>
</div>
