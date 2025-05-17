<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js'
	import { Label } from '$lib/components/ui/label/index.js'
	import { Textarea } from '$lib/components/ui/textarea/index.js'
	import { Checkbox } from '$lib/components/ui/checkbox/index.js'
	import { Input } from '$lib/components/ui/input/index.js'
	import { Button } from '$lib/components/ui/button/index.js'
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import { superForm } from 'sveltekit-superforms'

	let { data } = $props()
	let isBillingSame = $state(false)

	const { form, enhance, errors, reset, message } = superForm(data.form, {
		resetForm: true,
		clearOnSubmit: 'errors-and-message',
		dataType: 'json',
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
		NO: 'Norge (NO)',
		DE: 'Tyskland (DE)',
		PL: 'Polen (PL)',
	}

	function addItem(category: string) {
		$form.categories = [...$form.categories, category]
	}

	function removeItem(category: string) {
		$form.categories = $form.categories.filter(cat => cat !== category)
	}

	let acceptTerms = $state(false)

	let modalContent = $state<'terms' | 'privacy' | null>(null)
	function openModal(content: 'terms' | 'privacy') {
		modalContent = content
	}
</script>

<form method="POST" use:enhance class="mx-auto max-w-3xl space-y-4 p-6">
	<div>
		<h1 class="text-lg font-semibold leading-tight">Opret RMA</h1>
		<p class="text-muted-foreground text-sm">
			Udfyld formularen og vedlæg den genererede PDF-fil, når du sender varen retur.
		</p>
	</div>
	{#if $message}
		<p>{$message}</p>
	{/if}
	<div>
		<h2 class="font-semibold leading-tight">Afsender Information</h2>
		<p class="text-muted-foreground text-sm">
			Disse oplysninger bruges til kommunikation vedrørende din RMA og som returadresse.
		</p>
	</div>
	<div class="grid gap-2">
		<Label for="senderName">Firma<span class="text-red-500">*</span></Label>
		<Input
			id="senderName"
			bind:value={$form.sender.name}
			aria-invalid={$errors?.sender?.name ? 'true' : undefined}
		/>
		{#if $errors?.sender?.name}<span class="text-destructive text-sm">{$errors.sender?.name}</span
			>{/if}
	</div>

	<div class="flex items-start justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="senderEmail">E-mail<span class="text-red-500">*</span></Label>
			<Input
				id="senderEmail"
				type="email"
				bind:value={$form.sender.email}
				aria-invalid={$errors?.sender?.email ? 'true' : undefined}
			/>
			{#if $errors?.sender?.email}<span class="text-destructive text-sm"
					>{$errors.sender?.email}</span
				>{/if}
		</div>

		<div class="grid flex-1 gap-2">
			<Label for="senderPhone">Tlf.<span class="text-red-500">*</span></Label>
			<Input
				id="senderPhone"
				type="tel"
				bind:value={$form.sender.phone}
				aria-invalid={$errors?.sender?.phone ? 'true' : undefined}
			/>
			{#if $errors?.sender?.phone}<span class="text-destructive text-sm"
					>{$errors.sender?.phone}</span
				>{/if}
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="senderStreet">Vejnavn<span class="text-red-500">*</span></Label>
		<Input
			id="senderStreet"
			bind:value={$form.sender.street}
			aria-invalid={$errors?.sender?.street ? 'true' : undefined}
		/>
		{#if $errors?.sender?.street}<span class="text-destructive text-sm"
				>{$errors.sender?.street}</span
			>{/if}
	</div>

	<div class="flex items-start justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="senderCity">By<span class="text-red-500">*</span></Label>
			<Input
				id="senderCity"
				bind:value={$form.sender.city}
				aria-invalid={$errors?.sender?.city ? 'true' : undefined}
			/>
			{#if $errors?.sender?.city}<span class="text-destructive text-sm">{$errors.sender?.city}</span
				>{/if}
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="senderZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input
				id="senderZip"
				bind:value={$form.sender.zip}
				aria-invalid={$errors?.sender?.zip ? 'true' : undefined}
			/>
			{#if $errors?.sender?.zip}<span class="text-destructive text-sm">{$errors.sender?.zip}</span
				>{/if}
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
			{#if $errors?.sender?.country}<span class="text-destructive text-sm"
					>{$errors.sender?.country}</span
				>{/if}
		</div>
	</div>

	<div>
		<h2 class="font-semibold leading-tight">Fakturerings Information</h2>
		<p class="text-muted-foreground text-sm">Disse oplysninger bruges til fakturering.</p>
	</div>
	<div class="flex items-center gap-2">
		<Checkbox
			id="billing-checkbox"
			bind:checked={isBillingSame}
			onCheckedChange={checked => handleIsBillingSame(checked)}
		/>
		<Label class="font-normal" for="billing-checkbox">Fakturering er den samme som afsender</Label>
	</div>

	<div class="grid gap-2">
		<Label for="billingName">Firma<span class="text-red-500">*</span></Label>
		<Input
			id="billingName"
			disabled={isBillingSame}
			bind:value={$form.billing.name}
			aria-invalid={$errors?.billing?.name ? 'true' : undefined}
		/>
		{#if $errors?.billing?.name}<span class="text-destructive text-sm">{$errors.billing?.name}</span
			>{/if}
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="billingEmail">E-mail<span class="text-red-500">*</span></Label>
			<Input
				id="billingEmail"
				disabled={isBillingSame}
				type="email"
				bind:value={$form.billing.email}
				aria-invalid={$errors?.billing?.email ? 'true' : undefined}
			/>
			{#if $errors?.billing?.email}<span class="text-destructive text-sm"
					>{$errors.billing?.email}</span
				>{/if}
		</div>

		<div class="grid flex-1 gap-2">
			<Label for="billingPhone">Tlf.<span class="text-red-500">*</span></Label>
			<Input
				id="billingPhone"
				disabled={isBillingSame}
				type="tel"
				bind:value={$form.billing.phone}
				aria-invalid={$errors?.billing?.phone ? 'true' : undefined}
			/>
			{#if $errors?.billing?.phone}<span class="text-destructive text-sm"
					>{$errors.billing?.phone}</span
				>{/if}
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="billingStreet">Vejnavn<span class="text-red-500">*</span></Label>
		<Input
			id="billingStreet"
			disabled={isBillingSame}
			bind:value={$form.billing.street}
			aria-invalid={$errors?.billing?.street ? 'true' : undefined}
		/>
		{#if $errors?.billing?.street}<span class="text-destructive text-sm"
				>{$errors.billing?.street}</span
			>{/if}
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="billingCity">By<span class="text-red-500">*</span></Label>
			<Input
				id="billingCity"
				disabled={isBillingSame}
				bind:value={$form.billing.city}
				aria-invalid={$errors?.billing?.city ? 'true' : undefined}
			/>
			{#if $errors?.billing?.city}<span class="text-destructive text-sm"
					>{$errors.billing?.city}</span
				>{/if}
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="billingZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input
				id="billingZip"
				disabled={isBillingSame}
				bind:value={$form.billing.zip}
				aria-invalid={$errors?.billing?.zip ? 'true' : undefined}
			/>
			{#if $errors?.billing?.zip}<span class="text-destructive text-sm">{$errors.billing?.zip}</span
				>{/if}
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
			{#if $errors?.billing?.country}<span class="text-destructive text-sm"
					>{$errors.billing?.country}</span
				>{/if}
		</div>
	</div>

	<div>
		<h2 class="font-semibold leading-tight">RMA Information</h2>
		<p class="text-muted-foreground text-sm">
			Beskriv problemet så præcist og detaljeret som muligt.
		</p>
	</div>
	<div class="grid gap-2">
		<Label for="categorties">Kategorier<span class="text-red-500">*</span></Label>
		<div class="flex items-center gap-2">
			{#each ['software', 'hardware'] as category}
				{@const checked = $form.categories.includes(category)}
				<Label class="flex items-center gap-2 font-normal capitalize">
					<Checkbox
						value={category}
						{checked}
						onCheckedChange={checked => {
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
		{#if $errors?.categories}<span class="text-destructive text-sm"
				>{$errors.categories._errors?.join('. ')}</span
			>{/if}
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid w-2/4 gap-2">
			<Label for="model">Model</Label>
			<Input
				id="model"
				bind:value={$form.model}
				aria-invalid={$errors?.model ? 'true' : undefined}
			/>
			{#if $errors?.model}<span class="text-destructive text-sm">{$errors.model}</span>{/if}
		</div>
		<div class="grid w-2/4 gap-2">
			<Label for="serial">Serienummer</Label>
			<Input
				id="serial"
				bind:value={$form.serial_number}
				aria-invalid={$errors?.serial_number ? 'true' : undefined}
			/>
			{#if $errors?.serial_number}<span class="text-destructive text-sm"
					>{$errors.serial_number}</span
				>{/if}
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="issue">Fejlbeskrivelse<span class="text-red-500">*</span></Label>
		<Textarea
			id="issue"
			bind:value={$form.issue}
			aria-invalid={$errors?.issue ? 'true' : undefined}
		/>
		{#if $errors?.issue}<span class="text-destructive text-sm">{$errors.issue}</span>{/if}
	</div>

	<div>
		<h2 class="text-sm font-medium">Præmisser</h2>
		<p class="text-muted-foreground text-sm">
			Skancode A/S dækker reparation og returfragt ved garantisager. Ved øvrige henvendelser
			pålægges en minimumspris på 450 DKK pr. enhed, ekskl. fragt.
		</p>
	</div>
	<RadioGroup.Root bind:value={$form.quote}>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="yes" id="quote-yes" />
			<Label for="quote-yes"
				>Jeg ønsker at modtage et tilbud hvis reparation overstiger minimumsprisen</Label
			>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="no" id="quote-no" />
			<Label for="quote-no">
				Jeg ønsker <span class="font-bold">IKKE</span> at modtage et tilbud hvis reparation overstiger
				minimumsprisen
			</Label>
		</div>
	</RadioGroup.Root>
	{#if $errors?.quote}<span class="text-destructive text-sm">{$errors.quote}</span>{/if}

	<div>
		<h2 class="text-sm font-medium">Garanti</h2>
		<p class="text-muted-foreground text-sm">
			Angiv om enheden er dækket af garanti. Dette hjælper os med korrekt håndtering af sagen.
		</p>
	</div>
	<RadioGroup.Root bind:value={$form.warranty}>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="yes" id="warranty-yes" />
			<Label for="warranty-yes">Enheden er dækket af garanti</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="no" id="warranty-no" />
			<Label for="warranty-no"
				>Enheden er <span class="font-bold">IKKE</span> dækket af garanti</Label
			>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="unknown" id="warranty-unknown" />
			<Label for="warranty-unknown"
				>Det er <span class="font-bold">UVIST</span>, om enheden er dækket af garanti</Label
			>
		</div>
	</RadioGroup.Root>
	{#if $errors?.warranty}<span class="text-destructive text-sm">{$errors.warranty}</span>{/if}

	<div>
		<h2 class="text-sm font-medium">Betingelser</h2>
		<p class="text-muted-foreground text-sm">
			For at indsende en RMA skal du acceptere vores vilkår samt privatlivspolitik.
		</p>
	</div>
	<div class="flex items-center space-x-2">
		<Checkbox id="terms" bind:checked={acceptTerms} aria-labelledby="terms-label" />
		<Label
			id="terms-label"
			for="terms"
			class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
		>
			Jeg accepterer
			<button
				type="button"
				class="hover:text-primary cursor-pointer font-medium underline"
				onclick={() => openModal('terms')}
			>
				vilkår og betingelser
			</button>
			og
			<button
				type="button"
				class="hover:text-primary cursor-pointer font-medium underline"
				onclick={() => openModal('privacy')}
			>
				privatlivspolitik
			</button>.
		</Label>
	</div>

	<div class="flex w-full items-center justify-end gap-2">
		<Button variant="secondary" type="button" onclick={() => reset()}>Nulstil</Button>
		<Button type="submit" disabled={!acceptTerms}>Opret</Button>
	</div>
</form>

<Dialog.Root open={modalContent === 'terms'} onOpenChange={open => (modalContent = null)}>
	<Dialog.DialogContent>
		<Dialog.DialogHeader>
			<Dialog.DialogTitle>Vilkår og betingelser</Dialog.DialogTitle>
			<Dialog.DialogDescription>
				Her er vores vilkår og betingelser for RMA-processen.
			</Dialog.DialogDescription>
		</Dialog.DialogHeader>
		<div class="max-h-[60vh] overflow-y-auto">
			<p>Vilkår og betingelser her...</p>
		</div>
	</Dialog.DialogContent>
</Dialog.Root>

<Dialog.Root open={modalContent === 'privacy'} onOpenChange={open => (modalContent = null)}>
	<Dialog.DialogContent>
		<Dialog.DialogHeader>
			<Dialog.DialogTitle>Privatlivspolitik</Dialog.DialogTitle>
			<Dialog.DialogDescription>
				Sådan håndterer vi dine personoplysninger.
			</Dialog.DialogDescription>
		</Dialog.DialogHeader>
		<div class="max-h-[60vh] overflow-y-auto">
			<p>Privatlivspolitik her...</p>
		</div>
	</Dialog.DialogContent>
</Dialog.Root>
