<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte'
	import Input from '$lib/components/ui/input/input.svelte'
	import { MoveRight } from '@lucide/svelte'

	let search = $state()
</script>

<div class="mx-auto flex flex-1 flex-col items-center justify-center gap-16">
	<div class="flex flex-col items-center gap-8">
		<div class="space-y-4 text-center">
			<h1 class="max-w-[36rem] text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl">
				Professionel, hurtig og autoriseret service.
			</h1>
			<p class="text-muted-foreground max-w-[32rem]">
				Vi reparerer din stregkodescanner, håndterminal eller etiketprinter med originale
				reservedele – enten på vores værksted eller on-site.
			</p>
		</div>

		<div class="flex items-center gap-4">
			<Button href="/opret">Opret ny RMA</Button>
			<Button
				variant="outline"
				onclick={() => {
					const input = document.getElementById('input')?.focus()
				}}>Følg din RMA</Button
			>
		</div>
	</div>

	<div class="flex w-full items-center gap-2">
		<Input id="input" bind:value={search} placeholder="F.eks. 0003-LEET" />
		<Button class="group">
			<span>Søg</span>
			<MoveRight class="size-3.5 transition-transform group-hover:translate-x-0.5" />
		</Button>
	</div>
</div>
