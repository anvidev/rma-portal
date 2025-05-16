<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js'
	import { Label } from '$lib/components/ui/label/index.js'
	import { Textarea } from '$lib/components/ui/textarea/index.js'
	import { Checkbox } from '$lib/components/ui/checkbox/index.js'
	import { Input } from '$lib/components/ui/input/index.js'
	import { Button } from '$lib/components/ui/button/index.js'
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js'
	import SuperDebug, { superForm } from 'sveltekit-superforms'

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
</script>

<form method="POST" use:enhance class="mx-auto max-w-3xl space-y-4 p-6">
	<div>
		<h1 class="text-lg font-semibold leading-tight">Opret RMA</h1>
		<p class="text-sm text-muted-foreground">
			Udfyld formen og vedlæg den downloadede PDF-fil efter oprettelse.
		</p>
	</div>
	{#if $message}
		<p>{$message}</p>
	{/if}
	<div>
		<h2 class="font-semibold leading-tight">Afsender Information</h2>
		<p class="text-sm text-muted-foreground">
			Denne information bruges til evt. til kommunikation ang. RMA og som returaddresse.
		</p>
	</div>
	<div class="grid gap-2">
		<Label for="senderName">Firma<span class="text-red-500">*</span></Label>
		<Input
			id="senderName"
			bind:value={$form.sender.name}
			aria-invalid={$errors?.sender?.name ? 'true' : undefined}
		/>
		{#if $errors?.sender?.name}<span class="text-sm text-destructive">{$errors.sender?.name}</span
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
			{#if $errors?.sender?.email}<span class="text-sm text-destructive"
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
			{#if $errors?.sender?.phone}<span class="text-sm text-destructive"
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
		{#if $errors?.sender?.street}<span class="text-sm text-destructive"
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
			{#if $errors?.sender?.city}<span class="text-sm text-destructive">{$errors.sender?.city}</span
				>{/if}
		</div>

		<div class="grid w-1/4 gap-2">
			<Label for="senderZip">Postnr.<span class="text-red-500">*</span></Label>
			<Input
				id="senderZip"
				bind:value={$form.sender.zip}
				aria-invalid={$errors?.sender?.zip ? 'true' : undefined}
			/>
			{#if $errors?.sender?.zip}<span class="text-sm text-destructive">{$errors.sender?.zip}</span
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
			{#if $errors?.sender?.country}<span class="text-sm text-destructive"
					>{$errors.sender?.country}</span
				>{/if}
		</div>
	</div>

	<div>
		<h2 class="font-semibold leading-tight">Fakturerings Information</h2>
		<p class="text-sm text-muted-foreground">Denne information bruges til fakturering.</p>
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
		{#if $errors?.billing?.name}<span class="text-sm text-destructive">{$errors.billing?.name}</span
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
			{#if $errors?.billing?.email}<span class="text-sm text-destructive"
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
			{#if $errors?.billing?.phone}<span class="text-sm text-destructive"
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
		{#if $errors?.billing?.street}<span class="text-sm text-destructive"
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
			{#if $errors?.billing?.city}<span class="text-sm text-destructive"
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
			{#if $errors?.billing?.zip}<span class="text-sm text-destructive">{$errors.billing?.zip}</span
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
			{#if $errors?.billing?.country}<span class="text-sm text-destructive"
					>{$errors.billing?.country}</span
				>{/if}
		</div>
	</div>

	<div>
		<h2 class="font-semibold leading-tight">RMA Information</h2>
		<p class="text-sm text-muted-foreground">Beskriv problemet så detaljeret som muligt.</p>
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
		{#if $errors?.categories}<span class="text-sm text-destructive"
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
			{#if $errors?.model}<span class="text-sm text-destructive">{$errors.model}</span>{/if}
		</div>
		<div class="grid w-2/4 gap-2">
			<Label for="serial">Serienummer</Label>
			<Input
				id="serial"
				bind:value={$form.serial}
				aria-invalid={$errors?.serial ? 'true' : undefined}
			/>
			{#if $errors?.serial}<span class="text-sm text-destructive">{$errors.serial}</span>{/if}
		</div>
	</div>

	<div class="grid gap-2">
		<Label for="issue">Fejlbeskrivelse<span class="text-red-500">*</span></Label>
		<Textarea
			id="issue"
			bind:value={$form.issue}
			aria-invalid={$errors?.issue ? 'true' : undefined}
		/>
		{#if $errors?.issue}<span class="text-sm text-destructive">{$errors.issue}</span>{/if}
	</div>

	<h2 class="text-sm font-medium">Præmisser</h2>
	<p class="text-sm text-muted-foreground">
		Skancode A/S dækker reparation og returfragt i garantisager. For øvrige henvendelser pålægges en
		minimumspris på 450 DKK pr. indsendt enhed, ekskl. fragt.
	</p>
	<RadioGroup.Root bind:value={$form.wantsQuote}>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="Yes" id="option-one" />
			<Label for="option-one"
				>Jeg ønsker at modtage et tilbud hvis reparation overstiger minimumsprisen</Label
			>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="No" id="option-two" />
			<Label for="option-two">
				Jeg ønsker IKKE at modtage et tilbud hvis reparation overstiger minimumsprisen
			</Label>
		</div>
	</RadioGroup.Root>
	{#if $errors?.wantsQuote}<span class="text-sm text-destructive">{$errors.wantsQuote}</span>{/if}

	<div class="flex w-full items-center justify-end gap-2">
		<Button variant="secondary" type="button" onclick={() => reset()}>Nulstil</Button>
		<Button type="submit">Opret</Button>
	</div>
</form>

<SuperDebug data={$form} />
