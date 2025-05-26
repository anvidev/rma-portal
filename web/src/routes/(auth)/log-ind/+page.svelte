<script>
	import Button from '$lib/components/ui/button/button.svelte'
	import Input from '$lib/components/ui/input/input.svelte'
	import Label from '$lib/components/ui/label/label.svelte'
</script>

<title>Log ind - Skancode RMA Service Portal</title>

<section class="w-72 space-y-4">
	<div>
		<h1 class="text-lg font-semibold leading-tight">Velkommen tilbage!</h1>
		<p class="text-sm text-muted-foreground">Log ind for at servicere RMA sager.</p>
	</div>

	<form method="POST" class="space-y-4">
		<div class="grid gap-2">
			<Label for="email">E-mail</Label>
			<Input autocomplete="work email" autofocus name="email" id="email" />
		</div>

		<div class="grid gap-2">
			<Label for="password">Adgangskode</Label>
			<Input autocomplete="current-password" name="password" id="password" type="password" />
		</div>

		<Button class="w-full" type="submit">Log ind</Button>
	</form>
</section>
