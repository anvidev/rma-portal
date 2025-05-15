<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { Phone, Mail, MapPin, Building } from '@lucide/svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-2 rounded-lg border bg-white px-3 py-2 shadow-sm">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">Afsender</h3>
		<p class="text-sm text-muted-foreground">Returaddresse</p>
	</div>

	<div>
		<div class="flex items-center gap-2">
			<Building class="size-3.5 text-muted-foreground" />
			<span class="text-sm">{ticket.sender.name}</span>
		</div>
		<div class="flex items-center gap-2">
			<Phone class="size-3.5 text-muted-foreground" />
			<span class="text-sm">{ticket.sender.phone}</span>
		</div>
		<div class="flex items-center gap-2">
			<Mail class="size-3.5 text-muted-foreground" />
			<span class="text-sm">{ticket.sender.email}</span>
		</div>
		<div class="flex items-center gap-2">
			<MapPin class="size-3.5 text-muted-foreground" />
			<span class="text-sm"
				>{ticket.sender.street}, {ticket.sender.zip}
				{ticket.sender.city}, {ticket.sender.country}</span
			>
		</div>
	</div>
</div>
