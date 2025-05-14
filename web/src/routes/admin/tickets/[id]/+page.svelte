<script lang="ts">
	import Billing from '$lib/components/admin/tickets/billing.svelte'
	import Details from '$lib/components/admin/tickets/details.svelte'
	import Logs from '$lib/components/admin/tickets/logs.svelte'
	import Sender from '$lib/components/admin/tickets/sender.svelte'
	import Title from '$lib/components/admin/tickets/title.svelte'
	import { Shield } from '@lucide/svelte'
	import { formatDate } from '$lib/utils.js'

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
	<!-- <Logs {ticket} /> -->

	<div class="">
		<h3 class="font-4xl font-semibold leading-tight">Seneste opdateringer</h3>
		<p class="text-muted-foreground text-sm">Følg status og opdateringer på din RMA sag</p>
	</div>
	<div class="mt-8 flex flex-col gap-4">
		{#each ticket?.logs as log (log.id)}
			<div class="rounded-lg border border-l-[3px] border-l-sky-700 bg-white p-3 shadow-sm">
				<div class="space-y-2">
					<div class="flex items-center justify-between">
						<p class="text-sm font-medium capitalize leading-tight">{log.status}</p>
						<p class="text-sm leading-tight">{formatDate(log.inserted)}</p>
					</div>
					<small class="text-muted-foreground">{log.initiator}</small>
					<p class="text-sm">{log.external_comment}</p>
				</div>
				{#if log.internal_comment}
					<div class="bg-muted/40 mt-3 rounded-lg border p-3">
						<div class="text-muted-foreground mb-1 flex items-center gap-1 text-sm">
							<Shield class="size-3.5" />
							<span class="font-medium">Intern kommentar</span>
						</div>
						<p class="text-sm">{log.internal_comment}</p>
					</div>
				{/if}
			</div>
		{/each}
	</div>
</article>

<!-- <pre>{JSON.stringify(data.ticket, null, 2)}</pre> -->
