<script lang="ts">
	import { fade, fly } from 'svelte/transition'
	import { activeLightbox } from './lightbox-store'
	import { onMount } from 'svelte'
	import { outside } from '$lib/outside-event'

	function handleKeyDown(e: KeyboardEvent) {
		if (!$activeLightbox) return
		if ($activeLightbox.current()) {
			if (e.key === 'Escape') {
				$activeLightbox.close()
			} else if (e.key === 'ArrowRight' || e.key === 'l') {
				$activeLightbox.next()
			} else if (e.key === 'ArrowLeft' || e.key === 'h') {
				$activeLightbox.previous()
			} else if (e.ctrlKey && e.shiftKey && e.key == 'S') {
				$activeLightbox.download()
			}
		}
	}

	onMount(() => {
		document.addEventListener('keydown', handleKeyDown)

		return () => {
			document.removeEventListener('keydown', handleKeyDown)
		}
	})
</script>

{#if $activeLightbox && $activeLightbox.current()}
	<section
		out:fade
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
	>
		<div
			in:fly={{ y: 50 }}
			out:fly={{ y: -50 }}
			class="mx-2 w-fit rounded-lg bg-white p-3 shadow-xl sm:mx-0"
			use:outside
			onoutside={() => $activeLightbox.close()}
		>
			<img
				class="h-auto w-full rounded-md object-contain max-w-[90vw] max-h-[90dvh]"
				alt={$activeLightbox.current()?.id.toString()}
				src={$activeLightbox.current()?.url}
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
