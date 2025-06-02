<script lang="ts">
	import { type Log } from '$lib/types'
	import { cn, formatDate } from '$lib/utils'
	import { Shield } from '@lucide/svelte'
	let { log }: { log: Log } = $props()

	function getStatusBorderColor(status: string): string {
		switch (status) {
			case 'registreret':
				return 'border-l-gray-600'
			case 'modtaget':
				return 'border-l-green-600'

			case 'lukket':
			case 'afvist':
				return 'border-l-red-600'

			case 'intern reparation':
			case 'ekstern reparation':
				return 'border-l-amber-600'

			case 'afventer reservedele':
			case 'tilbud accepteret':
			case 'tilbud sendt':
				return 'border-l-sky-600'

			default:
				return 'border-l-gray-600'
		}
	}
</script>

<div
	class={cn(
		'rounded-lg border border-l-[3px] bg-white p-3 shadow-sm',
		getStatusBorderColor(log.status),
	)}
>
	<div class="space-y-2">
		<div class="flex items-center justify-between">
			<p class="text-sm font-medium capitalize leading-tight">{log.status}</p>
			<p class="text-sm leading-tight">{formatDate(log.inserted)}</p>
		</div>
		<small class="text-muted-foreground">{log.initiator}</small>
		<p class="whitespace-pre-line text-sm">{log.external_comment}</p>
	</div>
	{#if log.internal_comment}
		<div class="mt-3 rounded-lg border bg-muted/40 p-3">
			<div class="mb-1 flex items-center gap-1 text-sm text-muted-foreground">
				<Shield class="size-3.5" />
				<span class="font-medium">Intern kommentar</span>
			</div>
			<p class="whitespace-pre-line text-sm">{log.internal_comment}</p>
		</div>
	{/if}
</div>
