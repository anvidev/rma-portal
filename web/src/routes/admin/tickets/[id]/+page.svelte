<script lang="ts">
	import type { PageData } from './$types'

	export let data: PageData
</script>

<h1>{data.ticket.sender.name}</h1>
