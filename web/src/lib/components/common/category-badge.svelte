<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { cn } from '$lib/utils'
	let { category, onclick }: { category: string; onclick?: () => void } = $props()

	function getCategoryBadgeBackgroundColor(status: string): string {
		switch (status) {
			default:
				return 'bg-secondary text-secondary-foreground'
		}
	}

	function getCategoryDotBackgroundColor(status: string): string {
		switch (status) {
			default:
				return 'bg-secondary-foreground'
		}
	}
</script>

<Badge
	onclick={onclick ? onclick : () => {}}
	class={cn(
		'flex items-center gap-2 capitalize',
		getCategoryBadgeBackgroundColor(category),
		onclick && 'cursor-pointer',
	)}
>
	<span class={cn('block size-1.5 rounded-full', getCategoryDotBackgroundColor(category))}></span>
	{category}</Badge
>
