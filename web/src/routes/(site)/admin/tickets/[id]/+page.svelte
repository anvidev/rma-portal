<script lang="ts">
	import ContactInformation from '$lib/components/tickets/contact-information.svelte'
	import Details from '$lib/components/tickets/details.svelte'
	import Logs from '$lib/components/tickets/logs.svelte'
	import Title from '$lib/components/tickets/title.svelte'

	let { data } = $props()
	const ticket = $derived(data.ticket)
	const statuses = data.statuses
	const form = data.form
</script>

<article class="space-y-4">
	<Title {ticket} />
	<div class="flex flex-col gap-4 md:flex-row">
		<Details {ticket} />
		<div class="flex flex-col gap-4 md:w-1/2">
			<ContactInformation contact={ticket.sender} title="Afsender" sub="Returaddresse" />
			<ContactInformation contact={ticket.billing} title="Fakturering" sub="Faktureringsaddresse" />
		</div>
	</div>
	<Logs {ticket} {statuses} superform={form} />
</article>
