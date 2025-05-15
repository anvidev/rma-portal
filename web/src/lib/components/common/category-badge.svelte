<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { cn } from '$lib/utils'
	let { category }: { category: string } = $props()

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

<Badge class={cn('flex items-center gap-2 capitalize', getCategoryBadgeBackgroundColor(category))}>
	<span class={cn('block size-1.5 rounded-full', getCategoryDotBackgroundColor(category))}></span>
	{category}</Badge
>
