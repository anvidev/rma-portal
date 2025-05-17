<script lang="ts">
	import UserButton from '$lib/components/common/user-button.svelte'
	import { ScanBarcode } from '@lucide/svelte'

	let { children, data } = $props()

	let logoAnchorPath = $derived(data.user ? '/admin/tickets' : '/opret')
</script>

<main class="bg-muted/20 relative min-h-screen w-full">
	<header
		class="sticky top-0 flex items-center justify-between border-b bg-white px-3 py-4 shadow-sm"
	>
		<a href={logoAnchorPath} class="flex items-center gap-2 font-medium">
			<div class="flex h-[30px] items-center gap-1 rounded-lg bg-red-500 px-2 text-base text-white">
				<ScanBarcode class="size-[18px]" />
				<span class="leading-none">Skancode</span>
			</div>
			<span class="text-base">RMA Service Portal</span>
		</a>
		{#if data.user}
			<UserButton user={data.user} />
		{/if}
	</header>
	<div class="mx-auto max-w-6xl px-3 py-4">
		{@render children()}
	</div>
</main>
