<script lang="ts">
	import type { Contact } from '$lib/types'
	import { Phone, Mail, MapPin, Building } from '@lucide/svelte'

	let { contact, title, sub }: { contact: Contact; title: string; sub: string } = $props()
</script>

<div class="space-y-2 rounded-lg border bg-white p-3 shadow-sm">
	<div>
		<h3 class="font-4xl font-semibold leading-tight">{title}</h3>
		<p class="text-sm text-muted-foreground">{sub}</p>
	</div>

	<div>
		<div class="flex items-center gap-2">
			<Building class="size-4 text-muted-foreground" />
			<span class="text-sm">{contact.name}</span>
		</div>
		<div class="flex items-center gap-2">
			<Phone class="size-4 text-muted-foreground" />
			<span class="text-sm">{contact.phone}</span>
		</div>
		<div class="flex items-center gap-2">
			<Mail class="size-4 text-muted-foreground" />
			<span class="text-sm">{contact.email}</span>
		</div>
		<div class="flex items-center gap-2">
			<MapPin class="size-4 text-muted-foreground" />
			<span class="text-sm"
				>{contact.street}, {contact.zip}
				{contact.city}, {contact.country}</span
			>
		</div>
	</div>
</div>
