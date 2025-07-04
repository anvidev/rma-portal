<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import { Package } from '@lucide/svelte'
	import Button from '../ui/button/button.svelte'
	let open = $state(false)
</script>

<Dialog.Root {open} onOpenChange={val => (open = val)}>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<Button {...props} size="sm" variant="secondary" class="flex items-center gap-2">
				<Package class="size-3.5" />
				<span class="hidden md:block">Send med GLS</span>
			</Button>
		{/snippet}
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Send RMA med GLS</Dialog.Title>
			<Dialog.Description>
				Køb og betal for en GLS pakkelabel – helt uden brugerregistrering
			</Dialog.Description>
		</Dialog.Header>
		<div class="flex flex-col gap-4">
			<p>
				Her kan du nemt købe en pakkelabel via GLS’ officielle portal:
				<strong><a href="https://send.gls.dk" target="_blank">send.gls.dk</a></strong>. Du skal blot
				indtaste afsender- og modtageroplysninger, betale med kort eller MobilePay, og så modtager
				du en QR-kode via mail.
			</p>
			<p>
				QR-koden kan scannes i en hvilken som helst GLS PakkeShop, hvor de printer pakkelabelen for
				dig. Du behøver hverken en konto eller en printer – bare vis QR‑koden i shoppen.
			</p>
			<p class="text-destructive">
				Vigtigt: Du må <u>ikke</u> vælge levering til en pakkeshop. Pakken skal sendes direkte til vores
				adresse:
			</p>
			<address class="not-italic leading-tight">
				Skancode A/S<br />Hejrevang 34, 1.tv<br />3450 Allerød<br />Danmark
			</address>
			<Button href="https://send.gls.dk" target="_blank">Gå til GLS’ sendeformular</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
