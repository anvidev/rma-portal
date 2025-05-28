<script lang="ts">
	import type { Lightbox } from '$lib/lightbox.svelte'
	import { fade, fly } from 'svelte/transition'

	let { lightbox = $bindable() }: { lightbox: Lightbox } = $props()

	function handleKeyDown(e: KeyboardEvent) {
		console.log('keys', e)
		if (lightbox.current()) {
			if (e.key === 'Escape') {
				lightbox.close()
			} else if (e.key === 'ArrowRight' || e.key === 'l') {
				lightbox.next()
			} else if (e.key === 'ArrowLeft' || e.key === 'h') {
				lightbox.previous()
			} else if (e.ctrlKey && e.shiftKey && e.key == 'S') {
				console.log('save')
				lightbox.download()
			}
		}
	}
</script>

<svelte:document on:keydown={handleKeyDown} />

{#if lightbox.current()}
	<section
		out:fade
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
	>
		<div
			in:fly={{ y: 50 }}
			out:fly={{ y: -50 }}
			class="mx-2 w-fit rounded-lg bg-white p-4 shadow-xl sm:mx-0 sm:max-w-3xl"
		>
			<img
				class="h-autobw-full max-w-md rounded-md object-contain"
				alt={lightbox.current().id.toString()}
				src={lightbox.current().url}
			/>

			<p class="mt-3 text-center text-sm text-gray-500">
				Tryk
				<kbd class="rounded border bg-gray-100 px-1 py-0.5 font-mono text-sm text-gray-700">←</kbd>
				/
				<kbd class="rounded border bg-gray-100 px-1 py-0.5 font-mono text-sm text-gray-700">→</kbd>
				eller
				<kbd class="rounded border bg-gray-100 px-1 py-0.5 font-mono text-sm text-gray-700">h</kbd>
				/
				<kbd class="rounded border bg-gray-100 px-1 py-0.5 font-mono text-sm text-gray-700">l</kbd>
				for at navigere,
				<kbd class="rounded border bg-gray-100 px-1 py-0.5 font-mono text-sm text-gray-700">Esc</kbd
				> for at lukke
			</p>
		</div>
	</section>
{/if}
