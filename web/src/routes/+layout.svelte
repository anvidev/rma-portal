<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<main class="min-h-screen w-full bg-slate-950 text-slate-50">
	<div class="mx-auto max-w-7xl">
		{@render children()}
	</div>
</main>
