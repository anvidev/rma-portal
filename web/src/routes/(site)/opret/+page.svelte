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

<title>Opret RMA - Skancode RMA Service Portal</title>

<form
	method="POST"
	use:enhance
	class="mx-auto max-w-3xl space-y-4 p-6"
>
	<div>
		<h1 class="text-lg font-semibold leading-tight">Opret RMA</h1>
		<p class="text-sm text-muted-foreground">
			Udfyld formularen og vedlæg den genererede PDF-label, når du sender varen retur.
		</p>
	</div>
	{#if $message}
		<p>{$message}</p>
	{/if}
	<div>
		<h2 class="font-semibold leading-tight">Afsender Information</h2>
		<p class="text-sm text-muted-foreground">
			Disse oplysninger bruges til kommunikation vedrørende din RMA og som returadresse.
		</p>
	</div>

	<div class="flex items-start justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="senderCompany">Firma<span class="text-red-500">*</span></Label>
			<Input
				id="senderCompany"
				bind:value={$form.sender.company}
				aria-invalid={$errors?.sender?.company ? 'true' : undefined}
			/>
			{#if $errors?.sender?.company}<span class="text-sm text-destructive"
					>{$errors.sender?.company}</span
				>{/if}
		</div>
		<div class="grid flex-1 gap-2">
			<Label for="senderName">Kontaktperson<span class="text-red-500">*</span></Label>
			<Input
				id="senderName"
				bind:value={$form.sender.name}
				aria-invalid={$errors?.sender?.name ? 'true' : undefined}
			/>
			{#if $errors?.sender?.name}<span class="text-sm text-destructive">{$errors.sender?.name}</span
				>{/if}
		</div>
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
		<p class="text-sm text-muted-foreground">Disse oplysninger bruges til fakturering.</p>
	</div>
	<div class="flex items-center gap-2">
		<Checkbox
			id="billing-checkbox"
			bind:checked={isBillingSame}
			onCheckedChange={checked => handleIsBillingSame(checked)}
		/>
		<Label class="font-normal" for="billing-checkbox">Fakturering er den samme som afsender</Label>
	</div>

	<div class="flex items-center justify-stretch gap-4">
		<div class="grid flex-1 gap-2">
			<Label for="billingCompany">Firma<span class="text-red-500">*</span></Label>
			<Input
				id="billingCompany"
				disabled={isBillingSame}
				bind:value={$form.billing.company}
				aria-invalid={$errors?.billing?.company ? 'true' : undefined}
			/>
			{#if $errors?.billing?.company}<span class="text-sm text-destructive"
					>{$errors.billing?.company}</span
				>{/if}
		</div>
		<div class="grid flex-1 gap-2">
			<Label for="billingName">Kontaktperson<span class="text-red-500">*</span></Label>
			<Input
				id="billingName"
				disabled={isBillingSame}
				bind:value={$form.billing.name}
				aria-invalid={$errors?.billing?.name ? 'true' : undefined}
			/>
			{#if $errors?.billing?.name}<span class="text-sm text-destructive"
					>{$errors.billing?.name}</span
				>{/if}
		</div>
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

	<div class="flex items-start justify-stretch gap-4">
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
		<p class="text-sm text-muted-foreground">
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
				bind:value={$form.serial_number}
				aria-invalid={$errors?.serial_number ? 'true' : undefined}
			/>
			{#if $errors?.serial_number}<span class="text-sm text-destructive"
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
		{#if $errors?.issue}<span class="text-sm text-destructive">{$errors.issue}</span>{/if}
	</div>

	<div>
		<h2 class="text-sm font-medium">Præmisser</h2>
		<p class="text-sm text-muted-foreground">
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
	{#if $errors?.quote}<span class="text-sm text-destructive">{$errors.quote}</span>{/if}

	<div>
		<h2 class="text-sm font-medium">Garanti</h2>
		<p class="text-sm text-muted-foreground">
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
	{#if $errors?.warranty}<span class="text-sm text-destructive">{$errors.warranty}</span>{/if}

	<div>
		<h2 class="text-sm font-medium">Betingelser</h2>
		<p class="text-sm text-muted-foreground">
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
				class="cursor-pointer font-medium underline hover:text-primary"
				onclick={() => openModal('terms')}
			>
				vilkår og betingelser
			</button>
			og
			<button
				type="button"
				class="cursor-pointer font-medium underline hover:text-primary"
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

<Dialog.Root open={modalContent === 'terms'} onOpenChange={_open => (modalContent = null)}>
	<Dialog.DialogContent class="max-w-2xl">
		<Dialog.DialogHeader>
			<Dialog.DialogTitle>Vilkår og betingelser</Dialog.DialogTitle>
			<Dialog.DialogDescription>
				Her er vores vilkår og betingelser for RMA-processen.
			</Dialog.DialogDescription>
		</Dialog.DialogHeader>
		<div class="max-h-[60vh] overflow-y-auto">
			<div class="space-y-4">
				<div class="space-y-1">
					<h2 class="text-sm font-medium">1. Introduktion</h2>
					<p class="text-sm">
						Disse vilkår gælder for brugen af Skancodes RMA-serviceportal. Portalen gør det muligt
						for kunder at indsende RMA-sager vedrørende hardware- eller softwareprodukter.
					</p>
				</div>
				<div class="space-y-1">
					<h2 class="text-sm font-medium">2. Brug og ansvar</h2>
					<p class="text-sm">
						Portalen må kun bruges af personer, der er bemyndiget til at handle på vegne af den
						virksomhed, som opretter RMA-sagen. Brugeren er ansvarlig for, at de angivne oplysninger
						er korrekte.
					</p>
				</div>
				<div class="space-y-1">
					<h2 class="text-sm font-medium">3. Behandling af RMA-sager</h2>
					<p class="text-sm">
						Vi behandler RMA-sager på baggrund af de indsendte oplysninger. Hvis enheden ikke er
						dækket af garanti, kontakter vi kunden og tilbyder reparation mod betaling, hvis kunden
						har givet samtykke til dette.
					</p>
				</div>
				<div class="space-y-1">
					<h2 class="text-sm font-medium">4. Garantidækning og tilbudspris</h2>
					<p class="text-sm">
						Ved indsendelse af en RMA-sag kan kunden vælge, om der ønskes et tilbud, hvis
						reparationen overstiger en minimumspris. I tilfælde af afvist reparation kan et
						undersøgelsesgebyr pålægges.
					</p>
				</div>
				<div class="space-y-1">
					<h2 class="text-sm font-medium">5. Ansvarsbegrænsning</h2>
					<p class="text-sm">
						Skancode kan ikke holdes ansvarlig for indirekte tab som følge af fejl eller forsinket
						behandling af en RMA-sag.
					</p>
				</div>
				<div class="space-y-1">
					<h2 class="text-sm font-medium">6. Ændringer</h2>
					<p class="text-sm">
						Skancode forbeholder sig retten til at ændre vilkårene uden varsel. De gældende vilkår
						vil altid være tilgængelige på portalen.
					</p>
				</div>
			</div>
		</div>
	</Dialog.DialogContent>
</Dialog.Root>

<Dialog.Root open={modalContent === 'privacy'} onOpenChange={_open => (modalContent = null)}>
	<Dialog.DialogContent class="max-w-2xl">
		<Dialog.DialogHeader>
			<Dialog.DialogTitle>Privatlivspolitik</Dialog.DialogTitle>
			<Dialog.DialogDescription>
				Sådan håndterer vi dine personoplysninger.
			</Dialog.DialogDescription>
		</Dialog.DialogHeader>
		<div class="max-h-[60vh] overflow-y-auto">
			<div class="space-y-4">
				<div class="space-y-1">
					<h2 class="text-sm font-medium">1. Dataansvarlig</h2>
					<p class="text-sm">Den dataansvarlige for behandlingen af dine personoplysninger er:</p>
					<ul class="text-sm">
						<li>Skancode A/S</li>
						<li>Hejrevang 13, 3450 Allerød</li>
						<li>CVR: 36917954</li>
						<li>Telefon: 72220211</li>
						<li>Email: info@skancode.dk</li>
					</ul>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">2. Formål og behandlingsgrundlag</h2>
					<p>Vi indsamler og behandler personoplysninger med følgende formål:</p>
					<ul class="list-inside list-disc">
						<li>Behandling af RMA-sager</li>
						<li>Kommunikation vedrørende din RMA-sag</li>
						<li>Fakturering og administration</li>
						<li>Overholdelse af retlige forpligtelser, herunder bogføringskrav</li>
					</ul>
					<p>Retsgrundlaget for behandlingen er:</p>
					<ul class="list-inside list-disc">
						<li>GDPR art. 6(1)(b): Opfyldelse af kontrakt</li>
						<li>GDPR art. 6(1)(c): Retlig forpligtelse</li>
						<li>GDPR art. 6(1)(f): Legitim interesse</li>
					</ul>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">3. Hvilke oplysninger vi behandler</h2>
					<p>Vi behandler følgende kategorier af personoplysninger:</p>
					<ul class="list-inside list-disc">
						<li>Firmanavn eller kontaktpersons navn</li>
						<li>E-mailadresse</li>
						<li>Telefonnummer</li>
						<li>Adresse (vejnavn, by, postnummer, land)</li>
						<li>Om sagen vedrører software eller hardware</li>
						<li>Modelnavn og serienummer</li>
						<li>Fejlbeskrivelse</li>
						<li>Oplysning om garanti</li>
						<li>Ønske om tilbud, hvis reparation overstiger minimumspris</li>
					</ul>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">4. Opbevaring af data</h2>
					<p>
						Vi opbevarer personoplysninger så længe det er nødvendigt til de formål, hvortil de blev
						indsamlet, og i overensstemmelse med gældende lovgivning, herunder bogføringsloven.
					</p>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">5. Videregivelse og databehandlere</h2>
					<p>Vi videregiver ikke dine oplysninger til tredjeparter med henblik på markedsføring.</p>
					<p>Vi benytter følgende databehandler:</p>
					<ul class="list-inside list-disc">
						<li>Hetzner Online GmbH</li>
					</ul>
					<p>Databasen hostes internt og kontrolleres udelukkende af os.</p>
					<p>Vi overfører ikke personoplysninger uden for EU.</p>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">6. Dine rettigheder</h2>
					<p>Du har følgende rettigheder efter GDPR:</p>
					<ul class="list-inside list-disc">
						<li>Ret til indsigt i de oplysninger, vi behandler om dig</li>
						<li>Ret til berigtigelse (rettelse) af forkerte oplysninger</li>
						<li>Ret til sletning (“retten til at blive glemt”) i visse tilfælde</li>
						<li>Ret til begrænsning af behandling</li>
						<li>Ret til dataportabilitet</li>
						<li>Ret til at gøre indsigelse mod behandlingen</li>
					</ul>
					<p>
						Du kan gøre brug af dine rettigheder ved at kontakte os via ovenstående
						kontaktoplysninger.
					</p>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">7. Sikkerhed</h2>
					<p>
						Vi har truffet tekniske og organisatoriske foranstaltninger for at beskytte dine
						oplysninger mod uautoriseret adgang, ændring, tab eller ødelæggelse.
					</p>
				</div>
				<div class="space-y-1 text-sm">
					<h2 class="font-medium">8. Ændringer</h2>
					<p>
						Vi forbeholder os retten til at opdatere denne privatlivspolitik. Du vil blive
						informeret om væsentlige ændringer via portalen eller e-mail.
					</p>
					<p>Senest opdateret: 21/05-25.</p>
				</div>
			</div>
		</div>
	</Dialog.DialogContent>
</Dialog.Root>
