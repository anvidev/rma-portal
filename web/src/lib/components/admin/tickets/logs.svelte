<script lang="ts">
	import type { AdminTicket } from '$lib/types'
	import { EyeOff } from '@lucide/svelte'

	let { ticket }: { ticket: AdminTicket } = $props()
</script>

<div class="space-y-1 rounded-md border bg-white px-3 py-2">
	<h3 class="text-lg font-medium leading-tight">Historik</h3>

	{#each ticket.logs as log (log.id)}
		<div class="relative mb-8">
			<div class="bg-border absolute bottom-0 left-2 top-4 w-0.5"></div>
			<div class="flex items-start">
				<div class=" relative z-10 mr-4 h-4 w-4 rounded-full"></div>
				<div class="flex-1">
					<div class="mb-2 flex flex-col sm:flex-row sm:items-center sm:justify-between">
						<h3 class="font-medium">{log.status}</h3>
						<span class="text-muted-foreground text-sm">{log.inserted}</span>
					</div>
					<p class="mb-1 text-sm">
						<span class="font-medium">Initieret af:</span>
						{log.initiator}
					</p>

					<div class="mb-3">
						<p class="text-sm">{log.external_comment}</p>
					</div>

					{#if log.internal_comment}
						<div class="rounded-r-md border-l-4 border-amber-500 bg-amber-50 p-3">
							<div class="mb-1 flex items-center gap-1 text-sm font-medium text-amber-700">
								<EyeOff class="h-3.5 w-3.5" />
								<span>Intern kommentar (kun synlig for medarbejdere)</span>
							</div>
							<p class="text-sm text-amber-800">{log.internal_comment}</p>
						</div>
					{/if}
				</div>
			</div>
		</div>
	{/each}
</div>
