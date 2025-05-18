<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { cn } from '$lib/utils'

	let { ticket }: { ticket: AdminTicket } = $props()

	const warranty = $derived.by(
		() =>
			({
				yes: 'Ja',
				no: 'Nej',
				unknown: 'Ukendt',
			})[ticket.warranty],
	)

	const quote = $derived.by(
		() =>
			({
				yes: 'Ja',
				no: 'Nej',
			})[ticket.quote],
	)
</script>

<div class="space-y-2 rounded-lg border bg-white p-3 shadow-sm md:w-1/2">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">Sags detaljer</h3>
		<p class="text-muted-foreground text-sm">Vare og problem beskrivelse</p>
	</div>
	<div class="flex items-start">
		<div class="w-1/2">
			<small class="text-sm font-medium capitalize leading-tight">Model</small>
			<p class={cn('text-sm', !ticket.model && 'text-muted-foreground italic')}>
				{#if ticket.model}{ticket.model}{:else}Ikke angivet{/if}
			</p>
		</div>
		<div>
			<small class="text-sm font-medium capitalize leading-tight">Serienr.</small>
			<p class={cn('text-sm', !!!ticket.serial_number && 'text-muted-foreground italic')}>
				{#if ticket.serial_number}{ticket.serial_number}{:else}Ikke angivet{/if}
			</p>
		</div>
	</div>
	<div class="flex items-start">
		<div class="w-1/2">
			<small class="text-sm font-medium capitalize leading-tight">Garantidækning</small>
			<p class={cn('text-sm')}>
				{warranty}
			</p>
		</div>
		<div class="w-1/2">
			<small class="text-sm font-medium leading-tight">Tilbud (over minimumspris)</small>
			<p class={cn('text-sm')}>
				{quote}
			</p>
		</div>
	</div>
	<div>
		<small class="text-sm font-medium capitalize leading-tight">Problem beskrivelse</small>
		<p class="text-pretty text-sm">{ticket.issue}</p>
	</div>
</div>
