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

<title>Tak for din henvendelse - Skancode RMA Service Portal</title>

<div class="flex flex-1 flex-col items-center justify-center gap-4">
	<div class="max-w-lg space-y-4 text-pretty text-center">
		<h1 class="text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl">
			Tak for din henvendelse!
		</h1>
		<p class="mx-auto text-muted-foreground">
			Vi har modtaget din RMA-anmodning og behandler den hurtigst muligt. Vi takker for din
			tålmodighed og ser frem til at hjælpe dig.
		</p>
	</div>

	<div class="flex items-center gap-4">
		<Button onclick={downloadLabel}>Download RMA Label</Button>
		<Button size="lg" variant="outline" href={`/sager/${page.params.id}`}>Følg din RMA</Button>
	</div>
</div>
