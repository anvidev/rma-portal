<script lang="ts">
	let { children } = $props()
</script>

<main class="bg-muted/20 min-h-screen w-full">
	<div class="mx-auto grid min-h-screen max-w-6xl place-items-center">
		{@render children()}
	</div>
</main>
