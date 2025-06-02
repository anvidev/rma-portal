<script lang="ts">
	import { page } from '$app/state'
	import Button from '$lib/components/ui/button/button.svelte'
	import { TriangleAlert } from '@lucide/svelte'
</script>

<section class="grid min-h-screen w-full place-items-center">
	<div class="flex flex-col items-center justify-center gap-8">
		<TriangleAlert class="size-16 animate-bounce text-orange-500" />
		<div class="flex flex-col items-center justify-center gap-2">
			<h1 class="text-xl font-semibold">Hov! Der gik noget galt</h1>
			{#if page.error?.requestId}
				<p class="text-xs text-muted-foreground">Request ID: {page.error?.requestId}</p>
			{/if}
			<p class="text-muted-foreground">{page.error?.message}</p>
		</div>
		<div class="flex items-center gap-4">
			<Button variant="secondary" onclick={() => history.back()}>Gå tilbage</Button>
			<Button onclick={() => location.reload()}>Prøv igen</Button>
		</div>
	</div>
</section>
