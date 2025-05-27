<script lang="ts">
	import { page } from '$app/state'
	import MetaMenu from '$lib/components/common/meta-menu.svelte'
	import UserButton from '$lib/components/common/user-button.svelte'
	import { cn } from '$lib/utils.js'
	import { ScanBarcode } from '@lucide/svelte'

	let { children, data } = $props()
</script>

{#if data.user}
	<MetaMenu />
{/if}
<main class="relative flex min-h-screen w-full flex-col bg-muted/20">
	{#if data.user}
		<header
			class="sticky top-4 mx-auto mt-4 flex w-full max-w-6xl items-center justify-between rounded-lg border-b bg-white px-3 py-2 shadow-sm"
		>
			<div class="flex items-center gap-8">
				<a href="/" class="flex items-center gap-2 font-medium">
					<div class="grid size-7 place-items-center rounded-lg bg-brand">
						<ScanBarcode class="size-[18px] text-white" />
					</div>
					<span class="text-sm leading-none">Skancode RMA</span>
				</a>
				<nav class="flex items-center gap-4 text-sm font-medium text-muted-foreground">
					<a
						class={cn(page.url.pathname == '/admin/sager' && 'text-foreground')}
						href="/admin/sager">Oversigt</a
					>
					<a class={cn(page.url.pathname == '/opret' && 'text-foreground')} href="/opret">Opret</a>
				</nav>
			</div>
			<UserButton user={data.user} />
		</header>
	{/if}
	<div class="mx-auto flex w-full max-w-6xl flex-1 flex-col px-3 py-4">
		{@render children()}
	</div>
</main>
