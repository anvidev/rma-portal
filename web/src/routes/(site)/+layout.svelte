<script lang="ts">
	import UserButton from '$lib/components/common/user-button.svelte'
	import { ScanBarcode } from '@lucide/svelte'

	let { children, data } = $props()

	let logoAnchorPath = $derived(data.user ? '/admin/tickets' : '/opret')
</script>

<main class="bg-muted/20 relative flex min-h-screen w-full flex-col">
	<header
		class="sticky top-0 flex items-center justify-between border-b bg-white px-3 py-4 shadow-sm"
	>
		<div class="flex items-center gap-8">
			<a href={logoAnchorPath} class="flex items-center gap-2 font-medium">
				<div
					class="bg-primary flex h-[30px] items-center gap-1 rounded-lg px-2 text-base text-white"
				>
					<ScanBarcode class="size-[18px]" />
					<span class="leading-none">Skancode</span>
				</div>
				<span class="text-base">RMA Service Portal</span>
			</a>
			{#if data.user}
				<nav class="text-muted-foreground flex items-center gap-4 text-sm">
					<a class="hover:underline" href="/admin/tickets">Sags oversigt</a>
					<a class="hover:underline" href="/opret">Opret RMA</a>
				</nav>
			{/if}
		</div>
		{#if data.user}
			<UserButton user={data.user} />
		{/if}
	</header>
	<div class="mx-auto flex w-full max-w-6xl flex-1 flex-col px-3 py-4">
		{@render children()}
	</div>
</main>
