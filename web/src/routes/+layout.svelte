<script lang="ts">
	import '../app.css'
	let { children } = $props()
</script>

<main class="min-h-screen w-full px-3 py-2">
	<div class="mx-auto max-w-7xl">
		{@render children()}
	</div>
</main>
