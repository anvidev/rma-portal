<script lang="ts">
	import type { WithElementRef, WithoutChildren } from 'bits-ui'
	import type { HTMLTextareaAttributes } from 'svelte/elements'
	import { cn } from '$lib/utils.js'

	let {
		ref = $bindable(null),
		value = $bindable(),
		class: className,
		maxlength,
		...restProps
	}: WithoutChildren<WithElementRef<HTMLTextareaAttributes>> = $props()
</script>

<div class="relative">
	<textarea
		bind:this={ref}
		class={cn(
			'flex min-h-[80px] w-full rounded-lg border border-input bg-background px-3 py-2 text-base ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
			className,
		)}
		bind:value
		{...restProps}
	></textarea>
	{#if maxlength}
		<span class={cn("absolute bottom-2 right-3 text-xs", String(value).length > maxlength && "text-destructive")}>{value ? String(value).length : 0} / {maxlength}</span>
	{/if}
</div>
