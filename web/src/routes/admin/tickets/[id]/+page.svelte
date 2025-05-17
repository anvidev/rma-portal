<script lang="ts">
	import Billing from '$lib/components/tickets/billing.svelte'
	import Details from '$lib/components/tickets/details.svelte'
	import Logs from '$lib/components/tickets/logs.svelte'
	import Sender from '$lib/components/tickets/sender.svelte'
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
		<div class="flex md:w-1/2 flex-col gap-4">
			<Sender {ticket} />
			<Billing {ticket} />
		</div>
	</div>
	<Logs {ticket} {statuses} superform={form} />
</article>
