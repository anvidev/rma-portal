import { goto } from '$app/navigation'
import { SvelteURLSearchParams } from 'svelte/reactivity'

const activeParams = new Set<string>()
let debounceTimer = 0

function updateUrl(searchParams: URLSearchParams) {
	clearTimeout(debounceTimer)
	console.log('active', activeParams, searchParams)

	let empty = $derived(searchParams.values().every(v => v == ''))

	if (searchParams.toString() === '') {
		goto('', { keepFocus: true, noScroll: true })
		return
	}

	debounceTimer = setTimeout(() => {
		const query = searchParams.toString()
		goto(query ? `?${query}` : '', { keepFocus: true, noScroll: true })
	}, 2000)
}

function createQueryState<T>(
	searchParams: SvelteURLSearchParams,
	key: string,
	{
		parse,
		serialize,
		default: defaultValue,
	}: {
		parse: (value: string) => T
		serialize: (value: T) => string
		default?: T
	},
) {
	activeParams.add(key)

	const initial = searchParams.has(key) ? parse(searchParams.get(key)!) : defaultValue
	let value = $state(initial)
	let empty = $derived(searchParams.values().every(v => v == ''))

	$effect(() => {
		console.log(key, value, searchParams)
		if (value === undefined || value === null || value === '') {
			console.log('delete me', key)
			searchParams.delete(key)
		} else {
			console.log('set me', key)
			searchParams.set(key, serialize(value))
		}

		// Special handling when last parameter is cleared
		if (empty) {
			console.log('effect is empty')
			goto('', { keepFocus: true, noScroll: true })
		} else {
			updateUrl(searchParams)
		}
	})

	return {
		get value() {
			return value
		},
		set value(v) {
			value = v
		},
	}
}

export function stringQueryState(
	searchParams: SvelteURLSearchParams,
	key: string,
	options: { default?: string } = {},
) {
	return createQueryState(searchParams, key, {
		parse: v => v,
		serialize: v => v,
		default: options.default,
	})
}

export function numberQueryState(
	searchParams: SvelteURLSearchParams,
	key: string,
	options: { default?: number | null } = {},
) {
	return createQueryState<number | null>(searchParams, key, {
		parse: v => (v === '' ? null : Number(v)),
		serialize: v => (v === null ? '' : String(v)),
		default: options.default ?? null,
	})
}

export function stringArrayQueryState(
	searchParams: SvelteURLSearchParams,
	key: string,
	options: { default?: string[] } = {},
) {
	return createQueryState(searchParams, key, {
		parse: v => {
			if (v === '') return []
			return v.split(',').filter(Boolean)
		},
		serialize: v => {
			return v.length > 0 ? v.join(',') : ''
		},
		default: options.default ?? [],
	})
}

export function booleanQueryState(
	searchParams: SvelteURLSearchParams,
	key: string,
	options: { default?: boolean } = {},
) {
	return createQueryState(searchParams, key, {
		parse: v => v === 'true',
		serialize: v => (v ? 'true' : 'false'),
		default: options.default ?? false,
	})
}
