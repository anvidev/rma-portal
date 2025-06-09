<script lang="ts">
	import { Clock, Download } from '@lucide/svelte'
	import type { TicketWithLogs } from '$lib/types'
	import { formatDate } from '$lib/utils.js'
	import StatusBadge from '$lib/components/common/status-badge.svelte'
	import CategoryBadge from '$lib/components/common/category-badge.svelte'
	import { Button } from '../ui/button'
	import { toast } from 'svelte-sonner'
	import { page } from '$app/state'

	let { ticket }: { ticket: TicketWithLogs } = $props()

	async function downloadLabel() {
		const id = page.params.id
		const params = new URLSearchParams({ rma: id })
		const labelResponse = await fetch(`/api/label?${params}`)
		if (!labelResponse.ok) {
			const { message } = await labelResponse.json()
			toast.error('Der gik noget galt', { description: message })
			return
		}
		const blob = await labelResponse.blob()
		const url = URL.createObjectURL(blob)
		window.open(url, '_blank')
	}
</script>

<div class="flex items-start justify-between">
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

	<Button size="sm" variant="secondary" onclick={downloadLabel}>
		<Download class="size-3.5" />
		<span class="hidden md:block"> Download RMA Label </span>
	</Button>
</div>
