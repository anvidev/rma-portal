import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs))
}

export const formatDate = (dateString: string) => {
	return new Date(dateString).toLocaleString('da-DK', {
		day: '2-digit',
		month: '2-digit',
		year: '2-digit',
		hour: '2-digit',
		minute: '2-digit',
	})
}

export function censor(s: string): string {
	return s.replaceAll(/./g, '*')
}

export function censorPhone(s: string): string {
	let count = 0
	let result = ''

	for (let i = s.length - 1; i >= 0; i--) {
		const char = s[i]
		if (char !== ' ' && /\d/.test(char)) {
			if (count < 4) {
				result = char + result
				count++
			} else {
				result = '*' + result
			}
		} else {
			result = char + result
		}
	}

	return result
}

export function censorEmail(email: string): string {
	const [local, domain] = email.split('@')
	if (!domain) return email

	if (local.length <= 2) {
		return '*'.repeat(local.length) + '@' + domain
	}

	const first = local[0]
	const last = local[local.length - 1]
	const stars = '*'.repeat(local.length - 2)
	return `${first}${stars}${last}@${domain}`
}
