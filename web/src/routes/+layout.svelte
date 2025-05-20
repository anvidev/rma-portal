<script lang="ts">
	import { Toaster } from '$lib/components/ui/sonner'
	import '../app.css'
	let { children } = $props()
</script>

<Toaster duration={4000} />
{@render children()}
