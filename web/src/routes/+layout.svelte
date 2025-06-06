<script lang="ts">
	import { Toaster } from '$lib/components/ui/sonner'
	import Lightbox from '$lib/lightbox/lightbox.svelte'
	import '../app.css'
	let { children } = $props()
</script>

<Lightbox />
<Toaster duration={4000} />
{@render children()}
