<script lang="ts">
	import '../app.css'
	let { children } = $props()
</script>

<main class="bg-muted/20 min-h-screen w-full">
	<header class="border-b px-3 py-4">
		<h1 class="font-medium leading-tight">RMA Portal</h1>
	</header>
	<div class="mx-auto max-w-7xl px-3 py-4">
		{@render children()}
	</div>
</main>
