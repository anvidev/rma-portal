import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, url }) => {
	const searchParams = url.searchParams

	const response = await fetch(
		new URL(`http://localhost:8080/v1/admin/tickets?${searchParams.toString()}`),
		{
			headers: {
				Authorization: `Bearer ${cookies.get('token')}`
			}
		}
	)
	const { tickets } = await response.json()

	return {
		tickets: tickets as Ticket[]
	}
}
