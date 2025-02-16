<script lang="ts">
	import { goto } from '$app/navigation'
	import { SvelteURLSearchParams } from 'svelte/reactivity'

	let searchParams = new SvelteURLSearchParams()
	let query: string = $state('')
	let status: string = $state('')
	let categories = $state([])
	let debounceTimer: ReturnType<typeof setTimeout>

	function handleFilterChange(key: string, value: any, waitMS: number = 300) {
		clearTimeout(debounceTimer)

		if (Array.isArray(value) && value.length == 0) {
			searchParams.delete(key)
		} else if (typeof value == 'string' && value == '') {
			searchParams.delete(key)
		} else {
			searchParams.set(key, value)
		}

		debounceTimer = setTimeout(() => {
			goto(`?${searchParams.toString()}`, { keepFocus: true })
		}, waitMS)
	}
</script>

<input type="text" bind:value={query} onkeyup={() => handleFilterChange('query', query)} />
<select bind:value={status} onchange={() => handleFilterChange('status', status, 0)}>
	<option value="">Alle</option>
	<option value="åben">Åben</option>
	<option value="lukket">Lukket</option>
</select>
<div>
	{#each ['software', 'hardware'] as category}
		<label>
			<input
				type="checkbox"
				name="flavours"
				value={category}
				bind:group={categories}
				onchange={() => handleFilterChange('categories', categories)}
			/>
			{category}
		</label>
	{/each}
</div>
