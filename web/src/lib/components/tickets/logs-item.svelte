<script lang="ts">
	import { type Log } from '$lib/types'
	import * as Tooltip from '$lib/components/ui/tooltip/index.js'
	import { cn, formatDate, isDocument, isImage } from '$lib/utils'
	import { FileText, FileUp, Shield } from '@lucide/svelte'
	import { Lightbox } from '$lib/lightbox/lightbox-state.svelte'
	import { toast } from 'svelte-sonner'
	import { flip } from 'svelte/animate'

	let { log, internal = false }: { log: Log; internal?: boolean } = $props()

	let files = $state(log.files ?? [])

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

	let lightbox = $derived.by(() => {
		return new Lightbox({
			images: files
				.filter(f => isImage(f.mime_type))
				.map(att => ({ id: att.id, url: att.file_url })),
			loop: true,
		})
	})

	async function onchange(event: Event) {
		const target = event.target as HTMLInputElement | null

		if (!target || !target.files) return

		for (let file of target.files) {
			const presignResponse = await fetch('/api/upload', {
				method: 'PUT',
				body: JSON.stringify({
					key: `files/${log.ticket_id}/logs/${Date.now()}-${file.name}`,
					reference_id: log.id.toString(10),
					file_name: file.name,
					file_domain: 'logs',
					content_type: file.type,
					content_length: file.size,
				}),
			})
			if (!presignResponse.ok) {
				toast.error(`Kunne ikke hente en upload url for ${file.name}`)
				return
			}

			const { url } = await presignResponse.json()

			files.push({
				id: Math.random() * Math.PI,
				reference_id: log.id.toString(10),
				file_name: file.name,
				file_domain: 'logs',
				file_url: URL.createObjectURL(file),
				mime_type: file.type,
				inserted: new Date().toLocaleDateString('en-GB'),
			})

			const putResponse = await fetch(url, {
				method: 'PUT',
				headers: {
					'Content-Type': file.type,
					'Content-Length': file.size.toString(10),
				},
				body: file,
			})
			if (!putResponse.ok) {
				toast.error(`${file.name} blev ikke uploadet til R2`)
				return
			}

			toast.success('Fil blev uploadet')
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
		<div class="bg-muted/40 mt-3 rounded-lg border p-3">
			<div class="text-muted-foreground mb-1 flex items-center gap-1 text-sm">
				<Shield class="size-3.5" />
				<span class="font-medium">Intern kommentar</span>
			</div>
			<p class="whitespace-pre-line text-sm">{log.internal_comment}</p>
		</div>
	{/if}
	{#if internal}
		<div class="mt-3 flex items-center gap-2">
			{#if files && files.length > 0}
				{#each files as file (file.id)}
					<div animate:flip>
						<Tooltip.Provider delayDuration={250}>
							<Tooltip.Root>
								<Tooltip.Trigger>
									{#snippet child({ props })}
										<div
											class="group flex size-12 cursor-pointer items-center justify-center overflow-hidden rounded-lg border"
											{...props}
										>
											{#if isDocument(file.mime_type)}
												<a
													href={file.file_url}
													target="_blank"
													class="bg-muted/40 group-hover:bg-muted/100 grid h-full w-full place-items-center transition-colors"
												>
													<FileText class="size-5 fill-slate-200 text-slate-600" />
												</a>
											{:else if isImage(file.mime_type)}
												<button
													class="h-full w-full"
													type="button"
													onclick={() => lightbox.open(file.id)}
												>
													<img
														alt={file.file_name}
														src={file.file_url}
														class="h-full w-full object-cover transition-transform group-hover:scale-105"
													/>
												</button>
											{/if}
										</div>
									{/snippet}
								</Tooltip.Trigger>
								<Tooltip.Content sideOffset={8}>
									<div class="">{file.file_name}</div>
								</Tooltip.Content>
							</Tooltip.Root>
						</Tooltip.Provider>
					</div>
				{/each}
			{/if}
			<Tooltip.Provider delayDuration={250}>
				<Tooltip.Root>
					<Tooltip.Trigger>
						{#snippet child({ props })}
							<div {...props} class="group flex items-center justify-center">
								<label
									for={`dropzone-log-${log.id}`}
									class="bg-muted/40 group-hover:bg-muted/100 flex size-12 cursor-pointer flex-col items-center justify-center rounded-lg border border-dashed transition-colors"
								>
									<div class="flex flex-col items-center justify-center">
										<FileUp class="size-5 fill-slate-200 text-slate-600" />
									</div>
									<input
										{onchange}
										id={`dropzone-log-${log.id}`}
										type="file"
										class="hidden"
										accept="image/*, application/pdf, application/ms"
										multiple
									/>
								</label>
							</div>
						{/snippet}
					</Tooltip.Trigger>
					<Tooltip.Content sideOffset={8}>
						<div>Upload filer</div>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>
		</div>
	{/if}
</div>
