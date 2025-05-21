<script lang="ts">
	import { page } from '$app/state'
	import UserButton from '$lib/components/common/user-button.svelte'
	import Button from '$lib/components/ui/button/button.svelte'
	import { cn } from '$lib/utils.js'
	import { ScanBarcode } from '@lucide/svelte'

	let { children, data } = $props()
</script>

<main class="bg-muted/20 relative flex min-h-screen w-full flex-col">
	<header
		class="sticky top-4 mx-auto mt-4 flex w-full max-w-6xl items-center justify-between rounded-lg border-b bg-white px-3 py-2 shadow-sm"
	>
		<div class="flex items-center gap-8">
			<a href="/" class="flex items-center gap-2 font-medium">
				<div class="grid size-7 place-items-center rounded-lg bg-red-600">
					<ScanBarcode class="size-[18px] text-white" />
				</div>
				<span class="text-sm leading-none">Skancode RMA</span>
			</a>
			{#if data.user}
				<nav class="text-muted-foreground flex items-center gap-4 text-sm font-medium">
					<a
						class={cn(page.url.pathname == '/admin/tickets' && 'text-foreground')}
						href="/admin/tickets">Oversigt</a
					>
					<a class={cn(page.url.pathname == '/opret' && 'text-foreground')} href="/opret">Opret</a>
				</nav>
			{/if}
		</div>
		{#if data.user}
			<UserButton user={data.user} />
		{:else}
			<Button href="/log-ind" size="sm" variant="secondary" class="h-7">Log ind</Button>
		{/if}
	</header>
	<div class="mx-auto flex w-full max-w-6xl flex-1 flex-col px-3 py-4">
		{@render children()}
	</div>
</main>
