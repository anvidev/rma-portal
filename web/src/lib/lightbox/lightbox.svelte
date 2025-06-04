<script lang="ts">
	import { fade, fly } from 'svelte/transition'
	import { activeLightbox } from './lightbox-store'
	import { cn } from '$lib/utils'

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

	// TODO: implement goto function for lightbox
</script>

<svelte:document on:keydown={handleKeyDown} />

{#if $activeLightbox && $activeLightbox.current()}
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
				class="h-auto w-full max-w-md rounded-md object-contain"
				alt={$activeLightbox.current()?.id.toString()}
				src={$activeLightbox.current()?.url}
			/>

			{#if false}
				<div class="mt-2 flex flex-wrap items-start gap-2">
					{#each $activeLightbox?.images() as img (img.id)}
						<button type="button" onclick={() => console.log('')}>
							<img
								class={cn(
									'size-10 rounded-md',
									$activeLightbox?.current()?.id == img.id &&
										'border border-primary opacity-80 transition-all',
								)}
								alt={img.id.toString()}
								src={img.url}
							/>
						</button>
					{/each}
				</div>
			{/if}

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
