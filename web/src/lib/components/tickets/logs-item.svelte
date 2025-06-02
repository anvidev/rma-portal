<script lang="ts">
	import { type Log } from '$lib/types'
	import * as Tooltip from '$lib/components/ui/tooltip/index.js'
	import { cn, formatDate } from '$lib/utils'
	import { FileText, Shield } from '@lucide/svelte'
	import { Lightbox } from '$lib/lightbox/lightbox-state.svelte'
	let { log, internal = false }: { log: Log; internal?: boolean } = $props()

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

	const attachments = [
		{
			id: 1,
			type: 'image',
			filename: 'mike.jpg',
			url: 'https://anvi.dev/files/00007-C43E6/1748259425-mike.jpg',
		},
		{
			id: 2,
			type: 'document',
			filename: 'report.pdf',
			url: 'https://anvi.dev/files/00007-C43E6/bd378640-2bc6-4c9a-8fb9-dd769a06f13b.pdf',
		},
		{
			id: 3,
			type: 'image',
			filename: '3.jpg',
			url: 'https://anvi.dev/files/00008-245C0/1748261219-3.jpg',
		},
	]

	let lightbox = new Lightbox({
		images: attachments
			.filter(att => att.type == 'image')
			.map(att => ({ id: att.id, url: att.url })),
		loop: true,
	})
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
		<div class="bg-muted/40 mt-3 rounded-lg border p-3">
			<div class="text-muted-foreground mb-1 flex items-center gap-1 text-sm">
				<Shield class="size-3.5" />
				<span class="font-medium">Intern kommentar</span>
			</div>
			<p class="whitespace-pre-line text-sm">{log.internal_comment}</p>
		</div>
	{/if}
	{#if internal && attachments.length > 0}
		<div class="mt-3 flex items-center gap-2">
			{#each attachments as att (att.id)}
				<Tooltip.Provider delayDuration={250}>
					<Tooltip.Root>
						<Tooltip.Trigger>
							<div
								class="group flex size-12 cursor-pointer items-center justify-center overflow-hidden rounded-lg border"
							>
								{#if att.type == 'document'}
									<a
										href={att.url}
										target="_blank"
										class="bg-muted/40 group-hover:bg-muted/100 grid h-full w-full place-items-center transition-colors"
									>
										<FileText class="size-5 fill-slate-200 text-slate-500" />
									</a>
								{:else if att.type == 'image'}
									<div role="img" onclick={() => lightbox.open(att.id)}>
										<img
											alt={att.filename}
											src={att.url}
											class="transition-transform group-hover:scale-105"
										/>
									</div>
								{/if}
							</div>
						</Tooltip.Trigger>
						<Tooltip.Content sideOffset={8}>
							<div class="">{att.filename}</div>
						</Tooltip.Content>
					</Tooltip.Root>
				</Tooltip.Provider>
			{/each}
		</div>
	{/if}
</div>
