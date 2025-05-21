<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte'
	import Input from '$lib/components/ui/input/input.svelte'
	import { MoveRight } from '@lucide/svelte'
	import { toast } from 'svelte-sonner'
	import { superForm } from 'sveltekit-superforms'

	let { data } = $props()

	let { form, enhance, errors } = superForm(data.form, {
		resetForm: true,
		clearOnSubmit: 'errors-and-message',
		dataType: 'json',
		onUpdated({ form }) {
			if (!form.valid) {
				const description = form.errors.query?.at(0) ?? 'Der skete en ukendt fejl'
				toast.error('Der skete en fejl', { description })
			}
		},
	})
</script>

<div class="mx-auto flex flex-1 flex-col items-center justify-center gap-16">
	<div class="flex flex-col items-center gap-8">
		<div class="space-y-4 text-center">
			<h1 class="max-w-[36rem] text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl">
				Professionel, hurtig og autoriseret service.
			</h1>
			<p class="mx-auto max-w-[33rem] text-muted-foreground">
				Vi reparerer din stregkodescanner, håndterminal eller etiketprinter med originale
				reservedele – enten på vores værksted eller hos leverandør.
			</p>
		</div>

		<div class="flex items-center gap-4">
			<Button size="lg" href="/opret">Opret ny RMA</Button>
			<Button
				size="lg"
				variant="outline"
				onclick={() => {
					document.getElementById('input')?.focus()
				}}>Følg din RMA</Button
			>
		</div>
	</div>

	<form method="POST" use:enhance class="flex w-full items-center gap-2">
		<Input
			id="input"
			name="query"
			type="search"
			class="h-12 text-base placeholder:text-base"
			bind:value={$form.query}
			placeholder="F.eks. 00001-R2D20"
			aria-invalid={$errors.query ? true : undefined}
		/>
		<Button size="lg" class="group" type="submit">
			<span>Søg</span>
			<MoveRight class="size-3.5 transition-transform group-hover:translate-x-1" />
		</Button>
	</form>
</div>
