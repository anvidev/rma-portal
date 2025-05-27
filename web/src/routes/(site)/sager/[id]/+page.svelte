<script lang="ts">
	import ContactInformation from '$lib/components/tickets/contact-information.svelte'
	import Details from '$lib/components/tickets/details.svelte'
	import LogsListExternal from '$lib/components/tickets/logs-list-external.svelte'
	import Title from '$lib/components/tickets/title.svelte'
	import type { Contact } from '$lib/types.js'

	let { data } = $props()
	const ticket = $derived(data.ticket)

	function censor(s: string): string {
		return s.replaceAll(/./g, '*')
	}

	const censoredSender: Contact = $derived.by(() => {
		return {
			name: censor(ticket.sender.name),
			phone: censor(ticket.sender.phone),
			email: censor(ticket.sender.email),
			street: censor(ticket.sender.street),
			city: censor(ticket.sender.city),
			zip: censor(ticket.sender.zip),
			country: censor(ticket.sender.country),
		}
	})

	const censoredBilling: Contact = $derived.by(() => {
		return {
			name: censor(ticket.billing.name),
			phone: censor(ticket.billing.phone),
			email: censor(ticket.billing.email),
			street: censor(ticket.billing.street),
			city: censor(ticket.billing.city),
			zip: censor(ticket.billing.zip),
			country: censor(ticket.billing.country),
		}
	})
</script>

<title>RMA Sag #{data.ticket.id} - Skancode RMA Service Portal</title>

<article class="space-y-4">
	<Title {ticket} />
	<div class="flex flex-col gap-4 md:flex-row">
		<Details {ticket} />
		<div class="flex flex-col gap-4 md:w-1/2">
			<ContactInformation contact={censoredSender} title="Afsender" sub="Returaddresse" />
			<ContactInformation
				contact={censoredBilling}
				title="Fakturering"
				sub="Faktureringsaddresse"
			/>
		</div>
	</div>
	<LogsListExternal {ticket} />
</article>
