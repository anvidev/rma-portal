<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js'
	import { Label } from '$lib/components/ui/label/index.js'
	import { Textarea } from '$lib/components/ui/textarea/index.js'
	import { Checkbox } from '$lib/components/ui/checkbox/index.js'
	import { Input } from '$lib/components/ui/input/index.js'
	import { Button } from '$lib/components/ui/button/index.js'
	import SuperDebug from 'sveltekit-superforms'
	import { superForm } from 'sveltekit-superforms'

	let { data } = $props()
	let isBillingSame = $state(false)

	const { form, enhance } = superForm(data.form, {
		resetForm: false,
		clearOnSubmit: 'errors-and-message',
		dataType: 'json'
	})

	function handleIsBillingSame(checked: boolean) {
		isBillingSame = checked
		if (checked) {
			$form.billing = $form.sender
		}
	}

	let availableCountries: { [key: string]: string } = {
		DK: 'Danmark (DK)',
		SE: 'Sverige (SE)',
		DE: 'Tyskland (DE)',
		PL: 'Polen (PL)'
	}

	function addItem(category: string) {
		$form.categories = [...$form.categories, category]
	}

	function removeItem(category: string) {
		$form.categories = $form.categories.filter((cat) => cat !== category)
	}
</script>

<form method="POST" use:enhance class="mx-auto max-w-2xl space-y-4 rounded-xl border p-6">
	<h1 class="text-lg font-medium">Opret RMA</h1>
	<h2 class="text-sm font-medium">Afsender Information</h2>
	<div class="grid gap-2">
		<Label for="senderName">Firma<span class="text-red-500">*</span></Label>
		<Input id="senderName" bind:value={$form.sender.name} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="senderEmail">E-mail<span class="text-red-500">*</span></Label>
			<Input id="senderEmail" type="email" bind:value={$form.sender.email} />
		</div>

		<div class="grid flex-1 gap-2">
			<Label for="senderPhone">Tlf.<span class="text-red-500">*</span></Label>
			<Input id="senderPhone" type="tel" bind:value={$form.sender.phone} />
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="senderStreet">Vejnavn<span class="text-red-500">*</span></Label>
		<Input id="senderStreet" bind:value={$form.sender.street} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="senderCity">By<span class="text-red-500">*</span></Label>
			<Input id="senderCity" bind:value={$form.sender.city} />
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="senderZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input id="senderZip" bind:value={$form.sender.zip} />
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="senderCountry">Land<span class="text-red-500">*</span></Label>
			<Select.Root type="single" bind:value={$form.sender.country}>
				<Select.Trigger
					>{$form.sender.country
						? availableCountries[$form.sender.country]
						: 'Vælg land'}</Select.Trigger
				>
				<Select.Content>
					{#each Object.entries(availableCountries) as [val, label]}
						<Select.Item value={val}>{label}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	<h2 class="text-sm font-medium">Fakturering Information</h2>
	<div class="flex items-center gap-2">
		<Checkbox
			id="billing-checkbox"
			bind:checked={isBillingSame}
			onCheckedChange={(checked) => handleIsBillingSame(checked)}
		/>
		<Label class="font-normal" for="billing-checkbox">Fakturering er den samme som afsender</Label>
	</div>

	<div class="grid gap-2">
		<Label for="billingName">Firma<span class="text-red-500">*</span></Label>
		<Input id="billingName" disabled={isBillingSame} bind:value={$form.billing.name} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="billingEmail">E-mail<span class="text-red-500">*</span></Label>
			<Input
				id="billingEmail"
				disabled={isBillingSame}
				type="email"
				bind:value={$form.billing.email}
			/>
		</div>

		<div class="grid flex-1 gap-2">
			<Label for="billingPhone">Tlf.<span class="text-red-500">*</span></Label>
			<Input
				id="billingPhone"
				disabled={isBillingSame}
				type="tel"
				bind:value={$form.billing.phone}
			/>
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="billingStreet">Vejnavn<span class="text-red-500">*</span></Label>
		<Input id="billingStreet" disabled={isBillingSame} bind:value={$form.billing.street} />
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="billingCity">By<span class="text-red-500">*</span></Label>
			<Input id="billingCity" disabled={isBillingSame} bind:value={$form.billing.city} />
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="billingZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input id="billingZip" disabled={isBillingSame} bind:value={$form.billing.zip} />
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="billingCountry">Land<span class="text-red-500">*</span></Label>
			<Select.Root type="single" disabled={isBillingSame} bind:value={$form.billing.country}>
				<Select.Trigger
					>{$form.billing.country
						? availableCountries[$form.billing.country]
						: 'Vælg land'}</Select.Trigger
				>
				<Select.Content>
					{#each Object.entries(availableCountries) as [val, label]}
						<Select.Item value={val}>{label}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	<h2 class="text-sm font-medium">RMA Information</h2>
	<div class="grid gap-2">
		<Label for="categorties">Kategorier<span class="text-red-500">*</span></Label>
		<div class="flex items-center gap-2">
			{#each ['software', 'hardware'] as category}
				{@const checked = $form.categories.includes(category)}
				<Label class="flex items-center gap-2 font-normal capitalize">
					<Checkbox
						value={category}
						onCheckedChange={(checked) => {
							if (checked) {
								addItem(category)
							} else {
								removeItem(category)
							}
						}}
					/>
					{category}
				</Label>
			{/each}
		</div>
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="model">Model</Label>
			<Input id="model" bind:value={$form.model} />
		</div>
		<div class="grid w-2/4 gap-2">
			<Label for="serial">Serienummer</Label>
			<Input id="serial" bind:value={$form.serial} />
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="issue">Problem<span class="text-red-500">*</span></Label>
		<Textarea id="issue" bind:value={$form.issue} />
	</div>

	<Button type="submit">Opret</Button>
</form>

<SuperDebug data={$form} />
