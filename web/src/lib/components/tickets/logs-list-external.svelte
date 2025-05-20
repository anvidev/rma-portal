<script lang="ts">
	import type { TicketWithLogs } from '$lib/types'
	import LogsItem from './logs-item.svelte'

	let {
		ticket,
	}: {
		ticket: TicketWithLogs
	} = $props()
</script>

<div class="flex items-center justify-between">
	<div>
		<h3 class="font-semibold leading-tight">Seneste opdateringer</h3>
		<p class="text-sm text-muted-foreground">Følg status og opdateringer på din RMA sag</p>
	</div>
</div>
<div class="mt-8 flex flex-col gap-4">
	{#each ticket.logs as log (log.id)}
		<LogsItem {log} />
	{/each}
</div>
