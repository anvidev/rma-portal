<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { Phone, Mail, MapPin, Building } from '@lucide/svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-2 rounded-lg border bg-white px-3 py-2 shadow-sm">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">Afsender</h3>
		<p class="text-muted-foreground text-sm">Returaddresse</p>
	</div>

	<div>
		<div class="flex items-center gap-2">
			<Building class="text-muted-foreground size-3.5" />
			<span class="">{ticket.sender.name}</span>
		</div>
		<div class="flex items-center gap-2">
			<Phone class="text-muted-foreground size-3.5" />
			<span class="">{ticket.sender.phone}</span>
		</div>
		<div class="flex items-center gap-2">
			<Mail class="text-muted-foreground size-3.5" />
			<span class="">{ticket.sender.email}</span>
		</div>
		<div class="flex items-center gap-2">
			<MapPin class="text-muted-foreground size-3.5" />
			<span class=""
				>{ticket.sender.street}, {ticket.sender.zip}
				{ticket.sender.city}, {ticket.sender.country}</span
			>
		</div>
	</div>
</div>
