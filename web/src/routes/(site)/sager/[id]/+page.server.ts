import { API_URL } from '$lib/server/env'
import { error } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import type { TicketWithLogs } from '$lib/types'

export const load: PageServerLoad = async ({ params }) => {
	const id = params.id

	const ticketResponse = await fetch(`${API_URL}/v1/tickets/${id}`)

	if (!ticketResponse.ok) {
		if (ticketResponse.status === 404) {
			error(404, {
				message: 'RMA sag blev ikke fundet.',
			})
		} else if (ticketResponse.status === 429) {
			error(429, {
				message: 'For mange forspørgelser til serveren. Vent et øjeblik.',
			})
		} else {
			error(500, 'Der skete en fejl på serveren.')
		}
	}

	const { ticket } = await ticketResponse.json()

	return {
		ticket: ticket as TicketWithLogs,
	}
}
