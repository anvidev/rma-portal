import { SvelteURLSearchParams } from 'svelte/reactivity'
// inspo https://www.reddit.com/r/sveltejs/comments/1d43d8p/svelte_5_runes_with_localstorage_thanks_to_joy_of/

export class URLState {
	params = $state(new SvelteURLSearchParams())
}
