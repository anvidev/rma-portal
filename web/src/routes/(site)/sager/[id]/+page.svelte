<script lang="ts">
	import ContactInformation from '$lib/components/tickets/contact-information.svelte'
	import Details from '$lib/components/tickets/details.svelte'
	import LogsListExternal from '$lib/components/tickets/logs-list-external.svelte'
	import Title from '$lib/components/tickets/title.svelte'

	let { data } = $props()
	const ticket = $derived(data.ticket)
</script>

<title>RMA Sag #{data.ticket.id} - Skancode RMA Service Portal</title>

<article class="space-y-4">
	<Title {ticket} />
	<div class="flex flex-col gap-4 md:flex-row">
		<Details {ticket} />
		<div class="flex flex-col gap-4 md:w-1/2">
			<ContactInformation contact={ticket.sender} title="Afsender" sub="Returaddresse" />
			<ContactInformation contact={ticket.billing} title="Fakturering" sub="Faktureringsaddresse" />
		</div>
	</div>
	<LogsListExternal {ticket} />
</article>
