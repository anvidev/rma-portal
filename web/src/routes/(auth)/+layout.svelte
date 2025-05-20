<script lang="ts">
	let { children } = $props()
</script>

<main class="min-h-screen w-full bg-muted/20">
	<div class="mx-auto grid min-h-screen max-w-6xl place-items-center">
		{@render children()}
	</div>
</main>
