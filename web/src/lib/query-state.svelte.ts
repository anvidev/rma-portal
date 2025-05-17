import { goto } from '$app/navigation'
import { SvelteURLSearchParams } from 'svelte/reactivity'

let debounceTimer: number | undefined = 0

function updateUrl(searchParams: URLSearchParams) {
	clearTimeout(debounceTimer)

	debounceTimer = setTimeout(() => {
		const query = searchParams.toString()
		const url = query !== '' ? `?${query}` : '?'
		goto(url, {
			keepFocus: true,
			noScroll: true,
		})
	}, 300)
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
	const initial = searchParams.has(key) ? parse(searchParams.get(key)!) : defaultValue
	let value = $state(initial)

	$effect(() => {
		const serialized = serialize(value)
		if (
			serialized === undefined ||
			serialized === null ||
			serialized === '' ||
			(Array.isArray(value) && Array.from(value).length === 0)
		) {
			searchParams.delete(key)
		} else {
			searchParams.set(key, serialized)
		}

		updateUrl(searchParams)
	})

	return {
		get value() {
			return value
		},
		set value(v) {
			value = v
		},
		...(!isNaN(value)
			? {
					increment() {
						value = (value ?? 1) + 1
					},
					decrement() {
						value = (value ?? 1) - 1
					},
				}
			: {}),
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
