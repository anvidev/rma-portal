<script lang="ts">
	import { Label } from '$lib/components/ui/label/index.js'
	import { Input } from '$lib/components/ui/input/index.js'
	import { Button } from '$lib/components/ui/button/index.js'
	import SuperDebug from 'sveltekit-superforms'
	import { superForm } from 'sveltekit-superforms'

	let { data } = $props()

	const { form } = superForm(data.form)

	function copySenderToBilling() {
		$form.billingName = $form.senderName
	}
</script>

<h1>Opret RMA</h1>

<form method="post" class="space-y-4 rounded-xl border p-6">
	<div class="grid gap-2">
		<Label for="senderName">Afsender navn</Label>
		<Input id="senderName" bind:value={$form.senderName} />
	</div>

	<Button type="button" variant="outline" class="mb-4" onclick={copySenderToBilling}
		>Samme som afsender</Button
	>

	<div class="grid gap-2">
		<Label for="billingName">Fakturerings navn</Label>
		<Input id="billingName" bind:value={$form.billingName} />
	</div>

	<div class="grid gap-2">
		<Label for="categorties">Kaegorier</Label>
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
