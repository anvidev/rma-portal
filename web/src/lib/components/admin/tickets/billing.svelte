<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { Phone, Mail, MapPin } from '@lucide/svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-1 rounded-md border bg-white px-3 py-2">
	<h3 class="text-lg font-medium leading-tight">Fakturering</h3>
	<small class="text-muted-foreground text-sm font-medium">{ticket.billing.name}</small>

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
