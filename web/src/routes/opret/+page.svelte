<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js'
	import {Label} from '$lib/components/ui/label/index.js'
	import {Checkbox} from '$lib/components/ui/checkbox/index.js'
	import {Input} from '$lib/components/ui/input/index.js'
	import {Button} from '$lib/components/ui/button/index.js'
	import SuperDebug from 'sveltekit-superforms'
	import {superForm} from 'sveltekit-superforms'

	let {data} = $props()
	let isBillingSame = $state(false)
	const {form} = superForm(data.form, {resetForm: false})

	$effect(() => {
		if (isBillingSame) {
			$form.billingName = $form.senderName
		}
	})

	let availableCountries: {[key: string]: string} = {
		DK: 'Danmark (DK)',
		SE: 'Sverige (SE)',
		DE: 'Tyskland (DE)',
		PL: 'Polen (PL)'
	}
</script>

<form method="post" class="mx-auto max-w-2xl space-y-4 rounded-xl border p-6">
	<h1 class="text-lg font-medium">Opret RMA</h1>
	<h2 class="text-sm font-medium">Afsender Information</h2>
	<div class="grid gap-2">
		<Label for="senderName">Firma<span class="text-red-500">*</span></Label>
		<Input id="senderName" disabled={isBillingSame} bind:value={$form.senderName} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="senderEmail">E-mail<span class="text-red-500">*</span></Label>
			<Input id="senderEmail" type="email" disabled={isBillingSame} bind:value={$form.senderEmail} />
		</div>

		<div class="grid flex-1 gap-2">
			<Label for="senderPhone">Tlf.<span class="text-red-500">*</span></Label>
			<Input id="senderPhone" type="tel" disabled={isBillingSame} bind:value={$form.senderPhone} />
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="senderStreet">Vejnavn<span class="text-red-500">*</span></Label>
		<Input id="senderStreet" disabled={isBillingSame} bind:value={$form.senderStreet} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-3/4 gap-2">
			<Label for="senderCity">By<span class="text-red-500">*</span></Label>
			<Input id="senderCity" disabled={isBillingSame} bind:value={$form.senderCity} />
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="senderZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input id="senderZip" disabled={isBillingSame} bind:value={$form.senderZip} />
		</div>
	</div>

	<div>
		<Label for="senderCountry">Land<span class="text-red-500">*</span></Label>
		<Select.Root type="single" bind:value={$form.senderCountry}>
			<Select.Trigger>{$form.senderCountry
				? availableCountries[$form.senderCountry]
				: 'Vælg land'}</Select.Trigger>
			<Select.Content>
				{#each Object.entries(availableCountries) as [val, label]}
				<Select.Item value={val}>{label}</Select.Item>
				{/each}
			</Select.Content>
		</Select.Root>
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
