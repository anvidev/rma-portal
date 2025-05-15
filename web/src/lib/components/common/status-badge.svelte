<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { cn } from '$lib/utils'
	let { status }: { status: string } = $props()

	function getStatusBadgeBackgroundColor(status: string): string {
		switch (status) {
			case 'ekstern reperation':
				return 'bg-amber-100 text-amber-600'
			case 'intern reperation':
				return 'bg-emerald-100 text-emerald-600'
			case 'lukket':
				return 'bg-stone-100 text-stone-600'
			default:
				return 'bg-sky-100 text-sky-700'
		}
	}

	function getStatusDotBackgroundColor(status: string): string {
		switch (status) {
			case 'ekstern reperation':
				return 'bg-amber-600'
			case 'intern reperation':
				return 'bg-emerald-600'
			case 'lukket':
				return 'bg-stone-600'
			default:
				return 'bg-sky-700'
		}
	}
</script>

<Badge class={cn('flex items-center gap-2 capitalize', getStatusBadgeBackgroundColor(status))}>
	<span class={cn('block size-1.5 rounded-full', getStatusDotBackgroundColor(status))}></span>
	{status}</Badge
>
