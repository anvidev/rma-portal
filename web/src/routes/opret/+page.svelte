<script lang="ts">
	import { Label } from '$lib/components/ui/label/index.js'
	import { Checkbox } from '$lib/components/ui/checkbox/index.js'
	import { Input } from '$lib/components/ui/input/index.js'
	import { Button } from '$lib/components/ui/button/index.js'
	import SuperDebug from 'sveltekit-superforms'
	import { superForm } from 'sveltekit-superforms'

	let { data } = $props()
	let isBillingSame = $state(false)
	const { form } = superForm(data.form)

	$effect(() => {
		if (isBillingSame) {
			$form.billingName = $form.senderName
		}
	})
</script>

<h1>Opret RMA</h1>

<form method="post" class="space-y-4 rounded-xl border p-6">
	<div class="grid gap-2">
		<Label for="senderName">Afsender firma<span class="text-red-500">*</span></Label>
		<Input id="senderName" disabled={isBillingSame} bind:value={$form.senderName} />
	</div>

	<div class="flex items-center gap-2">
		<Checkbox id="billing-checkbox" bind:checked={isBillingSame} />
		<Label for="billing-checkbox">Fakturering er den samme som afsender</Label>
	</div>

	<div class="grid gap-2">
		<Label for="billingName">Fakturerings firma<span class="text-red-500">*</span></Label>
		<Input id="billingName" disabled={isBillingSame} bind:value={$form.billingName} />
	</div>

	<div class="grid gap-2">
		<Label for="categorties">Kategorier<span class="text-red-500">*</span></Label>
		<div>
			{#each ['software', 'hardware'] as category}
				<label>
					<input type="checkbox" name="categories" value={category} bind:group={$form.categories} />
					{category}
				</label>
			{/each}
		</div>
	</div>

	<Button type="submit">Opret</Button>
</form>

<SuperDebug data={$form} />
