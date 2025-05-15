import { goto } from '$app/navigation'
import { SvelteURLSearchParams } from 'svelte/reactivity'

type ParamType = 'string' | 'string[]' | 'number' | 'number[]' | 'boolean'

interface URLParamConfig {
	type: ParamType
	default?: any
}

export class URLParams<T extends Record<string, URLParamConfig>> {
	params = $state<Record<string, any>>({})
	debounceTimer = 0
	private config: T

	constructor(searchParams: URLSearchParams, config: T) {
		this.config = config

		for (const [key, { type, default: defaultValue }] of Object.entries(config)) {
			const rawValue = searchParams.get(key)

			if (rawValue !== null) {
				this.params[key] = this.parseValue(rawValue, type)
			} else if (defaultValue !== undefined) {
				this.params[key] = defaultValue
			}
		}

		$effect(() => {
			clearTimeout(this.debounceTimer)
			console.log('params changed', this.params)

			const tempParams: Record<string, string | string[]> = {}

			Object.entries(this.params).forEach(([key, val]) => {
				if (val === undefined || val === null) return

				const paramConfig = this.config[key]
				if (!paramConfig) return // skip unconfigured params

				const shouldInclude =
					(typeof val === 'string' && val !== '') ||
					(Array.isArray(val) && val.length > 0) ||
					(typeof val === 'number' && !isNaN(val)) ||
					typeof val === 'boolean'

				if (shouldInclude) {
					tempParams[key] = this.serializeValue(val, paramConfig.type)
				}
			})

			const searchParams = new SvelteURLSearchParams()
			Object.entries(tempParams).forEach(([key, val]) => {
				if (Array.isArray(val)) {
					searchParams.set(key, val.join(','))
				} else {
					searchParams.set(key, val.toString())
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

	private parseValue(value: string, type: ParamType): any {
		if (type === 'string') return value
		if (type === 'string[]') {
			const tmp = value.split(',')
			if (tmp.length == 0) {
				return [value]
			}
		}
		if (type === 'number') return parseFloat(value)
		if (type === 'number[]') return value.split(',').map(Number)
		if (type === 'boolean') return value === 'true'
		return value
	}

	private serializeValue(value: any, type: ParamType): string | string[] {
		if (Array.isArray(value)) return value.map(String)
		if (type === 'boolean') return value ? 'true' : 'false'
		return String(value)
	}
}
