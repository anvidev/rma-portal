<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { Phone, Mail, MapPin, Building } from '@lucide/svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-2 rounded-lg border bg-white px-3 py-2 shadow-sm">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">Fakturering</h3>
		<p class="text-muted-foreground text-sm">Faktureringsaddresse</p>
	</div>

	<div>
		<div class="flex items-center gap-2">
			<Building class="text-muted-foreground size-4" />
			<span class="">{ticket.billing.name}</span>
		</div>
		<div class="flex items-center gap-2">
			<Phone class="text-muted-foreground size-4" />
			<span class="">{ticket.billing.phone}</span>
		</div>
		<div class="flex items-center gap-2">
			<Mail class="text-muted-foreground size-4" />
			<span class="">{ticket.billing.email}</span>
		</div>
		<div class="flex items-center gap-2">
			<MapPin class="text-muted-foreground size-4" />
			<span class=""
				>{ticket.billing.street}, {ticket.billing.zip}
				{ticket.billing.city}, {ticket.billing.country}</span
			>
		</div>
	</div>
</div>
