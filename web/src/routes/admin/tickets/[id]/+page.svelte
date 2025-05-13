<script lang="ts">
	import Billing from '$lib/components/admin/tickets/billing.svelte'
	import Details from '$lib/components/admin/tickets/details.svelte'
	import Logs from '$lib/components/admin/tickets/logs.svelte'
	import Sender from '$lib/components/admin/tickets/sender.svelte'
	import Title from '$lib/components/admin/tickets/title.svelte'

	let { data } = $props()
	const ticket = data.ticket
</script>

<article class="space-y-4">
	<Title {ticket} />
	<div class="flex flex-col gap-4 md:flex-row">
		<Details {ticket} />
		<div class="flex w-1/2 flex-col gap-4">
			<Sender {ticket} />
			<Billing {ticket} />
		</div>
	</div>
	<Logs {ticket} />
</article>

<pre>{JSON.stringify(data.ticket, null, 2)}</pre>
