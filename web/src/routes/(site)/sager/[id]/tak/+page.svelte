<script lang="ts">
	import { page } from '$app/state'
	import Button from '$lib/components/ui/button/button.svelte'
	import { toast } from 'svelte-sonner'

	async function downloadLabel() {
		const id = page.params.id
		const params = new URLSearchParams({ rma: id })
		const labelResponse = await fetch(`/api/label?${params}`)
		if (!labelResponse.ok) {
			const { message } = await labelResponse.json()
			toast.error('Der gik noget galt', { description: message })
			return
		}
		const blob = await labelResponse.blob()
		const url = URL.createObjectURL(blob)
		window.open(url, '_blank')
	}
</script>

<svelte:head>
	<meta name="robots" content="noindex, nofollow" />
	<title>Tak for din henvendelse - Skancode RMA Service Portal</title>
</svelte:head>

<div class="flex flex-1 flex-col items-center justify-center gap-4">
	<div class="flex flex-col items-center space-y-4 text-pretty text-center">
		<h1 class="max-w-lg text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl">
			Tak for din henvendelse!
		</h1>
		<p class="max-w-lg text-muted-foreground">
			Vi har modtaget din anmodning og behandler den hurtigst muligt. Tak for din tålmodighed – vi
			ser frem til at hjælpe dig.
		</p>
		<p class="max-w-lg text-muted-foreground">
			Venligst download din RMA-label og enten vedlæg den i pakken eller fastgør den udenpå, når du
			sender varen retur. <span class="font-bold"
				>Bemærk, at RMA-labelen ikke fungerer som en forsendelseslabel.</span
			>
		</p>
	</div>

	<div class="flex items-center gap-4">
		<Button onclick={downloadLabel}>Download RMA Label</Button>
		<Button size="lg" variant="outline" href={`/sager/${page.params.id}`}>Følg din RMA</Button>
	</div>
</div>
