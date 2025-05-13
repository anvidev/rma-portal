<script lang="ts">
	import type { AdminTicket } from '$lib/types'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="w-1/2 space-y-1 rounded-md border bg-white px-3 py-2">
	<h3 class="text-lg font-medium leading-tight">Sags detaljer</h3>
	<div>
		<small class="text-muted-foreground text-sm font-medium">Model</small>
		<p>
			{#if ticket.model}{ticket.model}{:else}Ikke angivet{/if}
		</p>
	</div>
	<div>
		<small class="text-muted-foreground text-sm font-medium">Serienr.</small>
		<p>
			{#if ticket.serial_number}{ticket.serial_number}{:else}Ikke angivet{/if}
		</p>
	</div>
	<div>
		<small class="text-muted-foreground text-sm font-medium">Problem beskrivelse</small>
		<p>{ticket.issue}</p>
	</div>
</div>
