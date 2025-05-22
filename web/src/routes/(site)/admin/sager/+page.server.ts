import { API_URL } from '$lib/server/env'
import type { Ticket } from '$lib/types'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, url, setHeaders }) => {
	const searchParams = url.searchParams

	const ticketsResponse = await fetch(`${API_URL}/v1/admin/tickets?${searchParams.toString()}`, {
		headers: {
			Authorization: `Bearer ${cookies.get('token')}`,
		},
	})
	const { tickets, total, limit } = await ticketsResponse.json()

	const statusResponse = await fetch(`${API_URL}/v1/tickets/statuses`)
	const categoriesResponse = await fetch(`${API_URL}/v1/tickets/categories`)

	const { statuses } = await statusResponse.json()
	const { categories } = await categoriesResponse.json()

	setHeaders({
		'Cache-Control': 'private, max-age=120',
	})

	return {
		tickets: tickets as Ticket[],
		statuses: statuses as string[],
		categories: categories as string[],
		total: total as number,
		limit: limit as number,
	}
}
