<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { Clock } from '@lucide/svelte'
	import type { AdminTicket } from '$lib/types'
	import { cn, formatDate } from '$lib/utils.js'
	import StatusBadge from '$lib/components/common/status-badge.svelte'
	import CategoryBadge from '$lib/components/common/category-badge.svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-3">
	<h1 class="text-2xl font-semibold leading-tight">Sag #{ticket.id}</h1>
	<div class="flex items-center gap-1">
		<StatusBadge status={ticket.status} />
		{#each ticket.categories as category}
			<CategoryBadge {category} />
		{/each}
	</div>
	<div>
		<div class="flex items-center gap-2">
			<Clock class="size-3.5 text-muted-foreground" />
			<span class="text-sm text-muted-foreground">Oprettet: {formatDate(ticket.inserted)}</span>
		</div>
		<div class="flex items-center gap-2">
			<Clock class="size-3.5 text-muted-foreground" />
			<span class="text-sm text-muted-foreground"
				>Sidst opdateret: {formatDate(ticket.updated)}</span
			>
		</div>
	</div>
</div>
