import { goto } from '$app/navigation'
import { SvelteURLSearchParams } from 'svelte/reactivity'
// inspo https://www.reddit.com/r/sveltejs/comments/1d43d8p/svelte_5_runes_with_localstorage_thanks_to_joy_of/

export class URLParams {
	params = $state<Record<string, string | string[]>>({})
	debounceTimer = 0

	constructor(searchParams: URLSearchParams) {
		this.params = Object.fromEntries(searchParams.entries())

		$effect(() => {
			clearTimeout(this.debounceTimer)

			console.log('url effect', this.params)

			const tempParams: Record<string, string | string[]> = {}

			Object.entries(this.params).forEach(([key, val]) => {
				if ((typeof val === 'string' && val !== '') || (Array.isArray(val) && val.length > 0)) {
					tempParams[key] = val
				}
			})

			const searchParams = new SvelteURLSearchParams()
			Object.entries(tempParams).forEach(([key, val]) => {
				if (Array.isArray(val)) {
					searchParams.set(key, val.join(','))
				} else {
					searchParams.set(key, val)
				}
			})

			this.debounceTimer = setTimeout(() => {
				const query = searchParams.toString()
				goto(query ? `?${query}` : '', {
					keepFocus: true,
				})
			}, 300)
		})
	}
}
