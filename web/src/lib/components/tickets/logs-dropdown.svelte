<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import { Ellipsis } from '@lucide/svelte'
	import { Button } from '../ui/button'
	import { Label } from '../ui/label'
	import { Textarea } from '../ui/textarea'
	import { superForm } from 'sveltekit-superforms'
	import { toast } from 'svelte-sonner'
	import type { Log } from '$lib/types'

	let {
		log,
	}: {
		log: Log
	} = $props()

	const { form, enhance, errors } = superForm(log, {
		resetForm: true,
		clearOnSubmit: 'errors-and-message',
		dataType: 'json',
		onUpdated({ form }) {
			if (form.message) {
				editOpen = false
				toast(form.message)
			}
		},
	})

	let editOpen = $state(false)
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		{#snippet child({ props })}
			<Button {...props} variant="ghost" class="flex size-6 items-center justify-center px-0">
				<Ellipsis />
			</Button>
		{/snippet}
	</DropdownMenu.Trigger>
	<DropdownMenu.Content>
		<DropdownMenu.Group>
			<DropdownMenu.Item onclick={() => (editOpen = true)}>Rediger</DropdownMenu.Item>
		</DropdownMenu.Group>
		<DropdownMenu.Separator />
		<DropdownMenu.Group>
			<DropdownMenu.Item disabled>Slet</DropdownMenu.Item>
		</DropdownMenu.Group>
	</DropdownMenu.Content>
</DropdownMenu.Root>

<Dialog.Root open={editOpen} onOpenChange={val => (editOpen = val)}>
	<Dialog.Content>
		<form action="?/update" method="POST" use:enhance>
			<Dialog.Header>
				<Dialog.Title>Opdater status</Dialog.Title>
				<Dialog.Description>
					Rediger kommentare for RMA status opdatering. Intern kommentar er valgfri.
				</Dialog.Description>
			</Dialog.Header>
			<div class="my-4 space-y-4">
				<div class="grid flex-1 gap-2">
					<Label for="external">Ekstern kommentar<span class="text-red-500">*</span></Label>
					<Textarea id="external" bind:value={$form.external_comment} />
					{#if $errors?.external_comment}<span class="text-destructive text-sm"
							>{$errors.external_comment}</span
						>{/if}
				</div>

				<div class="grid flex-1 gap-2">
					<Label for="internal">Intern kommentar</Label>
					<Textarea id="internal" bind:value={$form.internal_comment} />
					<p class="text-muted-foreground text-sm">
						Denne kommentar er kun synlig for Skancode A/S
					</p>
					{#if $errors?.internal_comment}<span class="text-destructive text-sm"
							>{$errors.internal_comment}</span
						>{/if}
				</div>
			</div>
			<Dialog.Footer>
				<Button type="submit">Opdater</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
