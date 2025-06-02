import { error } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { ApiError } from '$lib/server/api'

export const load: PageServerLoad = async ({ params, locals }) => {
	const id = params.id

	const [ticketData, err] = await locals.api.getPublicTicket(id)
	if (err != null) {
		let msg = 'Intern server fejl'
		if (err instanceof ApiError) {
			if (err.status === 429) {
				msg = 'For mange forspørgelser til serveren. Vent et øjeblik.'
			} else if (err.status === 404) {
				msg = 'RMA sag blev ikke fundet.'
			}
			return error(err.status, { message: err.message, requestId: err.requestID })
		} else {
			return error(500, msg)
		}
	}

	const { ticket } = ticketData

	return {
		ticket: ticket,
	}
}
