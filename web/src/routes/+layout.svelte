<script lang="ts">
	import { Toaster } from '$lib/components/ui/sonner'
	import '../app.css'
	let { children } = $props()
</script>

<Toaster />
<main class="min-h-screen w-full bg-muted/20">
	<header class="border-b bg-white px-3 py-4 shadow-sm">
		<h1 class="font-medium leading-tight">RMA Portal</h1>
	</header>
	<div class="mx-auto max-w-6xl px-3 py-4">
		{@render children()}
	</div>
</main>
