import { error } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { ApiError } from '$lib/server/api'

export const load: PageServerLoad = async ({ params, locals }) => {
	const id = params.id

	const [ticketData, err] = await locals.api.getPublicTicket(id)
	if (err != null) {
		if (err instanceof ApiError) {
			return error(err.status, { message: err.message, requestId: err.requestID })
		} else {
			return error(500, err.message)
		}
	}

	const { ticket } = ticketData

	return {
		ticket: ticket,
	}
}
