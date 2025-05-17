<script lang="ts">
	import UserButton from '$lib/components/common/user-button.svelte'

	let { children, data } = $props()
</script>

<main class="bg-muted/20 min-h-screen w-full">
	<header class="flex items-center justify-between border-b bg-white px-3 py-4 shadow-sm">
		<h1 class="font-medium leading-tight">RMA Portal</h1>
		{#if data.user}
			<UserButton user={data.user} />
		{/if}
	</header>
	<div class="mx-auto max-w-6xl px-3 py-4">
		{@render children()}
	</div>
</main>
