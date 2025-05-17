<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { cn } from '$lib/utils'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="md:w-1/2 space-y-2 rounded-lg border bg-white p-3 shadow-sm">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">Sags detaljer</h3>
		<p class="text-sm text-muted-foreground">Vare og problem beskrivelse</p>
	</div>
	<div>
		<small class="text-sm font-medium capitalize leading-tight">Model</small>
		<p class={cn('text-sm', !!!ticket.model && 'italic text-muted-foreground')}>
			{#if ticket.model}{ticket.model}{:else}Ikke angivet{/if}
		</p>
	</div>
	<div>
		<small class="text-sm font-medium capitalize leading-tight">Serienr.</small>
		<p class={cn('text-sm', !!!ticket.serial_number && 'italic text-muted-foreground')}>
			{#if ticket.serial_number}{ticket.serial_number}{:else}Ikke angivet{/if}
		</p>
	</div>
	<div>
		<small class="text-sm font-medium capitalize leading-tight">Problem beskrivelse</small>
		<p class="text-pretty text-sm">{ticket.issue}</p>
	</div>
</div>
