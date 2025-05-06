import { API_URL } from '$lib/server/env'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, url }) => {
	const searchParams = url.searchParams

	const response = await fetch(new URL(`${API_URL}/v1/admin/tickets?${searchParams.toString()}`), {
		headers: {
			Authorization: `Bearer ${cookies.get('token')}`,
		},
	})
	const { tickets } = await response.json()

	return {
		tickets: tickets as Ticket[],
	}
}
