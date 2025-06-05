<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import { Textarea } from '$lib/components/ui/textarea/index.js'
	import * as Select from '$lib/components/ui/select/index.js'
	import type { TicketWithLogs } from '$lib/types'
	import { Plus } from '@lucide/svelte'
	import Button from '../ui/button/button.svelte'
	import Label from '../ui/label/label.svelte'
	import { filesProxy, superForm, type SuperValidated } from 'sveltekit-superforms'
	import { toast } from 'svelte-sonner'
	import LogsItem from './logs-item.svelte'
	import Input from '../ui/input/input.svelte'

	let open = $state(false)

	let {
		ticket,
		statuses,
		superform,
	}: {
		ticket: TicketWithLogs
		statuses: string[]
		superform: SuperValidated<{
			status: string
			external_comment: string
			internal_comment: string
			files?: File[]
		}>
	} = $props()

	const { form, enhance, errors } = superForm(superform, {
		resetForm: true,
		clearOnSubmit: 'errors-and-message',
		dataType: 'json',
		onUpdated({ form }) {
			if (form.message) {
				open = false
				toast(form.message)
			}
		},
	})

	const files = filesProxy(form, 'files')
</script>

<div class="flex items-center justify-between">
	<div>
		<h3 class="font-semibold leading-tight">Seneste opdateringer</h3>
		<p class="text-muted-foreground text-sm">Følg status og opdateringer RMA sag</p>
	</div>

	<Dialog.Root {open} onOpenChange={val => (open = val)}>
		<Dialog.Trigger>
			{#snippet child({ props })}
				<Button {...props} size="sm" variant="secondary" class="flex items-center gap-2">
					<Plus class="size-3.5" />
					Tilføj opdatering
				</Button>
			{/snippet}
		</Dialog.Trigger>
		<Dialog.Content>
			<form method="POST" use:enhance enctype="multipart/form-data">
				<Dialog.Header>
					<Dialog.Title>Opdater RMA</Dialog.Title>
					<Dialog.Description>
						Opdater RMA sagen med ny status og en ekstern kommentar. Intern kommentar er valgfri
					</Dialog.Description>
				</Dialog.Header>
				<div class="my-4 space-y-4">
					<div class="grid flex-1 gap-2">
						<Label for="status">Status<span class="text-red-500">*</span></Label>
						<Select.Root type="single" bind:value={$form.status}>
							<Select.Trigger class="capitalize"
								>{$form.status ? $form.status : 'Vælg status'}</Select.Trigger
							>
							<Select.Content>
								{#each statuses as status, index (index)}
									<Select.Item class="capitalize" value={status}>{status}</Select.Item>
								{/each}
							</Select.Content>
						</Select.Root>
						{#if $errors?.status}<span class="text-destructive text-sm">{$errors.status}</span>{/if}
					</div>

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

					<!-- <div class="grid gap-2"> -->
					<!-- 	<Label for="files">Filer</Label> -->
					<!-- 	<Input -->
					<!-- 		placeholder="Der er ikke valgt nogen filer" -->
					<!-- 		type="file" -->
					<!-- 		multiple -->
					<!-- 		name="files" -->
					<!-- 		bind:files={$files} -->
					<!-- 	/> -->
					<!-- 	<p class="text-sm text-muted-foreground">Filer er kun synlig for Skancode A/S</p> -->
					<!-- 	{#if $errors?.files}<span class="text-sm text-destructive" -->
					<!-- 			>{$errors.files._errors[0]}</span -->
					<!-- 		>{/if} -->
					<!-- </div> -->
				</div>
				<Dialog.Footer>
					<Button type="submit">Opret</Button>
				</Dialog.Footer>
			</form>
		</Dialog.Content>
	</Dialog.Root>
</div>
<div class="mt-8 flex flex-col gap-4">
	{#each ticket.logs as log (log.id)}
		<LogsItem {log} internal={true} />
	{/each}
</div>
