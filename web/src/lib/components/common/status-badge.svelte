<script lang="ts">
	import { Badge } from '$lib/components/ui/badge/index.js'
	import { cn } from '$lib/utils'
	let { status }: { status: string } = $props()

	function getStatusBadgeBackgroundColor(status: string): string {
		switch (status) {
			case 'registreret':
				return 'bg-gray-100 text-gray-600'
			case 'modtaget':
				return 'bg-green-100 text-green-600'

			case 'lukket':
			case 'afvist':
				return 'bg-red-100 text-red-600'

			case 'intern reparation':
			case 'ekstern reparation':
				return 'bg-amber-100 text-amber-600'

			case 'afventer reservedele':
			case 'tilbud accepteret':
			case 'tilbud sendt':
				return 'bg-sky-100 text-sky-600'

			default:
				return 'bg-gray-100 text-gray-600'
		}
	}

	function getStatusDotBackgroundColor(status: string): string {
		switch (status) {
			case 'registreret':
				return 'bg-gray-600'
			case 'modtaget':
				return 'bg-green-600'

			case 'lukket':
			case 'afvist':
				return 'bg-red-600'

			case 'intern reparation':
			case 'ekstern reparation':
				return 'bg-amber-600'

			case 'afventer reservedele':
			case 'tilbud accepteret':
			case 'tilbud sendt':
				return 'bg-sky-600'

			default:
				return 'bg-gray-600'
		}
	}
</script>

<Badge class={cn('flex items-center gap-2 capitalize', getStatusBadgeBackgroundColor(status))}>
	<span class={cn('block size-1.5 rounded-full', getStatusDotBackgroundColor(status))}></span>
	{status}</Badge
>
